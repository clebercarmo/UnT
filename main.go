package main

import (
	"context"
	"fmt"
	"net/http"
	_ "utest/adapter/http/rest/docs"
	"utest/adapter/http/rest/middleware"
	"utest/adapter/postgres"
	"utest/di"
	"utest/shared"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

// @title UTest
// @version 1.0.0
// @contact.name Cleber
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()
	shared.SetLog("Info", "Executando Migrations")
	postgres.RunMigrations()
	shared.SetLog("Info", "Migration executada")
	feiraService := di.ConfigFeiraDI(conn)
	
	
	shared.SetLog("Info", "Importando dados")
	shared.ImportarDados()
	shared.SetLog("Info", "Dados importados")

	router := mux.NewRouter()

	router.PathPrefix("/api/swagger/").Handler(httpSwagger.WrapHandler)

	jsonApiRouter := router.PathPrefix("/").Subrouter()
	jsonApiRouter.Use(middleware.Cors)

	jsonApiRouter.Handle("/api/feiras", http.HandlerFunc(feiraService.Create)).Methods("POST", "OPTIONS")
	jsonApiRouter.Handle("/api/feiras/{id:[0-9]+}", http.HandlerFunc(feiraService.Update)).Methods("PUT")
	jsonApiRouter.Handle("/api/feiras/{id}", http.HandlerFunc(feiraService.Delete)).Methods("DELETE")
	jsonApiRouter.Handle("/api/feiras/{nome_feira}/nomes", http.HandlerFunc(feiraService.GetNome)).Methods("GET")

	serverport := viper.GetString("server.port")
	
	shared.SetLog("Info", "Servidor disponivel na porta 3000")
	http.ListenAndServe(fmt.Sprintf(":%v", serverport), router)

}
