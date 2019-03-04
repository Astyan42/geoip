package main

import (
	"flag"
	"fmt"
	"geoip/businesslogic"
	"geoip/repository"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 5000, "TCP port the service will listen on")
	flag.Parse()

	repository := repository.NewModule("data/GeoLite2-Country.mmdb")
	businessLogic := businesslogic.NewModule(repository)
	router := NewRouter(businessLogic)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
