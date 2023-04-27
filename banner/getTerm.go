package banner

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func GetTerm() ([]BannerTerm, error) {
	u, err := url.Parse(BANNER_SSB + "/classSearch/getTerms")
	if err != nil {
		return nil, err
	}

	queries := make(url.Values)
	queries["searchTerms"] = []string{""}
	queries["offset"] = []string{"1"}
	queries["max"] = []string{"500"}

	u.RawQuery = queries.Encode()

	var c http.Client

	res, err := c.Get(u.String())
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, ErrBannerServer
	}

	var terms []BannerTerm
	if err := json.NewDecoder(res.Body).Decode(&terms); err != nil {
		return nil, err
	}

	return terms, nil
}
