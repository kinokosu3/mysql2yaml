package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func Engine(user, password, host, port, database string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database))
	if err != nil {
		panic(err)
	}
	return db
}

type TableName []string

func (h *TableName) String() string {
	return fmt.Sprintf("%v", *h)
}

func (h *TableName) Set(s string) error {
	for _, v := range strings.Split(s, ",") {
		*h = append(*h, v)
	}
	return nil
}

var TableNames TableName

func init() {
	flag.Var(&TableNames, "table", "-table=test1,test2")
}

// ./mysql2yaml -user=root -password=-host= -port= -table= -database=
// id 模式
// sql 模式
func main() {
	user := flag.String("user", "root", "mysql username")
	pwd := flag.String("password", "", "mysql password")
	host := flag.String("host", "localhost", "mysql host")
	port := flag.String("port", "3306", "mysql port")
	dbName := flag.String("database", "test", "mysql database, default is test")
	flag.Parse()
	engine := Engine(*user, *pwd, *host, *port, *dbName)
	for _, v := range TableNames {
		data, err := TableData(engine, v)
		if err != nil {
			panic(err)
		}
		CreateYaml(data, v+".yaml")
	}

}

func dataHandle(raw []map[string]interface{}) []map[string]interface{} {
	res := []map[string]interface{}{}
	for _, val := range raw {
		data := make(map[string]interface{})
		for k, v := range val {
			switch v.(type) {
			case []uint8:
				data[k] = string(v.([]uint8))
			default:
				data[k] = v
			}
		}
		res = append(res, data)
	}
	return res
}

func TableData(db *sql.DB, tableName string) ([]map[string]interface{}, error) {
	rows, err := db.Query("select * from " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var result []map[string]interface{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		row := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				row[columns[i]] = col
			}
		}
		result = append(result, row)
	}
	return dataHandle(result), nil

}
