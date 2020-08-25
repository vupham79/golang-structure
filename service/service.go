package service

import (
	"context"

	leadReq "github.com/vupham79/convertlead-be/model/request/lead"
	leadRes "github.com/vupham79/convertlead-be/model/response/lead"
)

type Service struct {
	LeadService   LeadService
	RoleService   RoleService
	SourceService SourceService
}

type LeadService interface {
	Create(ctx context.Context, p *leadReq.Create) (*leadRes.Create, error)
	Update(ctx context.Context, p *leadReq.Update) (*leadRes.Update, error)
	Find(ctx context.Context, p *leadReq.Find) (*leadRes.Find, error)
	FindAll(ctx context.Context, p *leadReq.FindAll) (*leadRes.FindAll, error)
	Delete(ctx context.Context, p *leadReq.Delete) (*leadRes.Delete, error)
}

type RoleService interface {
	// Create(ctx context.Context, p *domain.Role) error
	// Update(ctx context.Context, p *domain.Role) (*domain.Role, error)
	// Find(ctx context.Context, p *domain.Role) (*domain.Role, error)
	// FindAll(ctx context.Context) ([]domain.Role, error)
	// Delete(ctx context.Context, p *domain.Role) error
}

type SourceService interface {
	// Create(ctx context.Context, p *domain.Source) error
	// Update(ctx context.Context, p *domain.Source) (*domain.Source, error)
	// Find(ctx context.Context, p *domain.Source) (*domain.Source, error)
	// FindAll(ctx context.Context) ([]domain.Source, error)
	// Delete(ctx context.Context, p *domain.Source) error
}
