package ip_country

import (
	"encoding/json"
	"net/http"
)

type IpCountryAdapter struct{}

func NewIpCountryCodeAdapter() *IpCountryAdapter {
	return &IpCountryAdapter{}
}

type IpCountryResponse struct {
	Country string `json:"country"`
}

func (a *IpCountryAdapter) GetCountryCodeFromIp(ip string) (string, error) {
	res, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var response IpCountryResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Country, nil
}
