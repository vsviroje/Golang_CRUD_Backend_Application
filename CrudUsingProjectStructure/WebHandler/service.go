package WebHandler

import (
	DBA "CrudUsingProjectStructure/DBAccess"
	mod "CrudUsingProjectStructure/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AllBook []mod.Book
type SingleBook mod.Book

func GetAllBook(c *gin.Context) {
	var bk []mod.Book
	bk = DBA.Getdata()
	c.JSON(200, bk)
}
func InsertBook(c *gin.Context) {
	var bk mod.Book
	c.Bind(&bk)
	i := DBA.Postdata(bk)
	c.JSON(200, i)

}
func UpdateBook(c *gin.Context) {
	var bk mod.Book
	c.Bind(&bk)
	i := DBA.Updateid(bk)
	c.JSON(200, i)
}
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Println(id)
	i := DBA.Deleteid(int32(id))
	c.JSON(200, i)
}
