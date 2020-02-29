package task1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	Version = 1
)

func ToPostgresFunc(uri, method string) (funcName string, id []interface{}) {
	if strings.Index(uri, ApiURI) == 0 {
		uri = strings.Replace(uri, ApiURI, "", 1)
	}
	if strings.Index(uri, "/") == 0 {
		uri = uri[1:len(uri)]
	}
	params := strings.Split(uri, `/`)
	var result []string
	for index, object := range params {
		if index%2 == 0 {
			result = append(result, object)
		} else {
			num, _ := strconv.ParseInt(object, 10, 64)
			id = append(id, num)
		}
	}
	funcName = strings.Join(result, "_")
	switch strings.ToUpper(method) {
	case "GET":
		funcName += "_get"
	case "POST":
		funcName += "_ins"
	case "PUT":
		funcName += "_upd"
	case "DELETE":
		funcName += "_del"
	}
	return
}

type controllerAnswer struct {
	FuncName string
	Params   []interface{}
	Body     string
}

func PostgresParamsMixin(inData []interface{}) string {
	var count = len(inData)
	var resultSl []string
	for i := 1; i <= count; i++ {
		resultSl = append(resultSl, fmt.Sprintf("$%v", i))
	}
	return strings.Join(resultSl, ",")
}

func Controller(writer http.ResponseWriter, request *http.Request) {
	funcName, params := ToPostgresFunc(request.URL.Path, request.Method)
	result := controllerAnswer{
		FuncName: funcName,
		Params:   params,
	}
	switch strings.ToUpper(request.Method) {
	case "POST", "PUT":
		byteSl, _ := ioutil.ReadAll(request.Body)
		result.Params = append(result.Params, string(byteSl))
	}
	result.FuncName = fmt.Sprintf("SELECT * FROM %v.%v(%v);", SchemaEnv, result.FuncName, PostgresParamsMixin(result.Params))
	_ = makeStructJSON(result.FuncName, result.Params, writer)
}

func makeStructJSON(queryText string, args []interface{}, w http.ResponseWriter) error {
	rows, err := db.Query(queryText, args...)
	if err != nil {
		return err
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var masterData []map[string]interface{}
	for rows.Next() {
		sqlData := make(map[string]interface{})
		err := rows.Scan(scanArgs...)
		if err != nil {
			return err
		}
		for i, v := range values {
			sqlData[columns[i]] = v
		}
		masterData = append(masterData, sqlData)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(masterData)
	if err != nil {
		return err
	}
	return err
}
