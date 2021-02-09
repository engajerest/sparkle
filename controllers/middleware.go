package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"os"

	"github.com/spf13/viper"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/engajerest/auth/Models/users"

	"github.com/gin-gonic/gin"

	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/accesstoken"

	"github.com/engajerest/sparkle/Models/subscription"
	"github.com/engajerest/sparkle/graph"
	"github.com/engajerest/sparkle/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
)

func PlaygroundHandlers() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

func GraphHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func TokenAuthMiddleware(contextkey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("tkn1")
		token := c.Request.Header.Get("token")

		fmt.Println("tkn2")
		print(token)
		if token == "" {
			// c.JSON(http.StatusUnauthorized, "token null")
			// c.Abort()
			c.Next()
			return
		}
		fmt.Println("tkn3")
		userId, err := accesstoken.ParseToken(token)
		fmt.Println("5")
		if err != nil {
			c.JSON(http.StatusUnauthorized, "token denied")
			c.Abort()
			return
		}
		fmt.Println("tkn4")
		id := int(userId)
		// create user and check if user exists in db
		// data1 := users.User{}
		user, status, errrr := subscription.UserAuthentication(int64(id))
		print(status)
		// print("testing")
		// user, err := data1.GetByUserId(int64(id))
		if errrr != nil {
			c.JSON(http.StatusBadRequest, "user not found")
			c.Abort()
			return
		}
		print(user.CreatedDate)

		ctx := context.WithValue(c.Request.Context(), contextkey, user)

		// and call the next with our new context
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		// put it in context
		// ctx := context.WithValue(c.Request.Context(), userCtxKey, user)

		// 	// and call the next with our new context
		// 	r= c.Request.WithContext(ctx)

		// 		c.Next()

	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForSparkleContext(ctx context.Context) (*users.User, *Errors.RestError) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	userCtxKey := viper.GetString("APP.USER_CONTEXT_KEY")
	noUserFoundError := errors.New("no user found")
	if ctx.Value(userCtxKey) == nil {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusBadRequest,
		}
	}
	user, ok := ctx.Value(userCtxKey).(*users.User)
	if !ok || user.ID == 0 {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusBadRequest,
		}
	}
	return user, nil
}
// func Location(c *gin.Context) {

// 	var locationGetAll []subscription.Tenantlocation
//     locationGetAll = subscription.LocationTest(37)
// 	c.JSON(http.StatusOK, locationGetAll)
// }
