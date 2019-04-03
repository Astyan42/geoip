package domainobjects

import (
	"github.com/renaudcalmont/geoip/domainobjects/models"
	"net"
)

type Repository interface {
	RetrieveCountry(ipAddress net.IP) (models.Country, error)
}
