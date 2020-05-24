package geoip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreeGeoIPGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", "", "google.com"}
	for _, ip := range ipsToTest {
		resp, err := getFreeGeoIP(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IPAddress)
		assert.NotEmpty(t, resp.Country)
		assert.NotEmpty(t, resp.Latitude)
		assert.NotEmpty(t, resp.Longitude)
		//Some ips doesn't return state and city. Lattitude and Longitude will always return
	}
}

func TestFreeGeoIPGetFail(t *testing.T) {
	ipsToTest := []string{"abc", "9999.9999.9999.9999"}
	for _, ip := range ipsToTest {
		resp, err := getFreeGeoIP(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
	}
}

func TestIPAPIGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", "2001:4860:4860:0:0:0:0:6464", "google.com"}
	for _, ip := range ipsToTest {
		resp, err := getIPAPI(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IPAddress)
	}
}

func TestIPAPIInvalidIPAddress(t *testing.T) {
	ipsToTest := []string{"9999.9999.9999.9999", "2001:4860:4860:a:6464"}
	for _, ip := range ipsToTest {
		resp, err := getIPAPI(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
		assert.Equal(t, "ip-api.com invalid query", err.Error())
	}
}

func TestExtremeIPLookupGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", ""}
	for _, ip := range ipsToTest {
		resp, err := getExtremeIPLookup(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IPAddress)
		assert.NotEmpty(t, resp.Country)
		assert.NotEmpty(t, resp.Region)
		assert.NotEmpty(t, resp.City)
		assert.NotEmpty(t, resp.Latitude)
		assert.NotEmpty(t, resp.Longitude)
	}
}

func TestExtremeIPLookupNonIPV4(t *testing.T) {
	fakeIpsToTest := []string{"10.address.3.5", "10.address.3", "999.5.3.1", "fd44:e4da:5347:d30d:ffff:ffff:ffff:ffff"}
	for _, ip := range fakeIpsToTest {
		resp, err := getExtremeIPLookup(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
	}
}

func TestIPWhoisGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", "2001:4860:4860:0:0:0:0:6464"}
	for _, ip := range ipsToTest {
		resp, err := getIPWhois(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IPAddress)
	}
}

func TestIPWhoisInvalidIPAddress(t *testing.T) {
	ipsToTest := []string{"9999.9999.9999.9999", "2001:4860:4860:a:6464"}
	for _, ip := range ipsToTest {
		resp, err := getIPWhois(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
	}
}

func TestApilityGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", "2001:4860:4860:0:0:0:0:6464"}
	for _, ip := range ipsToTest {
		resp, err := getApility(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IPAddress)
	}
}

func TestApilityInvalidIPAddress(t *testing.T) {
	ipsToTest := []string{"9999.9999.9999.9999", "2001:4860:4860:a:6464"}
	for _, ip := range ipsToTest {
		resp, err := getApility(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
	}
}
