package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysql2yaml/handle"
)

func Engine(user, password, host, port, database string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database))
	if err != nil {
		panic(err)
	}
	return db
}

var SqlL handle.SqlList
var TableList handle.TableList

func init() {
	SqlL = make(handle.SqlList)
	flag.Var(&SqlL, "sql", "-sql user='select * from user' -sql role='select * from role' ")
	flag.Var(&TableList, "table", "-table=user,role")
}

// ./mysql2yaml -user=root -password=-host= -port= -table= -database=

func main() {
	user := flag.String("user", "root", "mysql username")
	pwd := flag.String("password", "", "mysql password")
	host := flag.String("host", "localhost", "mysql host")
	port := flag.String("port", "3306", "mysql port")
	dbName := flag.String("database", "test", "mysql database, default is test")
	globalLimit := flag.String("cond", "", `use by -table, example: -where="where id=1"`)
	flag.Parse()
	engine := Engine(*user, *pwd, *host, *port, *dbName)
	handle.DoSql(engine, TableList, SqlL, *globalLimit)
}
