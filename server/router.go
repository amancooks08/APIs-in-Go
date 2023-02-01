package server

import (
	"net/http"

	"github.com/amancooks08/go_apis/StatusChecker"
)

func InitRouter(dp *dependencies) {
	http.HandleFunc("/POST/websites", StatusChecker.SubmitHandler(dp.httpchecker))
	http.HandleFunc("/GET/websites", StatusChecker.StatusHandler(dp.httpchecker))
}
