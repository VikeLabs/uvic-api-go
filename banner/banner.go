package banner

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	bannerSsb   string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb/"
	PageMaxSize int    = 500
)

var (
	ErrEmptyOffset error = errors.New("BANNER offset empty")
)

type BannerClient struct {
	http.Client
	term string
}

// Set term and captures cookie set by Banner, used for subsequent Banner requests
func New(term string) (*BannerClient, error) {
	cookies, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	var cx BannerClient
	cx.Jar = cookies

	u, err := url.Parse(bannerSsb + "term/search")
	if err != nil {
		return nil, err
	}

	setQueries(u, map[string]string{"mode": "search"})

	response, err := cx.PostForm(u.String(), url.Values{"term": []string{term}})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("BANNER server response: %d", response.StatusCode)
	}

	cx.term = term
	return &cx, nil
}

func setQueries(u *url.URL, q map[string]string) {
	s := make(url.Values)
	for k, v := range q {
		s[k] = []string{v}
	}

	u.RawQuery = s.Encode()
}
