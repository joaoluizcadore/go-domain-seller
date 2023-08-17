package database

import (
	"database/sql"

	"github.com/joaoluizcadore/domain-seller/internal/domain"
	"github.com/joaoluizcadore/domain-seller/internal/shared"
)

type VisitSqlRepository struct {
	Db *sql.DB
}

func NewVisitCountSqlRepository(db *sql.DB) *VisitSqlRepository {
	return &VisitSqlRepository{Db: db}
}

func (r *VisitSqlRepository) Save(visit *domain.Visit) error {
	_, err := r.Db.Exec(`INSERT INTO visits (visit_date, domain, ip) VALUES ($1, $2, $3)`,
		visit.VisitDate, visit.Domain, visit.IP)

	return err
}

func (r *VisitSqlRepository) GetListSummary() ([]shared.VisitCountSummaryResponse, error) {
	rows, err := r.Db.Query(`SELECT domain, COUNT(*) AS total FROM visits GROUP BY domain`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	visits := []shared.VisitCountSummaryResponse{}

	for rows.Next() {
		var domain string
		var total int64
		err = rows.Scan(&domain, &total)
		if err != nil {
			return nil, err
		}

		visits = append(visits, shared.VisitCountSummaryResponse{Domain: domain, Qty: total})
	}

	return visits, nil
}
