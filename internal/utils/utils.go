package utils

import "strings"

func CleanDomain(domain string) string {
	bannedWords := []string{"http://", "https://", "www."}
	for _, word := range bannedWords {
		domain = strings.Replace(domain, word, "", -1)
	}

	if strings.Contains(domain, "/") {
		domain = strings.Split(domain, "/")[0]
	}
	return domain
}
