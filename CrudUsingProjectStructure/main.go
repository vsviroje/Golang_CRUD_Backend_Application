package main

import (
	wh "CrudUsingProjectStructure/WebHandler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	wh.RequestHandler(router)
	router.Run()
}
