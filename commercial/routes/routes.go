package routes

import (
	"commercial-propfloor/controller/location"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(privateRoutes *gin.Engine) {
    privateRoutes.POST("/Manage/Commercialportal/SaveBasicBuildingDetails",location.InsertInDatabase())
	// privateRoutes.GET("/commercialRequest/SelectLocality/insertcity", location.InsertInDatabase())
	privateRoutes.GET("/commercialRequest/SelectLocality/Deletecity", location.DeleteInDatabase())
	privateRoutes.GET("/commercialRequest/SelectLocality/Updatecity", location.UpdateInDatabase())
	privateRoutes.GET("/commercialRequest/SelectLocality/Selectindatabase", location.Selectindatabase())
	privateRoutes.GET("/commercialRequest/SelectLocality/Queryindatabase", location.Queryindatabase())
}

func PublicRoutes(publicRoutes *gin.Engine) {

}
