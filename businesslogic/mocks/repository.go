package mocks

import (
	"errors"
	"github.com/renaudcalmont/geoip/domainobjects"
	"github.com/renaudcalmont/geoip/domainobjects/models"
	"net"
)

type MockedRepository struct{}

func NewMockedRepository() domainobjects.Repository {
	return &MockedRepository{}
}

func (self *MockedRepository) RetrieveCountry(ipAddress net.IP) (models.Country, error) {
	if ipAddress.String() == "83.201.225.105" {
		return models.Country{Code: "FR"}, nil
	}

	return models.Country{}, errors.New("Not Found")
}
