package extremeiplookup

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", ""}
	for _, ip := range ipsToTest {
		resp, err := Get(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.Query)
		assert.NotEmpty(t, resp.Country)
		assert.NotEmpty(t, resp.Region)
		assert.NotEmpty(t, resp.City)
		assert.NotEmpty(t, resp.Lat)
		assert.NotEmpty(t, resp.Lon)
	}
}

func TestNonIPV4(t *testing.T) {
	fakeIpsToTest := []string{"10.address.3.5", "10.address.3", "999.5.3.1", "fd44:e4da:5347:d30d:ffff:ffff:ffff:ffff"}
	for _, ip := range fakeIpsToTest {
		resp, err := Get(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
		assert.Equal(t, "Only take IPv4 address", err.Error())
	}
}

func TestRateLimit(t *testing.T) {
	countStart = time.Now().Add(-10 * time.Hour)
	resetDuration = 1 * time.Hour
	rateCounter = 11
	rateLimit = 10
	resp, err := Get("")
	assert.Empty(t, resp)
	assert.Equal(t, "Rate limit reached", err.Error())
}
