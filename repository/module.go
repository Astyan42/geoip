package repository

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"geoip/domainobjects"
	"geoip/domainobjects/models"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/oschwald/maxminddb-golang"
)

type Module struct {
	countriesDatabaseFile string
}

func NewModule(countriesDatabaseFile string) domainobjects.Repository {
	module := &Module{countriesDatabaseFile}
	module.Initialize()
	return module
}

func (self *Module) RetrieveCountry(ipAddress net.IP) (models.Country, error) {
	db, err := maxminddb.Open(self.countriesDatabaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}

	if err = db.Lookup(ipAddress, &record); err != nil {
		log.Fatal(err)
	}
	if record.Country.ISOCode == "" {
		return models.Country{}, errors.New("Not Found")
	}

	return models.Country{
		Code: record.Country.ISOCode}, nil
}

func (self *Module) Initialize() {
	dir, _ := path.Split(self.countriesDatabaseFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(self.countriesDatabaseFile); os.IsNotExist(err) {
		downloadMMDB(
			self.countriesDatabaseFile,
			"https://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.tar.gz")
	}
}

func downloadMMDB(filepath string, url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("DownloadFile: %s", err.Error())
	}
	defer response.Body.Close()

	uncompressedStream, err := gzip.NewReader(response.Body)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			// if err := os.Mkdir(header.Name, 0755); err != nil {
			// 	log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			// }
		case tar.TypeReg:
			if strings.HasSuffix(header.Name, "mmdb") {
				outFile, err := os.Create(filepath)
				if err != nil {
					log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
				}
				defer outFile.Close()
				if _, err := io.Copy(outFile, tarReader); err != nil {
					log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
				}
			}
		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %v in %s",
				header.Typeflag,
				header.Name)
		}
	}
}
