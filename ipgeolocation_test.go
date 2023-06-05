package ipgeolocation

import (
	"os"
	"testing"
)

const (
	ip           = "8.8.8.8"
	countryCode2 = "US"
)

func TestGetGeolocation(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	verbose := os.Getenv("VERBOSE") == "true"

	if len(apiKey) != 0 {
		client := NewClient(apiKey)
		client.Verbose = verbose

		// query an IP address
		if result, err := client.GetGeolocation(ip); err != nil {
			t.Errorf("Failed to get geolocation: %s", err)
		} else {
			if result.CountryCode2 != countryCode2 {
				t.Errorf("Expected result country to be `%s`, got `%s`.", countryCode2, result.CountryCode2)
			}
		}
	} else {
		t.Errorf("Environment variable `API_KEY` was not found.")
	}
}
