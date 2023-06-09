package banner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetTerm() ([]BannerTerm, error) {
	u, err := url.Parse(bannerSsb + "/classSearch/getTerms")
	if err != nil {
		return nil, err
	}

	query := map[string]string{
		"searchTerms": "",
		"offset":      "1",
		"max":         "500",
	}

	setQueries(u, query)

	var c http.Client

	res, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("banner server status: %d", res.StatusCode)
	}

	var terms []BannerTerm
	if err := json.NewDecoder(res.Body).Decode(&terms); err != nil {
		return nil, err
	}

	return terms, nil
}
