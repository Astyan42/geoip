package main

import (
	"geoip/domainobjects"
	"geoip/endpoints"
	"net/http"
)

func NewRouter(businessLogic domainobjects.BusinessLogic) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/country", endpoints.GuessCountry(businessLogic))

	return router
}
