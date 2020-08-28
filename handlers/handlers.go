//Package handlers implements methods to serve http requests for the API.
package handlers

import (
	"github.com/gin-gonic/gin"
)

//RegisterRoutes registers routes to the appropriate handler functions on the Gin Engine instance.
func RegisterRoutes(handler *gin.Engine) {

	homeHandler := GetHomeHandler()
	userHandler := GetUserHandler()
	handler.GET("/get/", homeHandler.get)
	handler.GET("/", homeHandler.index)

	handler.POST("/post/", userHandler.post)
	handler.PUT("/put/:uuid", userHandler.put)
	handler.DELETE("/delete/:uuid", userHandler.delete)

}
