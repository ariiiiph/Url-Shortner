package utils

import "os"

func IsDiffrentDomain(url string) bool {
	domain := os.Getenv("DOMAIN")

	if url == domain {
		return false
	}
}
