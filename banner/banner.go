package banner

import (
	"errors"
)

const (
	BANNER_SSB string = "https://banner.uvic.ca/StudentRegistrationSsb/ssb"
)

var (
	ErrBannerServer error = errors.New("ErrBanner")
)
