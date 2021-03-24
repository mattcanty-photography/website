package website

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httplog"
)

type RoutingConfiguration struct {
	RoutePatternPrefix string
	JsonLogging        bool
}

func SetupRouting(config RoutingConfiguration) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("httplog", httplog.Options{
		JSON: config.JsonLogging,
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(9, "text/html"))

	r.HandleFunc(fmt.Sprintf("%s/", config.RoutePatternPrefix), homePageHandler)
	r.HandleFunc(fmt.Sprintf("%s/portfolio/{portfolioID}", config.RoutePatternPrefix), portfolioPageHandler)
	r.HandleFunc(fmt.Sprintf("%s/portfolio/{portfolioID}/album/{albumID}", config.RoutePatternPrefix), albumPageHandler)
	r.HandleFunc(fmt.Sprintf("%s/portfolio/{portfolioID}/album/{albumID}/photo/{photoID}", config.RoutePatternPrefix), photoPageHandler)

	r.NotFound(notFoundPageHandler)

	return r
}
