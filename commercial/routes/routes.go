package routes

import (
	"commercial-propfloor/controller/location"

	"github.com/gin-gonic/gin"
)

/* func PublicRoutes(publicRoutes *gin.Engine){
	publicRoutes.GET("/admin",controller.Anki())
} */

func PrivateRoutes(privateRoutes *gin.Engine) {
	privateRoutes.GET("/commencialRequest/SelectLocality", location.CreateCity())
}
