package routers

import (
	"license/api/routers/partner"

	"github.com/gin-gonic/gin"
)

// SetupRouter Register All Routers
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("licence/api")
	partner.RegisterRoutes(v1.Group("partner"))

	return r
}
