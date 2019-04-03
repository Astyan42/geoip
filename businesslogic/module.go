package businesslogic

import (
	"errors"
	"github.com/renaudcalmont/geoip/domainobjects"
	"github.com/renaudcalmont/geoip/domainobjects/models"
	"net"
)

type Module struct {
	repository domainobjects.Repository
}

func NewModule(repository domainobjects.Repository) domainobjects.BusinessLogic {
	return &Module{repository}
}

func (self *Module) RetrieveCountryFromIPAddress(ipAddress string) (models.Country, error) {
	parsedIPAddress := net.ParseIP(ipAddress)
	if parsedIPAddress == nil {
		return models.Country{}, errors.New("Invalid IP address")
	}

	country, err := self.repository.RetrieveCountry(parsedIPAddress)
	if err != nil {
		return models.Country{}, errors.New("No match")
	}

	return country, nil
}
