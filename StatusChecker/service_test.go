package StatusChecker

import (
	"context"
	"testing"
	"time"
	"sync"
	"github.com/stretchr/testify/suite"
)

var wg sync.WaitGroup
type WebsiteSuite struct {
	suite.Suite
}

func (s *WebsiteSuite) SetupTest() {
	websitesMap = make(map[string]Website)
	weblist := []string{"www.google.com", "www.facebook.com", "www.amazon.com", "www.fakebook.com"}
	for _, v := range weblist {
		websitesMap[v] = Website{v, "Not Ready", time.Now()}
	}
}

func (s *WebsiteSuite) TestCheckStatus() {
	w := Website{}

	// Test with a valid website
	status, err := w.CheckStatus(context.Background(), "www.google.com")
	s.NoError(err)
	s.Equal("Not Ready", status)

	// Test with an invalid website
	status, err = w.CheckStatus(context.Background(), "www.invalid.com")
	s.Error(err)
	s.Equal("Key not present.", status)
}

func (s *WebsiteSuite) TestMonitorWebsite() {
	// Test if the status of the websites is being updated correctly
	wg.Add(1)
	go func() {
		defer wg.Done()
		MonitorWebsite()
	}()

	wg.Wait()
	time.Sleep(time.Second * 2)
	s.Equal("UP", websitesMap["www.google.com"].Status)
	s.Equal("DOWN", websitesMap["www.fakebook.com"].Status)
	s.Equal("UP", websitesMap["www.amazon.com"].Status)
}

func TestWebsiteSuite(t *testing.T) {
	suite.Run(t, new(WebsiteSuite))
}
