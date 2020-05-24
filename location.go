package geoip

//Location information from ip addresses
type Location struct {
	IPAddress string  `json:"ipAddress" xml:"ipAddress" yaml:"ipAddress"`
	Country   string  `json:"country" xml:"country" yaml:"country"`
	Region    string  `json:"region" xml:"region" yaml:"region"`
	City      string  `json:"city" xml:"city" yaml:"city"`
	ZipCode   string  `json:"zipCode" xml:"zipCode" yaml:"zipCode"`
	Latitude  float64 `json:"latitude" xml:"latitude" yaml:"latitude"`
	Longitude float64 `json:"longitude" xml:"longitude" yaml:"longitude"`
}
