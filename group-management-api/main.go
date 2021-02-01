package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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



	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}