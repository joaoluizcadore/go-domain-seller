package repository

import (
	"github.com/joaoluizcadore/domain-seller/internal/domain"
	"github.com/joaoluizcadore/domain-seller/internal/shared"
)

type VisitCountRepository interface {
	Save(visit *domain.Visit) error
	GetListSummary() ([]shared.VisitCountSummaryResponse, error)
}
