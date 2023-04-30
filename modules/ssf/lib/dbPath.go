package lib

import (
	"os"
	"strings"
)

func GetDSN() string {
	p := []string{"modules", "ssf", "database.db"}
	path := strings.Join(p, string(os.PathSeparator))
	return path
}
