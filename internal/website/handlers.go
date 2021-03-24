package website

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mattcanty-photography/website/internal/website/templates"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	templates.WritePageTemplate(w, &templates.HomePage{})
}

func getPortfolio(w http.ResponseWriter, r *http.Request) {
	portfolioID := chi.URLParam(r, "portfolioID")

	data := &templates.PortfolioPage{
		PortfolioID: portfolioID,
		Albums:      []templates.Album{},
	}

	for _, album := range getAlbums(portfolioID) {
		data.Albums = append(data.Albums, templates.Album{
			ID:           album.ID,
			CoverPhotoID: album.CoverID,
		})
	}

	templates.WritePageTemplate(w, data)
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	portfolioID := chi.URLParam(r, "portfolioID")
	albumID := chi.URLParam(r, "albumID")

	data := &templates.AlbumPage{
		PortfolioID: portfolioID,
		AlbumID:     albumID,
		Photos:      []templates.Photo{},
	}

	for _, photo := range getPhotos(portfolioID, albumID) {
		data.Photos = append(data.Photos, templates.Photo{
			ID: photo.ID,
		})
	}

	templates.WritePageTemplate(w, data)
}

func getPhoto(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.PhotoPage{
		PortfolioID: chi.URLParam(r, "portfolioID"),
		AlbumID:     chi.URLParam(r, "albumID"),
		PhotoID:     chi.URLParam(r, "photoID"),
	})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.NotFoundPage{})
}
