package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// Initialise api router group
	// apiGroup := r.Group("/api")
	//api.InitApiGroup(apiGroup) // this will be implemented later depending on plans

	return r
}
