package banner

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	BANNER_SSB      string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb"
	BANNER_PAGE_MAX uint   = 500
)

var (
	ErrBannerServer error = errors.New("BANNER request error")
)

type BannerClient struct {
	http.Client
}

// Set term and captures cookie set by Banner, used for subsequent Banner requests
func New(term *BannerTerm) (*BannerClient, error) {
	cookies, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	var cx BannerClient
	cx.Jar = cookies

	u, err := url.Parse(BANNER_SSB + "term/search")
	if err != nil {
		return nil, err
	}

	setQueries(u, map[string]string{"mode": "search"})

	response, err := cx.PostForm(u.String(), url.Values{"term": []string{term.Code}})
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, ErrBannerServer
	}

	return &cx, nil
}

func setQueries(u *url.URL, q map[string]string) {
	s := make(url.Values)
	for k, v := range q {
		s[k] = []string{v}
	}

	u.RawQuery = s.Encode()
}
