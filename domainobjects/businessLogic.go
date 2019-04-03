package domainobjects

import "github.com/renaudcalmont/geoip/domainobjects/models"

type BusinessLogic interface {
	RetrieveCountryFromIPAddress(ipAddress string) (models.Country, error)
}
