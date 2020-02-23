package task1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testControllerCase struct {
	name       string
	uri        string
	method     string
	body       string
	answerCode int
	answerBody controllerAnswer
}

func TestController(t *testing.T) {
	EndpointEnv = "test_endpoint"
	SchemaEnv = "test"
	ApiURI = fmt.Sprintf("/%v/v%v/", EndpointEnv, Version)
	var (
		ts        = httptest.NewServer(http.HandlerFunc(Controller))
		testCases = []testControllerCase{
			testControllerCase{
				name:       "1 тест get",
				uri:        "user/1/comment/2",
				body:       "",
				method:     "GET",
				answerCode: 200,
				answerBody: controllerAnswer{
					FuncName: "SELECT * FROM test.user_comment_get($1,$2);",
					Params:   []interface{}{1, 2},
					Body:     "",
				},
			},
			testControllerCase{
				name:       "2 тест post",
				uri:        "henry/1/qwerty/2/kinder",
				body:       `{"test":2,"test_2":1}`,
				method:     "POST",
				answerCode: 200,
				answerBody: controllerAnswer{
					FuncName: "SELECT * FROM test.user_ins($1,$2,$3,$4);",
					Params:   []interface{}{1, 2, "test", 2, "test_2", 1},
					Body:     `{"test":2,"test_2":1}`,
				},
			},
		}
		answerBody controllerAnswer
	)
	defer ts.Close()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r, errR := http.NewRequest(testCase.method, ts.URL+"/"+testCase.uri, strings.NewReader(testCase.body))
			if errR != nil {
				t.Fatal(errR)
			}
			resp, err := http.DefaultClient.Do(r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			byteSl, _ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(byteSl, &answerBody)
			if answerBody.Body != testCase.answerBody.Body {
				t.Fatalf("ошибка теста # %v несовпадение body, должно быть %v получено %v", testCase.name, testCase.answerBody.Body, answerBody.Body)
			}
		})
	}
}

type testPostgresParamsMixinCase struct {
	inData  []interface{}
	outData string
}

func TestPostgresParamsMixin(t *testing.T) {
	var testCases = []testPostgresParamsMixinCase{
		testPostgresParamsMixinCase{
			inData:  []interface{}{1, 2, 3, 4, 51, 2, 3},
			outData: "$1,$2,$3,$4,$5,$6,$7",
		},
		testPostgresParamsMixinCase{
			inData:  []interface{}{"1", 2, "3", 4, "51", 2, 3},
			outData: "$1,$2,$3,$4,$5,$6,$7",
		},
		testPostgresParamsMixinCase{
			inData:  []interface{}{"1", 2},
			outData: "$1,$2",
		},
		testPostgresParamsMixinCase{
			inData:  []interface{}{},
			outData: "",
		},
	}
	for index, testCase := range testCases {
		var outData = PostgresParamsMixin(testCase.inData)
		if testCase.outData != outData {
			t.Fatalf("тест # %v, не прошел, ожидалось %v, получилось %v.", index, testCase.outData, outData)
		}
	}
}

type testToPostgresFuncCase struct {
	uri      string
	method   string
	funcName string
	id       []interface{}
}

func TestToPostgresFunc(t *testing.T) {
	var testCases = []testToPostgresFuncCase{
		testToPostgresFuncCase{
			uri:      "user/1/comment/1000/kinder",
			method:   "POST",
			funcName: "user_comment_kinder_ins",
			id:       []interface{}{1, 1000},
		},
	}
	for index, testCase := range testCases {
		funcName, id := ToPostgresFunc(testCase.uri, testCase.method)
		if funcName != testCase.funcName {
			t.Fatalf("тест # %v, не прошел, funcName ожидалось %v, получилось %v.", index, testCase.funcName, funcName)
		}
		if len(id) != len(testCase.id) {
			t.Fatalf("тест # %v, не прошел, id ожидалось %v, получилось %v.", index, testCase.id, id)
		} else {
			for i := 0; i < len(id); i++ {
				if id[i] != testCase.id[i] {
					t.Fatalf("тест # %v, не прошел, id ожидалось %v, получилось %v.", index, testCase.id, id)
				}
			}
		}
	}
}
