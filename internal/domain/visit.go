package domain

import (
	"strings"
	"time"
)

type Visit struct {
	Id        int64      `json:"id"`
	VisitDate *time.Time `json:"visit_date"`
	Domain    string     `json:"domain"`
	IP        string     `json:"ip"`
}

func NewVisit(domain string, ip string, visitDate *time.Time) *Visit {

	if visitDate == nil {
		d := time.Now()
		visitDate = &d
	}

	return &Visit{
		VisitDate: visitDate,
		Domain:    cleanDomain(domain),
		IP:        ip,
	}
}

func cleanDomain(domain string) string {
	return strings.Replace(domain, "www.", "", -1)
}
