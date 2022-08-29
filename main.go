package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	_ "utest/adapter/http/rest/docs"
	"utest/adapter/http/rest/middleware"
	"utest/adapter/postgres"
	"utest/di"
)

func init() {

	//CARGA BASE.

}

// @title UTest
// @version 1.0.0
// @contact.name Cleber
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
func main() {

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	feiraService := di.ConfigFeiraDI(conn)

	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	jsonApiRouter.Handle("/api/feiras", http.HandlerFunc(feiraService.Create)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/api/feiras/{id}", http.HandlerFunc(feiraService.Update)).Methods("PUT")
	jsonApiRouter.Handle("/api/feiras/{id}", http.HandlerFunc(feiraService.Update)).Methods("DELETE")
	jsonApiRouter.Handle("/api/feiras/{nome_feira}/nomes", http.HandlerFunc(feiraService.GetNome)).Methods("GET")

	serverport := os.Getenv("serverport")

	log.Printf("LISTEN ON PORT: %v", serverport)
	http.ListenAndServe(fmt.Sprintf(":%v", serverport), router)

}
