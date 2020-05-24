//Package geoip library is powered by a number of free API services
//Author: Z-M-Huang
//Repository: https://github.com/Z-M-Huang/go-geoip
package geoip

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Z-M-Huang/go-extremeiplookup"
	"github.com/Z-M-Huang/go-freegeoip"
	"github.com/Z-M-Huang/go-geoip/apility"
	"github.com/Z-M-Huang/go-geoip/ipwhois"
	"github.com/Z-M-Huang/go-ipapi"
)

//GetLocation gets the location info from host.
//host can be ipv4 ipv6 or domain host name
func GetLocation(host string) (*Location, error) {
	loc, err := getFreeGeoIP(host)
	if err == nil && loc != nil {
		return loc, nil
	}
	loc, err = getIPAPI(host)
	if err == nil && loc != nil {
		return loc, nil
	}
	loc, err = getExtremeIPLookup(host)
	if err == nil && loc != nil {
		return loc, nil
	}
	loc, err = getIPWhois(host)
	if err == nil && loc != nil {
		return loc, nil
	}
	loc, err = getApility(host)
	if err == nil && loc != nil {
		return loc, nil
	}
	return nil, errors.New("no service can identify the host name")
}

func getFreeGeoIP(host string) (*Location, error) {
	resp, err := freegeoip.Get(host)
	if err != nil {
		return nil, fmt.Errorf("freegeoip.app %s", err.Error())
	}
	loc := &Location{}
	loc.IPAddress = resp.IP
	loc.Country = resp.CountryName
	loc.Region = resp.RegionName
	loc.City = resp.City
	loc.ZipCode = resp.ZipCode
	loc.Latitude = resp.Latitude
	loc.Longitude = resp.Longitude
	return loc, nil
}

func getIPAPI(host string) (*Location, error) {
	resp, err := ipapi.Get(host)
	if err != nil {
		return nil, fmt.Errorf("ip-api.com %s", err.Error())
	}
	loc := &Location{}
	loc.IPAddress = resp.Query
	loc.Country = resp.Country
	loc.Region = resp.RegionName
	loc.City = resp.City
	loc.ZipCode = resp.Zip
	loc.Latitude = resp.Lat
	loc.Longitude = resp.Lon
	return loc, nil
}

func getExtremeIPLookup(host string) (*Location, error) {
	resp, err := extremeiplookup.Get(host)
	if err != nil {
		return nil, fmt.Errorf("extreme-ip-lookup %s", err.Error())
	}
	loc := &Location{}
	loc.IPAddress = resp.Query
	loc.Country = resp.Country
	loc.Region = resp.Region
	loc.City = resp.City
	loc.ZipCode = ""
	lat, err := strconv.ParseFloat(resp.Lat, 64)
	if err != nil {
		return nil, errors.New("extreme-ip-lookup lat is not float64")
	}
	loc.Latitude = lat
	lon, err := strconv.ParseFloat(resp.Lon, 64)
	if err != nil {
		return nil, errors.New("extreme-ip-lookup lon is not float64")
	}
	loc.Longitude = lon
	return loc, nil
}

func getIPWhois(host string) (*Location, error) {
	resp, err := ipwhois.Get(host)
	if err != nil {
		return nil, fmt.Errorf("ipwhoise.io %s", err.Error())
	}
	loc := &Location{}
	loc.IPAddress = resp.IP
	loc.Country = resp.Country
	loc.Region = resp.Region
	loc.City = resp.City
	loc.ZipCode = ""
	lat, err := strconv.ParseFloat(resp.Latitude, 64)
	if err != nil {
		return nil, errors.New("ipwhoise.io lat is not float64")
	}
	loc.Latitude = lat
	lon, err := strconv.ParseFloat(resp.Longitude, 64)
	if err != nil {
		return nil, errors.New("ipwhoise.io lon is not float64")
	}
	loc.Longitude = lon
	return loc, nil
}

func getApility(host string) (*Location, error) {
	resp, err := apility.Get(host)
	if err != nil {
		return nil, fmt.Errorf("apility %s", err.Error())
	}
	loc := &Location{}
	loc.IPAddress = resp.IP.Address
	loc.Country = resp.IP.CountryNames.En
	loc.Region = resp.IP.Region
	loc.City = resp.IP.City
	loc.ZipCode = ""
	//lat and long are not always returned
	loc.Latitude = resp.IP.Latitude
	loc.Longitude = resp.IP.Longitude
	return loc, nil
}
