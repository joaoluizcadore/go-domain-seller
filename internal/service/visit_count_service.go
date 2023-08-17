package service

import (
	"github.com/joaoluizcadore/domain-seller/internal/domain"
	"github.com/joaoluizcadore/domain-seller/internal/repository"
	"github.com/joaoluizcadore/domain-seller/internal/shared"
)

type VisitCountService struct {
	visitCountRepository repository.VisitCountRepository
}

func NewVisitCountService(visitCountRepository repository.VisitCountRepository) *VisitCountService {
	return &VisitCountService{
		visitCountRepository: visitCountRepository,
	}
}

func (s *VisitCountService) AddVisit(visit *domain.Visit) error {
	return s.visitCountRepository.Save(visit)
}

func (s *VisitCountService) GetListSummary() ([]shared.VisitCountSummaryResponse, error) {
	return s.visitCountRepository.GetListSummary()
}
