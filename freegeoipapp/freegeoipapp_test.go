package freegeoipapp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", "", "google.com"}
	for _, ip := range ipsToTest {
		resp, err := Get(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IP)
		assert.NotEmpty(t, resp.CountryName)
		assert.NotEmpty(t, resp.Latitude)
		assert.NotEmpty(t, resp.Longitude)
		//Some ips doesn't return state and city. Lattitude and Longitude will always return
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
