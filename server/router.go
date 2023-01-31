package server

import (
	"net/http"

	"github.com/amancooks08/go_apis/StatusChecker"
)

func InitRouter(dp *dependencies) {
	http.HandleFunc("/submit", StatusChecker.SubmitHandler(dp.httpchecker))
	http.HandleFunc("/status", StatusChecker.StatusHandler(dp.httpchecker))
}
