package main

import (
	"net/http"

	"github.com/amancooks08/go_apis/StatusChecker"
	"github.com/amancooks08/go_apis/server"
)

func main() {
	dep := server.InitDependencies()
	server.InitRouter(dep)
	go StatusChecker.MonitorWebsite()
	http.ListenAndServe(":8000", nil)
}
