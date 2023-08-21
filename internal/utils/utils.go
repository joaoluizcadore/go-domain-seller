package utils

import "strings"

func CleanDomain(domain string) string {
	return strings.Replace(domain, "www.", "", -1)
}
