package library

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/lib/pq" //posgres driver
)

//Database (database object)
type Database struct {
	SQL *sql.DB
}

//Connect (connect to database)
func (d *Database) Connect(driver, host, port, database, username, password string) *sql.DB {
	conStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s", driver, username, password, host, port, database)
	db, err := sql.Open(driver, conStr)
	if err != nil {
		log.Fatal("Koneksi database gagal")
	}
	d.SQL = db
	return db
}

//CreateTable (create table in database using stuct model)
func (d *Database) CreateTable(model interface{}) {
	ref := reflect.TypeOf(model)
	modelName := ref.Name()
	sqlFields := []string{}
	for i := 0; i < ref.NumField(); i++ {
		sqlFields = append(sqlFields, ref.Field(i).Tag.Get("json")+" "+ref.Field(i).Tag.Get("sql"))
	}
	str := strings.Join(sqlFields, ",")
	sql := fmt.Sprintf("create table if not exists %s (%s)", modelName, str)
	_, err := d.SQL.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
}
