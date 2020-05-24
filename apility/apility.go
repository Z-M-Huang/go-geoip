package apility

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// Response from https://api.apility.net/geoip/
// The response is simplified version. It seems this API is a really good one to use
// However, lat and long may not return
type Response struct {
	IP struct {
		Address      string `json:"address"`
		Country      string `json:"country"`
		CountryNames struct {
			En string `json:"en"`
		} `json:"country_names"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		TimeZone  string  `json:"time_zone"`
		Region    string  `json:"region"`
		City      string  `json:"city"`
	} `json:"ip"`
}

//Get from https://api.apility.net/geoip/
func Get(host string) (*Response, error) {
	ip := net.ParseIP(host)
	if ip == nil {
		return nil, errors.New("Invalid IPAddress")
	}

	url := fmt.Sprintf("https://api.apility.net/geoip/%s", host)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %s", res.Status)
	}
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
