package endpoints

import (
	"encoding/json"
	"github.com/renaudcalmont/geoip/domainobjects"
	"net/http"
)

func GuessCountry(businessLogic domainobjects.BusinessLogic) http.HandlerFunc {
	type requestBody struct {
		Address string
	}

	type responseBody struct {
		Code string `json:"code"`
	}

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			var requestBody requestBody
			if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
				http.Error(responseWriter, err.Error(), http.StatusBadRequest)
				return
			}

			country, err := businessLogic.RetrieveCountryFromIPAddress(requestBody.Address)
			if err != nil {
				http.Error(responseWriter, err.Error(), http.StatusNotFound)
				return
			}

			responseWriter.Header().Set("Content-Type", "application/json")
			json.NewEncoder(responseWriter).Encode(responseBody{country.Code})

		default:
			http.Error(responseWriter, "Invalid method", http.StatusBadRequest)
		}
	}
}
