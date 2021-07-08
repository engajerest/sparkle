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
	router.POST("/dev/sparkle/fstenant",controllers.Tenantinsert)
	router.POST("/dev/sparkle/fstenantupdate",controllers.Tenantupdate)
	router.POST("/dev/sparkle/fslocationcreate",controllers.Locationcreate)
	router.POST("/dev/sparkle/fslocationupdate",controllers.Locationupdate)
//live
	router.GET("/v1", controller.PlaygroundHandlers())
	router.POST("/v1/sparkle", controllers.GraphHandler())
	router.POST("/v1/sparkle/fstenant",controllers.Tenantinsert)
	router.POST("/v1/sparkle/fstenantupdate",controllers.Tenantupdate)
	router.POST("/v1/sparkle/fslocationcreate",controllers.Locationcreate)
	router.POST("/v1/sparkle/fslocationupdate",controllers.Locationupdate)
}
