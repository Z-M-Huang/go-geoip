package extremeiplookup

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	countStart    time.Time
	resetDuration time.Duration
	rateLimit     int64
	rateCounter   int64
)

// Response from https://extreme-ip-lookup.com/
type Response struct {
	BusinessName    string `json:"businessName"`
	BusinessWebsite string `json:"businessWebsite"`
	City            string `json:"city"`
	Continent       string `json:"continent"`
	Country         string `json:"country"`
	CountryCode     string `json:"countryCode"`
	IPName          string `json:"ipName"`
	IPType          string `json:"ipType"`
	Isp             string `json:"isp"`
	Lat             string `json:"lat"`
	Lon             string `json:"lon"`
	Org             string `json:"org"`
	Query           string `json:"query"`
	Region          string `json:"region"`
	Status          string `json:"status"`
}

// Get from https://extreme-ip-lookup.com/
// Rate limit: 20 per minute
func Get(ipv4 string) (*Response, error) {
	ip := net.ParseIP(ipv4)
	if ipv4 != "" && ip.To4() == nil {
		return nil, errors.New("Only take IPv4 address")
	}

	if reachedLimit() {
		return nil, errors.New("Rate limit reached")
	}

	url := fmt.Sprintf("https://extreme-ip-lookup.com/json/%s", ipv4)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Per documentation
	if res.StatusCode == http.StatusForbidden ||
		res.StatusCode == http.StatusTooManyRequests {
		rateCounter = rateLimit
		return nil, errors.New("Rate limit reached")
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %s", res.Status)
	}

	rateCounter++
	response := &Response{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %s", err.Error())
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %s", err.Error())
	}
	return response, nil
}

func reachedLimit() bool {
	if countStart.IsZero() {
		countStart = time.Now()
		resetDuration = 1 * time.Minute
		rateLimit = 20
		rateCounter = 0
		return false
	}
	if time.Now().After(countStart.Add(resetDuration)) {
		return rateCounter >= rateLimit
	}
	countStart = time.Now()
	rateCounter = 0
	return false
}
