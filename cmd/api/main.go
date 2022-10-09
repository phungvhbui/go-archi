package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phungvhbui/go-archi/internal/connector"
	"github.com/phungvhbui/go-archi/internal/controller"
	"github.com/phungvhbui/go-archi/internal/externalservice"
	"github.com/phungvhbui/go-archi/internal/repository"
	"github.com/phungvhbui/go-archi/internal/service"
	"github.com/rs/zerolog/log"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check
	health := new(controller.HealthController)
	router.GET("/healthz", health.Status)

	// General deps
	db, err := connector.InitializeDB(
		"mysql", "172.17.0.2", 3306, "test_db", "root", "mypass", "",
	)
	if err != nil {
		panic(err)
	}

	// System deps
	stripe := externalservice.NewStripe()

	transactor := repository.New(db)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, transactor, stripe)

	organizationRepository := repository.NewOrganizationRepository(db)
	organizationService := service.NewOrganizationService(organizationRepository)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			userController := controller.NewUserController(userService)
			userGroup.GET("/", userController.GetAll)
			userGroup.POST("/", userController.Create)
			// organizationGroup.GET("/:id", organizationController.Get)
			// organizationGroup.POST("/", organizationController.Create)
		}

		organizationGroup := v1.Group("organizations")
		{
			organizationController := controller.NewOrganizationController(organizationService)
			organizationGroup.GET("/", organizationController.GetAll)
		}

	}
	return router
}

func main() {
	r := NewRouter()

	hostPort := "localhost" + ":" + "3000"
	log.Info().Msgf("server started at %s", hostPort)

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
