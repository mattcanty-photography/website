package main

import (
	"log"
	"net/http"

	"github.com/mattcanty-photography/website/internal/website"
)

func main() {
	config := website.RoutingConfiguration{
		RoutePatternPrefix: "",
		JsonLogging:        false,
	}

	log.Fatal(http.ListenAndServe(":3000", website.SetupRouting(config)))
}
