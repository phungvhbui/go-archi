package main

import (
	"github.com/gin-gonic/gin"
	"github.com/phungvhbui/go-archi/internal/api/health"
	"github.com/phungvhbui/go-archi/internal/api/organization"
	"github.com/phungvhbui/go-archi/internal/api/user"
	"github.com/phungvhbui/go-archi/internal/datastore"
	"github.com/phungvhbui/go-archi/internal/datastore/repository"
	"github.com/phungvhbui/go-archi/internal/datastore/transaction"
	"github.com/phungvhbui/go-archi/internal/stripe"
	"github.com/rs/zerolog/log"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check
	health := new(health.HealthController)
	router.GET("/healthz", health.Status)

	// General deps
	db, err := datastore.InitializeDB(
		"mysql", "172.17.0.2", 3306, "test_db", "root", "mypass", "",
	)
	if err != nil {
		panic(err)
	}

	// System deps
	stripeInstance := stripe.NewStripe()

	transactor := transaction.NewTransactor(db)

	userRepository := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepository, transactor, stripeInstance)

	organizationRepository := repository.NewOrganizationRepository(db)
	organizationService := organization.NewOrganizationService(organizationRepository)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			userController := user.NewUserController(userService)
			userGroup.GET("/", userController.GetAll)
			userGroup.POST("/", userController.Create)
			// organizationGroup.GET("/:id", organizationController.Get)
			// organizationGroup.POST("/", organizationController.Create)
		}

		organizationGroup := v1.Group("organizations")
		{
			organizationController := organization.NewOrganizationController(organizationService)
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
