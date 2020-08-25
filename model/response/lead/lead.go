package lead

import "github.com/vupham79/convertlead-be/model/domain"

type Create struct{}
type Update struct{}
type Find struct{}
type FindAll struct {
	Leads []domain.Lead `json:"leads"`
}
type Delete struct{}
