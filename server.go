package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/engajerest/auth/controller"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/dbconfig"
	"github.com/engajerest/sparkle/graph"
	"github.com/engajerest/sparkle/graph/generated"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)



func main() {

	// Config
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Declare var
	defaultPort := viper.GetString("APP.PORT")
	dbName := viper.GetString("APP.DATABASE_NAME")
	password := viper.GetString("APP.DATABASE_PASSWORD")
	userName := viper.GetString("APP.DATABASE_USERNAME")
	_ = viper.GetString("APP.DATABASE_PORT")
	host := viper.GetString("APP.DATABASE_SERVER_HOST")

	fmt.Println("PORT :", defaultPort)

	router := chi.NewRouter()
	router.Use(controller.Middleware())
	dbconfig.InitDB(dbName, userName, password, host)
    logger.Info("sparkle application started")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/sparkle", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
