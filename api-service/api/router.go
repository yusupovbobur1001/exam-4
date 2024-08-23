package api

import (
	_ "api_service/api/docs"
	"api_service/api/handler"
	"api_service/api/middleware"
	"api_service/config"
	"api_service/pkg/kafka"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        On-Demand API
// @version      1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @security [ApiKeyAuth]
func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	enforcer, err := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	kafkaProducer, err := kafka.NewKafkaProducerInit([]string{"kafka1:9093"}) //kafka1:9092
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(cfg, kafkaProducer)

	router.Use(middleware.JWTMiddleware())
	router.Use(middleware.CasbinMiddleware(enforcer))

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.PUT("/users/:id", h.UpdateUserProfile)
		auth.DELETE("/users/:id", h.DeleteUserProfile)
		auth.GET("/users/:id", h.GetByIdProfile)
		auth.GET("/users", h.GetAllProfile)
	}

	booking := router.Group("/booking")
	{
		booking.POST("/bookings", h.CreateBooking)
		booking.GET("/bookings/:id", h.GetBooking)
		booking.PUT("/bookings/:id", h.UpdateBooking)
		booking.DELETE("/bookings/:id", h.DeleteBooking)
		booking.GET("/bookings", h.GetBookingList)
		booking.GET("/most-frequent-service-id", h.GetMostFrequentServiceID)
	}

	payment := router.Group("/payment")
	{
		payment.POST("/payments", h.CreatePayment)
		payment.GET("/payments/:id", h.GetPayment)
		payment.GET("/payments", h.ListPayments)
	}

	service := router.Group("/service")
	{
		service.POST("/services", h.CreateService)
		service.PUT("/services/:id", h.UpdateService)
		service.DELETE("/services/:id", h.DeleteService)
		service.GET("/services", h.ListServices)
		service.GET("/services/search", h.SearchServices)
	}

	provider := router.Group("/provider")
	{
		provider.POST("/providers", h.CreateProviders)
		provider.PUT("/providers/:id", h.UpdateProviders)
		provider.DELETE("/providers/:id", h.DeleteProviders)
		provider.GET("/providers/:id", h.GetProviders)
		provider.GET("/providers", h.ListProviders)
		provider.GET("/providers/search", h.SearchProviders)
	}

	review := router.Group("/review")
	{
		review.POST("/reviews", h.CreateReview)
		review.PUT("/reviews/:id", h.UpdateReview)
		review.DELETE("/reviews/:id", h.DeleteReview)
		review.GET("/reviews", h.ListReviews)
	}

	return router
}
