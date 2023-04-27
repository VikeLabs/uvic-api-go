package banner

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

func (cx *BannerClient) GetData(offset, maxSize int) (*BannerResponse, error) {
	if maxSize > BANNER_PAGE_MAX || maxSize < 0 {
		panic("max size out of range: 0 <= maxSize <= 500")
	}

	u, err := url.Parse(BANNER_SSB + "searchResults/searchResults")
	if err != nil {
		return nil, err
	}

	q := map[string]string{
		"txt_term":    cx.term,
		"pageOffset":  strconv.Itoa(offset),
		"pageMaxSize": strconv.Itoa(maxSize),
	}

	setQueries(u, q)
	log.Println(u.String())

	res, err := cx.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, ErrBannerServer
	}

	var buf BannerResponse
	if err := json.NewDecoder(res.Body).Decode(&buf); err != nil {
		return nil, err
	}

	log.Println(len(buf.Data))

	if buf.Data == nil || len(buf.Data) == 0 {
		return nil, ErrBannerEmptyOffset
	}

	return &buf, nil
}
