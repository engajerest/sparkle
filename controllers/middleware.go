package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"os"

	"github.com/spf13/viper"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"

	"github.com/engajerest/auth/utils/Errors"

	"github.com/engajerest/sparkle/Models/subscription"
	"github.com/engajerest/sparkle/graph"
	"github.com/engajerest/sparkle/graph/generated"
	"github.com/engajerest/sparkle/helper"

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
	
		userId, configid, err := helper.ParseToken(token)
		// fmt.Println("5")
		if err != nil {
			c.JSON(http.StatusUnauthorized, "token denied")
			c.Abort()
			return
		}
		// fmt.Println("tkn4")
		id := int(userId)
		id1 := int(configid)
		print("confiid", id1)
		if id1 == 1 {
			print("configid==1")
			
			user, status, errrr := subscription.UserAuthentication(int64(id))
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}
			print(user.CreatedDate)
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			print("configid>1")
			
			user, status, errrr := subscription.Customerauthenticate(int64(id))
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}
			print(user.CreatedDate)
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()

		}

	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForSparkleContext1(ctx context.Context) (*subscription.User, *Errors.RestError) {
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
	user, ok := ctx.Value(userCtxKey).(*subscription.User)
	if !ok || user.ID == 0 {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusBadRequest,
		}
	}
	return user, nil
}

func Tenantinsert(c *gin.Context) {

	var fs subscription.Fstenant
	var res subscription.Result
	if err := c.ShouldBindJSON(&fs); err != nil {
		formatErr := subscription.NewBadRequestErr("invalid json body")
		c.JSON(http.StatusBadRequest, formatErr)
		return
	}
	

	custerr := fs.Firestoretenantweb()
	if custerr != nil {
		res.Status=false
		res.Code=http.StatusBadRequest
		res.Message="Internal Error in Firestore"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res.Status=true
	res.Code=http.StatusCreated
	res.Message="Tenant Created in Firestore Sucessfully"
	c.JSON(http.StatusCreated,res )
}