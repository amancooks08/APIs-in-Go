package server

import (
	"github.com/amancooks08/go_apis/StatusChecker"
)

// dependencies is a struct that contains all the dependencies for the server
type dependencies struct {
	httpchecker StatusChecker.WebsiteChecker
}

// InitRouter initializes the router with the dependencies

func InitDependencies() *dependencies {
	return &dependencies{
		httpchecker: StatusChecker.NewFunc(),
	}
}
