package lead

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"

	reqModel "github.com/vupham79/convertlead-be/model/request/lead"
	resModel "github.com/vupham79/convertlead-be/model/response/lead"
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

func (s *LeadService) Create(_ context.Context, p *reqModel.Create) (*resModel.Create, error) {
	// return s.db.Create(p).Error
	return nil, nil
}

func (s *LeadService) Update(_ context.Context, p *reqModel.Update) (*resModel.Update, error) {
	// old := domain.Lead{Model: domain.Model{ID: p.ID}}
	// if err := s.db.Find(&old).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return nil, ErrNotFound
	// 	}
	// 	return nil, err
	// }

	// old.Name = p.Name
	// old.Email = p.Email

	// return &old, s.db.Save(&old).Error
	return nil, nil
}

func (s *LeadService) Find(_ context.Context, p *reqModel.Find) (*resModel.Find, error) {
	// res := p
	// if err := s.db.Find(&res).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return nil, ErrNotFound
	// 	}
	// 	return nil, err
	// }

	// return res, nil
	return nil, nil
}

// FindAll implement FindAll for Lead service
func (s *LeadService) FindAll(_ context.Context, p *reqModel.FindAll) (*resModel.FindAll, error) {
	// res := []domain.Lead{}
	// return res, s.db.Find(&res).Error
	fmt.Println("Logic FindAll!")
	return nil, nil
}

// Delete implement Delete for Lead service
func (s *LeadService) Delete(_ context.Context, p *reqModel.Delete) (*resModel.Delete, error) {
	// old := domain.Lead{Model: domain.Model{ID: p.ID}}
	// if err := s.db.Find(&old).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return ErrNotFound
	// 	}
	// 	return err
	// }
	// return s.db.Delete(old).Error
	return nil, nil
}
