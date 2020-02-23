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
	if strings.Index(uri, apiURI) == 0 {
		uri = strings.Replace(uri, apiURI, "", 1)
	}
	params := strings.Split(uri, `/`)
	var result []string
	for index, object := range params {
		if index%2 == 0 {
			result = append(result, object)
		} else {
			num, _ := strconv.ParseInt(object, 10, 64)
			id = append(id, int(num))
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
	writer.Header().Set("Content-Type", "application/json")
	funcName, params := ToPostgresFunc(request.URL.Path, request.Method)
	result := controllerAnswer{
		FuncName: funcName,
		Params:   params,
	}
	switch strings.ToUpper(request.Method) {
	case "POST", "PUT":
		var data map[string]interface{}
		byteSl, _ := ioutil.ReadAll(request.Body)
		result.Body = string(byteSl)
		_ = json.Unmarshal(byteSl, &data)
		for name, val := range data {
			result.Params = append(result.Params, name, val)
		}
	}
	result.FuncName = fmt.Sprintf("SELECT * FROM %v.%v(%v);", SchemaEnv, result.FuncName, PostgresParamsMixin(result.Params))
	jsonAnswer, _ := json.Marshal(&result)
	_, _ = writer.Write(jsonAnswer)
}
