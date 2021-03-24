package main

import (
	"log"
	"net/http"

	"github.com/mattcanty-photography/website/internal/website"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Fatal(http.ListenAndServe(":3000", website.SetupRouting()))
}
