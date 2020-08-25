package lead

import (
	"context"
	"fmt"

	reqModel "github.com/vupham79/convertlead-be/model/request/lead"
	resModel "github.com/vupham79/convertlead-be/model/response/lead"
	"github.com/vupham79/convertlead-be/service"
)

type authenMiddleware struct {
	nextSvc service.LeadService
}

func AuthenMiddleware() func(service.LeadService) service.LeadService {
	return func(next service.LeadService) service.LeadService {
		return &authenMiddleware{
			nextSvc: next,
		}
	}
}

func (mw authenMiddleware) Create(ctx context.Context, req *reqModel.Create) (*resModel.Create, error) {
	return mw.nextSvc.Create(ctx, req)
}

func (mw authenMiddleware) Find(ctx context.Context, req *reqModel.Find) (*resModel.Find, error) {
	return mw.nextSvc.Find(ctx, req)
}

func (mw authenMiddleware) Delete(ctx context.Context, req *reqModel.Delete) (*resModel.Delete, error) {
	return mw.nextSvc.Delete(ctx, req)
}

func (mw authenMiddleware) FindAll(ctx context.Context, req *reqModel.FindAll) (*resModel.FindAll, error) {
	fmt.Println("Authenticate FindAll!")
	return mw.nextSvc.FindAll(ctx, req)
}

func (mw authenMiddleware) Update(ctx context.Context, req *reqModel.Update) (*resModel.Update, error) {
	return mw.nextSvc.Update(ctx, req)
}
