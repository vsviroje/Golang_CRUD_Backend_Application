package DBAccess

import (
	help "CrudUsingProjectStructure/helper"
	mod "CrudUsingProjectStructure/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var databasesqlx *sqlx.DB

type AllBook []mod.Book
type SingleBook mod.Book
type mapJSON map[string]string

func Getdata() []mod.Book {
	var bk []mod.Book
	db := help.Connect_to_database()
	var JSONKeyValue mapJSON = help.ReadJSONFileData()

	qry := JSONKeyValue["SELECTALL"]

	error := db.Select(&bk, qry)
	if error != nil {
		fmt.Println("Failed to execute query or to bind data with structure")
	}
	return bk
}

func Updateid(bk mod.Book) int32 {

	db := help.Connect_to_database()
	var JSONKeyValue mapJSON = help.ReadJSONFileData()

	qry := JSONKeyValue["UPDATEID"]

	result := db.MustExec(qry, bk.BookName, bk.BookDetails, bk.BookAuthor, bk.BookPrice, bk.BookRequest, bk.BookID)

	i, _ := result.RowsAffected()
	if i < 1 {
		fmt.Println("Failed to insert")
		i = 0
	}
	return int32(i)
}

func Postdata(bk mod.Book) int32 {

	db := help.Connect_to_database()
	var JSONKeyValue mapJSON = help.ReadJSONFileData()
	qry := JSONKeyValue["INSERT"]

	result := db.MustExec(qry, bk.BookName, bk.BookDetails, bk.BookAuthor, bk.BookPrice, bk.BookRequest)

	i, _ := result.RowsAffected()
	if i < 1 {
		fmt.Println("Failed to insert")
		i = 0
	}
	return int32(i)
}

func Deleteid(id int32) int32 {

	db := help.Connect_to_database()
	var JSONKeyValue mapJSON = help.ReadJSONFileData()
	qry := JSONKeyValue["DELETEID"]
	result := db.MustExec(qry, id)

	i, _ := result.RowsAffected()

	if i < 1 {
		fmt.Println("Failed to insert")
		i = 0
	}

	return int32(i)

}
