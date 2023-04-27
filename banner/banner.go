package banner

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	BANNER_SSB string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb"
)

var (
	ErrBannerServer error = errors.New("ErrBanner")
)

type BannerClient http.Client

func New(term BannerTerm) (*BannerClient, error) {

	return nil, nil
}

func setQueries(u *url.URL, q map[string]string) {
	s := make(url.Values)
	for k, v := range q {
		s[k] = []string{v}
	}

	u.RawQuery = s.Encode()
}
