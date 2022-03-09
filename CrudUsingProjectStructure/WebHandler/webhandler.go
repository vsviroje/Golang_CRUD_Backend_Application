package WebHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestHandler(conf *gin.Engine) {

	conf.GET("/get", handle)
	conf.POST("/post", handle)
	conf.PUT("/put", handle)
	conf.DELETE("/delete", handle)

}

func handle(c *gin.Context) {
	switch c.Request.RequestURI {

	case "/get":
		GetAllBook(c)
	case "/post":
		InsertBook(c)
	case "/put":
		UpdateBook(c)
	case "/delete":
		DeleteBook(c)
	default:
		DeleteBook(c)
		fmt.Println("Route went wrong")
	}
}
