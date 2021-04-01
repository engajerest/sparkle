package main

import (
	"fmt"
	"os"

	"github.com/engajerest/auth/controller"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/dbconfig"
	"github.com/engajerest/sparkle/controllers"
	"github.com/gin-gonic/gin"
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
	UserContextKey := viper.GetString("APP.USER_CONTEXT_KEY")

	fmt.Println("PORT :", defaultPort)

	router := gin.Default()
	// router.GET("/locations", controllers.Location)
	router.Use(controller.TokenNoAuthMiddleware(UserContextKey))

	dbconfig.InitDB(dbName, userName, password, host)
	logger.Info("sparkle application started")

	router.GET("/", controllers.PlaygroundHandlers())
	router.POST("/sparkle", controllers.GraphHandler())
	router.Run(defaultPort)
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//  http.Handle("/", playground.Handler("Engaje", "/query"))
	// // router.Handle("/middle", middleware(http.HandlerFunc(pong)))
	// http.Handle("/query", srv)
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	// log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
