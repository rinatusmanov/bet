package main

import (
	"bet/task1"
	"fmt"
	"net/http"
)

func main() {
	task1.LoadVariables()
	task1.OpenSqlConnections()
	mux := http.NewServeMux()
	mux.HandleFunc(fmt.Sprintf("/%v/v%v", task1.EndpointEnv, task1.Version), task1.Controller)
	_ = http.ListenAndServe(fmt.Sprintf(":%v", task1.PortEnv), mux)
}
