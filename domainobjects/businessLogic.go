package domainobjects

import "geoip/domainobjects/models"

type BusinessLogic interface {
	RetrieveCountryFromIPAddress(ipAddress string) (models.Country, error)
}
