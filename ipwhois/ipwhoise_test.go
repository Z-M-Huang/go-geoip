package ipwhois

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ipsToTest := []string{"8.8.8.8", "2001:4860:4860:0:0:0:0:6464"}
	for _, ip := range ipsToTest {
		resp, err := Get(ip)
		assert.Empty(t, err)
		assert.NotEmpty(t, resp)
		assert.NotEmpty(t, resp.IP)
	}
}

func TestInvalidIPAddress(t *testing.T) {
	ipsToTest := []string{"9999.9999.9999.9999", "2001:4860:4860:a:6464"}
	for _, ip := range ipsToTest {
		resp, err := Get(ip)
		assert.Empty(t, resp)
		assert.NotEmpty(t, err)
		assert.Equal(t, "Invalid IPAddress", err.Error())
	}
}

func TestRateLimit(t *testing.T) {
	countStart = time.Now()
	resetDuration = 1 * time.Hour
	rateCounter = 11
	rateLimit = 10
	resp, err := Get("")
	assert.Empty(t, resp)
	assert.Equal(t, "Rate limit reached", err.Error())
}
