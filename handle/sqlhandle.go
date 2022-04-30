package handle

import (
	"database/sql"
	"fmt"
	"strings"
)

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

func TableData(db *sql.DB, sql string) ([]map[string]interface{}, error) {
	fmt.Println(sql)
	rows, err := db.Query(sql)
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

func relationSql(tName []string, sql string) []string {
	resSql := make([]string, 0, len(tName))
	// get first 'from' index
	fromIndex := strings.Index(sql, "from")
	suffixSql := sql[fromIndex:]
	for _, v := range tName {
		resSql = append(resSql, "select "+v+".* "+suffixSql)
	}
	return resSql

}

func DoSql(db *sql.DB, t TableList, s SqlList, gLimit string) {
	// table flag
	if len(t) != 0 {
		for _, v := range t {
			bufSql := fmt.Sprintf("select * from %s", v)
			if gLimit != "" {
				bufSql = fmt.Sprintf("%s %s", bufSql, gLimit)
			}
			data, err := TableData(db, bufSql)
			if err != nil {
				panic(err)
			}
			CreateYaml(data, v+".yaml")
		}
	}
	// sql flag
	if len(s) != 0 {
		//key is table name, value is sql
		for k, v := range s {
			bufK := strings.Split(k, ",")
			if len(bufK) >= 2 {
				l := relationSql(bufK, v)
				for index, val := range l {
					data, err := TableData(db, val)
					if err != nil {
						panic(err)
					}
					CreateYaml(data, bufK[index]+".yaml")
				}
				return
			}
			data, err := TableData(db, v)
			if err != nil {
				panic(err)
			}
			CreateYaml(data, k+".yaml")
		}
	}
}
