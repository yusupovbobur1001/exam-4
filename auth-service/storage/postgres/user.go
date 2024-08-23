package postgres

import (
	pb "auth_service/genproto/user"
	"auth_service/model"
	"auth_service/storage/repo"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type AuthRepoManagement struct {
	Db *sql.DB
}

func NewAuthRepoManagement(db *sql.DB) repo.StoageI {
	return &AuthRepoManagement{
		Db: db,
	}
}

func (a *AuthRepoManagement) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println(3)
	resp := pb.RegisterResponse{}
	pass := HashPassword(req.Password)
	fmt.Println(pass)

	q := `
	insert into users(first_name, last_name, email, password, phone_number, role, created_at)
	values($1, $2, $3, $4, $5, $6, $7)
	returning id, first_name, last_name, phone_number, created_at
	`

	err := a.Db.QueryRow(q, req.FirstName, req.LastName, req.Email, pass, req.PhoneNumber, req.Role, time.Now()).Scan(
		&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.PhoneNumber,
		&resp.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *AuthRepoManagement) UpdateUserProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	resp := pb.UpdateProfileResponse{}

	q := `
	UPDATE 
		users
	SET 
		first_name = $1,
		phone_number = $2,
		role = $3,
		updated_at = $4
	WHERE 
		id = $5
	AND 
		deleted_at IS NULL
	RETURNING 
		first_name, phone_number, role, updated_at
	`
	err := a.Db.QueryRow(q, req.NewFirstName, req.NewPhoneNumber, req.NewRole, time.Now(), req.Id).Scan(
		&resp.FirstName,
		&resp.PhoneNumber,
		&resp.Role,
		&resp.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &resp, nil

}

func (a *AuthRepoManagement) DeleteUserProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	q := `
	DELETE FROM 
		users
	WHERE 
		id = $1
	AND 
		deleted_at IS NULL
	`

	_, err := a.Db.ExecContext(ctx, q, req.Id)
	if err != nil {
		return &pb.DeleteProfileResponse{
			Message: "Error",
		}, err
	}

	return &pb.DeleteProfileResponse{
		Message: "Success",
	}, nil
}

func (a *AuthRepoManagement) GetByIdProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	fmt.Println(req.Id)
	resp := pb.GetProfileResponse{}
	q := `
	select 
		id,
		role,
		first_name,
		last_name,
		email,
		phone_number,
		created_at
	from 
		users
	where 
		id = $1
	and 
		deleted_at is null
	`

	err := a.Db.QueryRowContext(ctx, q, req.Id).Scan(
		&resp.Id,
		&resp.Role,
		&resp.FirstName,
		&resp.LastName,
		&resp.Email,
		&resp.PhoneNumber,
		&resp.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *AuthRepoManagement) GetAllProfile(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesResponse, error) {
	resp := pb.GetProfilesResponse{}

	q := `
	select  
		id,
		role,
		first_name,
		last_name,
		email,
		phone_number,
		created_at
	from 
		users
	where 
		deleted_at is null
	`

	rows, err := a.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := pb.GetProfileResponse{}
		err := rows.Scan(
			&r.Id,
			&r.Role,
			&r.FirstName,
			&r.LastName,
			&r.Email,
			&r.PhoneNumber,
			&r.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp.AllProfile = append(resp.AllProfile, &r)
	}
	return &resp, nil
}

func (a *AuthRepoManagement) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	resp := model.LoginResponse{}
	pass := HashPassword(req.Password)
	fmt.Println(pass)
	q := `	
	select 
		id,
		role,
		first_name
	from 
		users
	where 
		email = $1
	and 
		 password = $2
	and
		deleted_at is null
	`
	err := a.Db.QueryRowContext(ctx, q, req.Email, pass).Scan(
		&resp.Id,
		&resp.Role,
		&resp.FirstName,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (a *AuthRepoManagement) Logout(ctx context.Context, req *model.LogoutRequest) (*model.LogoutResponse, error) {
	return &model.LogoutResponse{
		Message: "Success",
	}, nil
}
