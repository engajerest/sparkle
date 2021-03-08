package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/sparkle/Models/subscription"
	"github.com/spf13/viper"
)

func ParseToken(tokenStr string) (userid, configid float64, Error error) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	Key := viper.GetString("APP.JWT_SECRET_KEY")
	SecretKey := []byte(Key)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userid := claims["userid"].(float64)
		configid := claims["configid"].(float64)
		var tm time.Time
		switch iat := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(iat), 0)
		case json.Number:
			v, _ := iat.Int64()
			tm = time.Unix(v, 0)
		}

		fmt.Println(tm)
		logger.Time("expiry time", tm)
		return userid, configid, nil
	} else {
		return 0, 0, err
	}
}
func ForSparkleContext(ctx context.Context) (*subscription.User, *Errors.RestError) {
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
