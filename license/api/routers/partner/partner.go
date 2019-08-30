package partner

import (
	"license/api/controller"

	"github.com/gin-gonic/gin"
)

//RegisterRoutes Partner
func RegisterRoutes(router *gin.RouterGroup) {
	pc := new(controller.PartnerController)
	router.GET("/", pc.ListPartner)
	router.POST("/", pc.AddPartner)
	router.POST("/:id/products", pc.AddProduct)

}
