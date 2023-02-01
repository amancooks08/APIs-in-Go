package StatusChecker

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type WebsiteChecker interface {
	CheckStatus(ctx context.Context, url string) (status string, err error)
}

func NewFunc() WebsiteChecker {
	return &Website{}
}

func (w Website) CheckStatus(ctx context.Context, url string) (status string, err error) {

	if _, ok := websitesMap[url]; !ok {
		return "Key not present.", errors.New("Website Not Found")
	}

	return websitesMap[url].Status, nil
}

func MonitorWebsite() {
	for {
		for key := range websitesMap {
			res, err := http.Get("http://" + key)
			temp := websitesMap[key]
			if res.StatusCode != http.StatusOK || err != nil {
				temp.Status = "DOWN"
				websitesMap[key] = temp
			} else {
				temp.Status = "UP"
				websitesMap[key] = temp
			}

		}
		time.Sleep(time.Minute)
	}
}
