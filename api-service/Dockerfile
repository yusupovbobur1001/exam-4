FROM golang:1.22.6 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .
COPY --from=builder /app/casbin/model.conf /app/casbin/model.conf
COPY --from=builder /app/casbin/policy.csv /app/casbin/policy.csv


EXPOSE 8080

CMD ["./myapp"]
