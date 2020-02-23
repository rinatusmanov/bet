package task1

import (
	"fmt"
	"os"
	"strconv"
)

var (
	MaxPoolSizeEnv int64
	PortEnv        int64
	EndpointEnv    string
	HostEnv        string
	UserEnv        string
	PasswordEnv    string
	SchemaEnv      string
	DBEnv          string
)

func LoadVariables() {
	MaxPoolSizeEnv, _ = strconv.ParseInt(os.Getenv("max_pool_size"), 10, 64)
	PortEnv, _ = strconv.ParseInt(os.Getenv("port"), 10, 64)
	EndpointEnv = os.Getenv("endpoint")
	HostEnv = os.Getenv("host")
	UserEnv = os.Getenv("user")
	PasswordEnv = os.Getenv("password")
	SchemaEnv = os.Getenv("schema")
	DBEnv = os.Getenv("db")
}

var (
	apiURI string
)

func init() {
	apiURI = fmt.Sprintf("/%v/v%v/", EndpointEnv, Version)
}
