package domainobjects

import (
	"geoip/domainobjects/models"
	"net"
)

type Repository interface {
	RetrieveCountry(ipAddress net.IP) (models.Country, error)
}
