package usecases

import (
	"geoip/domainobjects/models"
	"net"
)

type RepositoryForGuessCountry interface {
	RetrieveCountry(ipAddress net.IP) (models.Country, error)
}

type BusinessLogicForGuessCountry interface {
	RetrieveCountryFromIPAddress(ipAddress string) (models.Country, error)
}
