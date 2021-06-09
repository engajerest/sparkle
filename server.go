package main

import (
	"github.com/engajerest/sparkle/app"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	app.StartApplication()
	// router := gin.Default()

	// v2 := router.Group("/live")

	// Split after an empty string to get all letters.

	// Config
	// viper.SetConfigName("config") // config file name without extension
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	fmt.Println("fatal error config file: default \n", err)
	// 	os.Exit(1)
	// }

	// // Declare var
	// defaultPort := viper.GetString("APP.PORT")
	// dbName := viper.GetString("APP.DATABASE_NAME")
	// password := viper.GetString("APP.DATABASE_PASSWORD")
	// userName := viper.GetString("APP.DATABASE_USERNAME")
	// _ = viper.GetString("APP.DATABASE_PORT")
	// host := viper.GetString("APP.DATABASE_SERVER_HOST")
	// UserContextKey := viper.GetString("APP.USER_CONTEXT_KEY")
	// fmt.Println("PORT :", defaultPort)
	// router.Use(controller.TokenNoAuthMiddleware(UserContextKey))
	// dbconfig.InitDB(dbName, userName, password, host)
	// logger.Info("sparkle application started ")
	// router.GET("/", controller.PlaygroundHandlers())
	// router.POST("/sparkle", controllers.GraphHandler())
	// router.Run(defaultPort)

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//  http.Handle("/", playground.Handler("Engaje", "/query"))
	// // router.Handle("/middle", middleware(http.HandlerFunc(pong)))
	// http.Handle("/query", srv)
	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	// log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
