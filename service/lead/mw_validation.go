package lead

import (
	"context"
	"fmt"

	"github.com/vupham79/convertlead-be/service"

	reqModel "github.com/vupham79/convertlead-be/model/request/lead"
	resModel "github.com/vupham79/convertlead-be/model/response/lead"
)

type validationMiddleware struct {
	nextSvc service.LeadService
}

// ValidationMiddleware ...
func ValidationMiddleware() func(service.LeadService) service.LeadService {
	return func(next service.LeadService) service.LeadService {
		return &validationMiddleware{
			nextSvc: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, req *reqModel.Create) (*resModel.Create, error) {
	return mw.nextSvc.Create(ctx, req)
}
func (mw validationMiddleware) FindAll(ctx context.Context, req *reqModel.FindAll) (*resModel.FindAll, error) {
	fmt.Println("Validation FindAll!")
	return mw.nextSvc.FindAll(ctx, req)
}
func (mw validationMiddleware) Find(ctx context.Context, req *reqModel.Find) (*resModel.Find, error) {
	return mw.nextSvc.Find(ctx, req)
}

func (mw validationMiddleware) Update(ctx context.Context, req *reqModel.Update) (*resModel.Update, error) {
	return mw.nextSvc.Update(ctx, req)
}
func (mw validationMiddleware) Delete(ctx context.Context, req *reqModel.Delete) (*resModel.Delete, error) {
	return mw.nextSvc.Delete(ctx, req)
}
