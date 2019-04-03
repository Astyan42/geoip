package main

import (
	"github.com/renaudcalmont/geoip/domainobjects"
	"github.com/renaudcalmont/geoip/endpoints"
	"net/http"
)

func NewRouter(businessLogic domainobjects.BusinessLogic) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/country", endpoints.GuessCountry(businessLogic))

	return router
}
