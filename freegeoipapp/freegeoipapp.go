package freegeoipapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	countStart    time.Time
	resetDuration time.Duration
	rateLimit     int64
	rateCounter   int64
)

// Response from https://freegeoip.app/
type Response struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

// Get from https://freegeoip.app/
// Rate limit: 15,000 per hour
func Get(host string) (*Response, error) {
	if reachedLimit() {
		return nil, errors.New("Rate limit reached")
	}

	url := fmt.Sprintf("https://freegeoip.app/json/%s", host)
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
		resetDuration = 1 * time.Hour
		rateLimit = 15000
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
