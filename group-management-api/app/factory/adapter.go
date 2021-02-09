package factory

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"group-management-api/adapter"
	"group-management-api/adapter/restapi"
	"group-management-api/app/config/impl"
	"group-management-api/app/container"
	"group-management-api/app/logger"
	"log"
	"net/http"
	"time"
)

func InitAdapter(c *container.Container) error {
	switch c.AppConfig.AdapterConfig.Impl {
	case impl.RestAdapter:
		// Construct a REST adapter.
		err := restAdapterSetup(c)
		if err != nil {
			return err
		}

	default:
		// Unknown configuration.
		return errors.New("unknown adapter implementation in configurations")
	}

	return nil
}

func restAdapterSetup(c *container.Container) error {
	router := chi.NewRouter()
	// Default router middleware.
	setupDefaultRouterMiddleware(router, c)

	lg, err := GetListGroupUseCase(c)
	if err != nil {
		return err
	}

	lu, err := GetListUserUseCase(c)
	if err != nil {
		return err
	}

	mg, err := GetManageGroupUseCase(c)
	if err != nil {
		return err
	}

	mu, err := GetManageUserUseCase(c)
	if err != nil {
		return err
	}

	ur, err := GetUserRegistrationUseCase(c)
	if err != nil {
		return err
	}

	// Business domain
	businessDomain := &adapter.BusinessDomain{
		ListGroup: lg,
		ListUser: lu,
		ManageGroup: mg,
		ManageUser: mu,
		UserRegistration: ur,
	}

	// Secret for signing jwt tokens.
	jwtSecret := c.AppConfig.AdapterConfig.RestConfig.JwtKey
	// Setup the router for Rest Api.
	err = restapi.SetupRouterForRestAdapter(businessDomain, router, jwtSecret)

	if err != nil {
		return err
	}

	// RestAPI will be reachable trough this port.
	port := c.AppConfig.ExposedPort

	// Set the startable to HTTP Server.
	c.Startable = func() {
		// Start our Http server and register our router.
		logger.Log.Infof("Rest server started on port %v", port)
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
		if err != nil {
			// Log if it http does not start listening and serving.
			log.Fatalf("Can not start server %v", err)
		}
	}

	// Used for testing.
	c.Adapter = router

	return nil
}

// Default middleware. Could call a different middleware setup function based on config parameters.
func setupDefaultRouterMiddleware(router *chi.Mux, c *container.Container) {
	// CORS Options.
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(corsOptions.Handler)
	// Other settings.
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Compress(6, "application/json"))

	if c.AppConfig.AdapterConfig.RestConfig.EnableLogging {
		router.Use(middleware.Logger)
	}

	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(60 * time.Second))
}