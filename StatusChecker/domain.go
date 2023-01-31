package StatusChecker

import "time"

type Website struct {
	URL      string    `json:"url"`
	Status   string    `json:"status"`
	LastPing time.Time `json:"last_ping"`
}

var websitesMap = make(map[string]Website)
