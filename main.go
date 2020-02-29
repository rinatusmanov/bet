package main

import (
	"bet/task1"
	"fmt"
	"net/http"
)


// env
// max_pool_size=111;port=3000;endpoint=crud;host=localhost:5432;user=postgres;password=postgres_pass;schema=test;db=test
func main() {
	task1.LoadVariables()
	task1.OpenSqlConnections()
	mux := http.NewServeMux()
	fmt.Println(task1.ApiURI)
	mux.HandleFunc(task1.ApiURI, task1.Controller)
	_ = http.ListenAndServe(fmt.Sprintf(":%v", task1.PortEnv), mux)
}
