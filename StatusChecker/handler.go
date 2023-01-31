package StatusChecker

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func SubmitHandler(httpchecker WebsiteChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			newWebsite := make(map[string][]string)
			json.NewDecoder(r.Body).Decode(&newWebsite)
			weblist := newWebsite["websites"]
			for _, v := range weblist {
				websitesMap[v] = Website{v, "Not Ready", time.Now()}
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}

func StatusHandler(httpchecker WebsiteChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var status string
			url := r.URL.Query().Get("id")
			statusMap := make(map[string]string)
			if url != "" {
				status, _ = httpchecker.CheckStatus(context.Background(), url)
				statusMap[url] = status
			} else {
				for url := range websitesMap {
					statusMap[url], _ = httpchecker.CheckStatus(context.Background(), url)
				}
			}
			json.NewEncoder(w).Encode(statusMap)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
