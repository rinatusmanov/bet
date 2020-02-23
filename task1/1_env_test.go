package task1

import (
	"fmt"
	"os"
	"testing"
)

type testLoadVariables struct {
	MaxPoolSize int64
	Port        int64
	Endpoint    string
	Host        string
	User        string
	Password    string
	Schema      string
	DB          string
}

var testLoadVariablesCases = []testLoadVariables{
	testLoadVariables{
		MaxPoolSize: 1,
		Port:        2,
		Endpoint:    "EndpointEnv 1",
		Host:        "HostEnv 1",
		User:        "UserEnv 1",
		Password:    "PasswordEnv 1",
		Schema:      "SchemaEnv 1",
		DB:          "DBEnv 1",
	},
	testLoadVariables{
		MaxPoolSize: 112454,
		Port:        2,
		Endpoint:    "EndpointEnv 2",
		Host:        "HostEnv 2",
		User:        "UserEnv 2",
		Password:    "PasswordEnv 2",
		Schema:      "SchemaEnv 2",
		DB:          "DBEnv 2",
	},
}

func TestLoadVariables(t *testing.T) {
	for _, testCase := range testLoadVariablesCases {
		errMaxPoolSize := os.Setenv("max_pool_size", fmt.Sprintf("%v", testCase.MaxPoolSize))
		errPort := os.Setenv("port", fmt.Sprintf("%v", testCase.Port))
		errEndpoint := os.Setenv("endpoint", fmt.Sprintf("%v", testCase.Endpoint))
		errHost := os.Setenv("host", fmt.Sprintf("%v", testCase.Host))
		errUser := os.Setenv("user", fmt.Sprintf("%v", testCase.User))
		errPassword := os.Setenv("password", fmt.Sprintf("%v", testCase.Password))
		errSchema := os.Setenv("schema", fmt.Sprintf("%v", testCase.Schema))
		errDB := os.Setenv("db", fmt.Sprintf("%v", testCase.DB))
		if errMaxPoolSize != nil {
			panic(errMaxPoolSize)
		}
		if errPort != nil {
			panic(errPort)
		}
		if errEndpoint != nil {
			panic(errEndpoint)
		}
		if errHost != nil {
			panic(errHost)
		}
		if errUser != nil {
			panic(errUser)
		}
		if errPassword != nil {
			panic(errPassword)
		}
		if errSchema != nil {
			panic(errSchema)
		}
		if errDB != nil {
			panic(errDB)
		}
		LoadVariables()
		if MaxPoolSizeEnv != testCase.MaxPoolSize {
			t.Errorf("Не соответствие MaxPoolSizeEnv из тесткейса %v к полученному %v", testCase.MaxPoolSize, MaxPoolSizeEnv)
		}
		if PortEnv != testCase.Port {
			t.Errorf("Не соответствие PortEnv из тесткейса %v к полученному %v", testCase.Port, PortEnv)
		}
		if EndpointEnv != testCase.Endpoint {
			t.Errorf("Не соответствие EndpointEnv из тесткейса %v к полученному %v", testCase.Endpoint, EndpointEnv)
		}
		if HostEnv != testCase.Host {
			t.Errorf("Не соответствие HostEnv из тесткейса %v к полученному %v", testCase.Host, HostEnv)
		}
		if UserEnv != testCase.User {
			t.Errorf("Не соответствие UserEnv из тесткейса %v к полученному %v", testCase.User, UserEnv)
		}
		if PasswordEnv != testCase.Password {
			t.Errorf("Не соответствие PasswordEnv из тесткейса %v к полученному %v", testCase.Password, PasswordEnv)
		}
		if SchemaEnv != testCase.Schema {
			t.Errorf("Не соответствие SchemaEnv из тесткейса %v к полученному %v", testCase.Schema, SchemaEnv)
		}
		if DBEnv != testCase.DB {
			t.Errorf("Не соответствие DBEnv из тесткейса %v к полученному %v", testCase.DB, DBEnv)
		}
	}
}
