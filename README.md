# go-geoip
Free Geoip API
This API is a sum of free geo apis available to public. Thank you all for the awesome api.
This library is mainly for get the latitude and longitude. Please fork and modify based on your needs.

[![Build Status](https://travis-ci.com/Z-M-Huang/go-geoip.svg?branch=master)](https://travis-ci.com/Z-M-Huang/go-geoip)

# Example usage
```
  import "github.com/Z-M-Huang/go-geoip"

  resp, err := geoip.GetLocation(host)
  if err != nil {
    panic(err)
  }
  fmt.Println(resp.IPAddress)
```