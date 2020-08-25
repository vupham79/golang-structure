package source

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/vupham79/convertlead-be/model/domain"
	"github.com/vupham79/convertlead-be/service"
)

type LeadService struct {
	db *gorm.DB
}

func NewLeadService(db *gorm.DB) service.LeadService {
	return &LeadService{
		db: db,
	}
}

func (s *LeadService) Create(_ context.Context, p *domain.Lead) error {
	return s.db.Create(p).Error
}

func (s *LeadService) Update(_ context.Context, p *domain.Lead) (*domain.Lead, error) {
	old := domain.Lead{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	old.Name = p.Name
	old.Email = p.Email

	return &old, s.db.Save(&old).Error
}

func (s *LeadService) Find(_ context.Context, p *domain.Lead) (*domain.Lead, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Lead service
func (s *LeadService) FindAll(_ context.Context) ([]domain.Lead, error) {
	res := []domain.Lead{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Lead service
func (s *LeadService) Delete(_ context.Context, p *domain.Lead) error {
	old := domain.Lead{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}
