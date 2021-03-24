package website

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
	"github.com/mattcanty-photography/website/internal/website/templates"
	"github.com/rs/zerolog"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	templates.WritePageTemplate(w, &templates.HomePage{})
}

func portfolioPageHandler(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	portfolioID := chi.URLParam(r, "portfolioID")

	data := &templates.PortfolioPage{
		PortfolioID: portfolioID,
		Albums:      []templates.Album{},
	}

	portfolio, err := getPortfolio(portfolioID)
	if err != nil {
		errorHandler(w, oplog, err)
		return
	}

	for _, album := range portfolio.Albums {
		data.Albums = append(data.Albums, templates.Album{
			ID:           album.ID,
			CoverPhotoID: album.CoverID,
		})
	}

	templates.WritePageTemplate(w, data)
}

func albumPageHandler(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	portfolioID := chi.URLParam(r, "portfolioID")
	albumID := chi.URLParam(r, "albumID")

	data := &templates.AlbumPage{
		PortfolioID: portfolioID,
		AlbumID:     albumID,
		Photos:      []templates.Photo{},
	}

	album, err := getAlbum(portfolioID, albumID)
	if err != nil {
		errorHandler(w, oplog, err)
		return
	}

	for _, photo := range album.Photos {
		data.Photos = append(data.Photos, templates.Photo{
			ID: photo.ID,
		})
	}

	templates.WritePageTemplate(w, data)
}

func photoPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.PhotoPage{
		PortfolioID: chi.URLParam(r, "portfolioID"),
		AlbumID:     chi.URLParam(r, "albumID"),
		PhotoID:     chi.URLParam(r, "photoID"),
	})
}

func notFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.ErrorPage{
		Heading: "Not Found",
		Message: "Either you are lost, or I have done something wrong. Have a nice day :-)",
	})
}

func errorHandler(w http.ResponseWriter, oplog zerolog.Logger, err error) {
	oplog.Error().Msg(err.Error())
	templates.WritePageTemplate(w, &templates.ErrorPage{
		Heading: "Error",
		Message: "Something went really wrong. Let me know, if you know me. :-(",
	})
}
