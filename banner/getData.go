package banner

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func (cx *BannerClient) GetData(offset, maxSize int) (*BannerResponse, error) {
	if maxSize > BANNER_PAGE_MAX || maxSize < 0 {
		fmt.Println("max size out of range: 0 <= maxSize <= 500")
		os.Exit(1)
	}

	u, err := url.Parse(bannerSsb + "searchResults/searchResults")
	if err != nil {
		return nil, err
	}

	q := map[string]string{
		"txt_term":    cx.term,
		"pageOffset":  strconv.Itoa(offset),
		"pageMaxSize": strconv.Itoa(maxSize),
	}

	setQueries(u, q)

	res, err := cx.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("banner response status: %d", res.StatusCode)
	}

	var buf BannerResponse
	if err := json.NewDecoder(res.Body).Decode(&buf); err != nil {
		return nil, err
	}

	if buf.Data == nil || len(buf.Data) == 0 {
		return nil, ErrBannerEmptyOffset
	}

	return &buf, nil
}
