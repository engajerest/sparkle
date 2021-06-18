package app

import (
	"github.com/engajerest/auth/controller"
	"github.com/engajerest/sparkle/controllers"
	// "github.com/gin-gonic/gin"
)

func Mapurls() {

//dev

	router.GET("/dev", controller.PlaygroundHandlers())
	router.POST("/dev/sparkle", controllers.GraphHandler())
//live
	router.GET("/v1", controller.PlaygroundHandlers())
	router.POST("/v1/sparkle", controllers.GraphHandler())
}
