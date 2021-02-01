package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"group-management-api/adapter"
	"group-management-api/adapter/restapi"
	"group-management-api/domain/usecase/listgroup"
	"group-management-api/domain/usecase/listuser"
	"group-management-api/domain/usecase/managegroup"
	"group-management-api/domain/usecase/manageuser"
	"group-management-api/domain/usecase/userregistration"
	"group-management-api/service/dataservice/postgresds"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	fmt.Printf("Endpoint port: %s\n", port)
	fmt.Printf("Postgres host: %s\n", postgresHost)
	fmt.Printf("Postgres password: %s\n", postgresPassword)
	fmt.Printf("Postgres user: %s\n", postgresUser)

	// Postgres DB
	if postgresUser == "" || postgresPassword == "" || port == "" {
		log.Fatalf("Required environment variables not set.")
	}

	// Database connection pointer.
	DB := postgresds.NewDatabaseConnection(&pg.Options{
		User:     postgresUser,
		Password: postgresPassword,
		Addr:     postgresHost,
		PoolSize: 20,
	})
	defer DB.Close()

	// Create our database schema with this options.
	options := orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := postgresds.CreateSchema(DB, &options)
	if err != nil {
		log.Fatalf("Unable to create database schema:\n%v", err)
	}

	////////// INJECTION PART //////////////

	GroupDataService := postgresds.GroupData{DB: DB}
	UserDataService := postgresds.UserData{DB: DB}

	lguc := listgroup.ListGroupUseCase{
		GroupData: GroupDataService,
	}
	luuc := listuser.ListUserUseCase{
		UserData: UserDataService,
	}
	mguc := managegroup.ManageGroupUseCase{
		GroupData: GroupDataService,
	}
	muuc := manageuser.ManageUserUseCase{
		UserData: UserDataService,
	}
	uruc := userregistration.UserRegistrationUseCase{
		UserData: UserDataService,
	}
	
	bd := adapter.BusinessDomain{
		ListUser:         luuc,
		ListGroup:        lguc,
		ManageGroup:      mguc,
		ManageUser:       muuc,
		UserRegistration: uruc,
	}

	router := chi.NewRouter()
	setupMiddleWare(router) // TODO, do this somewhere else.

	var api adapter.ApiInterface = restapi.RestApi{}
	api.SetupRouter(&bd, router)

	// Start our HTTP server.
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatalf("Can not start server %v", err)
	}
}

func setupMiddleWare(router *chi.Mux) {
	// CORS Options.
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)
	// Other settings.
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Compress(6, "application/json"))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(60 * time.Second))
}