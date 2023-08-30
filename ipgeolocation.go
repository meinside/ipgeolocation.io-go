package ipgeolocation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	timeoutSeconds = 10
)

// Client for ipgeolocation API service.
type Client struct {
	apiKey     string
	httpClient *http.Client

	Verbose bool
}

// NewClient returns a new client with api key.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,

		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   timeoutSeconds * time.Second,
					KeepAlive: timeoutSeconds * time.Second,
				}).DialContext,
				IdleConnTimeout:       timeoutSeconds * time.Second,
				TLSHandshakeTimeout:   timeoutSeconds * time.Second,
				ResponseHeaderTimeout: timeoutSeconds * time.Second,
				ExpectContinueTimeout: timeoutSeconds * time.Second,
			},
		},
	}
}

// ResponseGeolocation struct
//
// https://ipgeolocation.io/documentation/ip-geolocation-api.html
type ResponseGeolocation struct {
	IP             string `json:"ip"`
	Hostname       string `json:"hostname"`
	ContinentCode  string `json:"continent_code"`
	ContinentName  string `json:"continent_name"`
	CountryCode2   string `json:"country_code2"`
	CountryCode3   string `json:"country_code3"`
	CountryName    string `json:"country_name"`
	CountryCapital string `json:"country_capital"`
	StateProvince  string `json:"state_prov"`
	District       string `json:"district"`
	City           string `json:"city"`
	Zipcode        string `json:"zipcode"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	IsEU           bool   `json:"is_eu"`
	CallingCode    string `json:"calling_code"`
	CountryTLD     string `json:"country_tld"`
	Languages      string `json:"languages"`
	CountryFlag    string `json:"country_flag"`
	GeonameID      string `json:"geoname_id"`
	ISP            string `json:"isp"`
	ConnectionType string `json:"connection_type"`
	Organization   string `json:"organization"`
	ASN            string `json:"asn"`
	Currency       struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	TimeZone struct {
		Name            string  `json:"name"`
		Offset          float32 `json:"offset"`
		CurrentTime     string  `json:"current_time"`
		CurrentTimeUnix float32 `json:"current_time_unix"`
		IsDST           bool    `json:"is_dst"`
		DSTSavings      int     `json:"dst_savings"`
	} `json:"time_zone"`
}

// GetGeolocation fetches the geolocation through API.
//
// NOTE: Pass `ip` as an empty string for querying the IP of this machine.
//
// https://ipgeolocation.io/documentation/ip-geolocation-api.html
func (c *Client) GetGeolocation(ip string) (response ResponseGeolocation, err error) {
	var baseURL *url.URL
	baseURL, err = url.Parse("https://api.ipgeolocation.io/ipgeo")

	if err == nil {
		params := url.Values{}
		params.Add("apiKey", c.apiKey)
		if len(ip) > 0 { // omit `ip` when it is empty
			params.Add("ip", ip)
		}
		baseURL.RawQuery = params.Encode()

		var resp *http.Response
		if resp, err = c.httpClient.Get(baseURL.String()); err == nil {
			if resp != nil {
				defer resp.Body.Close()
			}

			if err == nil {
				var bytes []byte
				bytes, err = io.ReadAll(resp.Body)

				if c.Verbose {
					log.Printf("[verbose] geolocation response: %s", string(bytes))
				}

				if resp.StatusCode != 200 {
					err = fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(bytes))
				} else {
					err = json.Unmarshal(bytes, &response)
					if err == nil {
						return response, nil
					}
				}
			}
		}
	}

	return ResponseGeolocation{}, err
}

// TODO: https://ipgeolocation.io/documentation/timezone-api.html

// TODO: https://ipgeolocation.io/documentation/user-agent-api.html

// TODO: https://ipgeolocation.io/documentation/astronomy-api.html
