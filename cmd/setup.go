package main

import (
	"fmt"
	"net/http"

	"github.com/Doreen-Onyango/portfolio.git/routers"
	"github.com/Doreen-Onyango/portfolio.git/utils"
)

// Load templates and set up the router.
func Setup() (http.Handler, error) {
	if err := utils.LoadTemplates(); err != nil {
		return nil, fmt.Errorf("error loading template: %w", err)
	}
	r := http.NewServeMux()
	routers.Routes(r)

	wrappedroute := routers.RouteChecker(r)

	return wrappedroute, nil
}
