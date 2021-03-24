package main

import (
	"log"

	"github.com/apex/gateway/v2"
	"github.com/mattcanty-photography/website/internal/website"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Fatal(gateway.ListenAndServe(":3000", website.SetupRouting()))
}
