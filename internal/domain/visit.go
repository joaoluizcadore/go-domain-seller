package domain

import (
	"time"

	"github.com/joaoluizcadore/domain-seller/internal/utils"
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
		Domain:    utils.CleanDomain(domain),
		IP:        ip,
	}
}
