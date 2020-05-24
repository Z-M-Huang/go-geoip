package ipwhois

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

//Response from http://free.ipwhois.io/json/
type Response struct {
	IP                string `json:"ip"`
	Success           bool   `json:"success"`
	Type              string `json:"type"`
	Continent         string `json:"continent"`
	ContinentCode     string `json:"continent_code"`
	Country           string `json:"country"`
	CountryCode       string `json:"country_code"`
	CountryFlag       string `json:"country_flag"`
	CountryCapital    string `json:"country_capital"`
	CountryPhone      string `json:"country_phone"`
	CountryNeighbours string `json:"country_neighbours"`
	Region            string `json:"region"`
	City              string `json:"city"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Asn               string `json:"asn"`
	Org               string `json:"org"`
	Isp               string `json:"isp"`
	Timezone          string `json:"timezone"`
	TimezoneName      string `json:"timezone_name"`
	TimezoneDstOffset string `json:"timezone_dstOffset"`
	TimezoneGmtOffset string `json:"timezone_gmtOffset"`
	TimezoneGmt       string `json:"timezone_gmt"`
	Currency          string `json:"currency"`
	CurrencyCode      string `json:"currency_code"`
	CurrencySymbol    string `json:"currency_symbol"`
	CurrencyRates     string `json:"currency_rates"`
	CurrencyPlural    string `json:"currency_plural"`
	CompletedRequests int    `json:"completed_requests"`
}

// Get from http://free.ipwhois.io/json/
// Rate limit: 10, 000 per month
func Get(ipAddress string) (*Response, error) {
	ip := net.ParseIP(ipAddress)
	if ipAddress != "" && ip.To4() == nil && ip.To16() == nil {
		return nil, errors.New("Invalid IPAddress")
	}

	if reachedLimit() {
		return nil, errors.New("Rate limit reached")
	}

	url := fmt.Sprintf("http://free.ipwhois.io/json/%s", ipAddress)
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
		resetDuration = 31 * 24 * 60 * time.Hour
		rateLimit = 10000
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
