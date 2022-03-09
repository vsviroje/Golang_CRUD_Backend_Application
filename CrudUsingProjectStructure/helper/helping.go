package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const datasourcename = "root:toor@tcp(127.0.0.1:3306)/fun"
const driver = "mysql"

func ReadJSONFileData() map[string]string {
	var MappedData map[string]string
	filePath := path.Join("qry", "BookQry.json")
	JSONfileData, _ := ioutil.ReadFile(filePath)
	_ = json.Unmarshal([]byte(JSONfileData), &MappedData)
	return MappedData
}

func Connect_to_database() *sqlx.DB {
	db, err := sqlx.Connect(driver, datasourcename)

	if err != nil {
		fmt.Println("failed to connect to database")
	} else if err == nil {
		fmt.Println("connected to database")
	}
	return db
}
