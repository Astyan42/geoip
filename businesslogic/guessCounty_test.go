package businesslogic

import (
	"github.com/renaudcalmont/geoip/businesslogic/mocks"
	"testing"
)

func TestRetrieveCountryFromIPAddress(t *testing.T) {
	module := NewModule(mocks.NewMockedRepository())
	if _, err := module.RetrieveCountryFromIPAddress(""); err == nil || err.Error() != "Invalid IP address" {
		t.Errorf("Expected 'Invalid IP address' error, got '%s'", err.Error())
	}
	if _, err := module.RetrieveCountryFromIPAddress("127.0.0.1"); err == nil || err.Error() != "No match" {
		t.Errorf("Expected 'No match' error, got '%s'", err.Error())
	}
	country, err := module.RetrieveCountryFromIPAddress("83.201.225.105")
	if err != nil {
		t.Errorf("Unexpected error: '%s'", err.Error())
	}
	if country.Code != "FR" {
		t.Errorf("Expected 'FR' country code, got '%s'", country.Code)
	}
}
