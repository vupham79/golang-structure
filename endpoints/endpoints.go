package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/vupham79/convertlead-be/service"

	"github.com/vupham79/convertlead-be/endpoints/lead"
)

// Endpoints .
type Endpoints struct {
	// Lead
	FindLead     endpoint.Endpoint
	FindAllLeads endpoint.Endpoint
	CreateLead   endpoint.Endpoint
	UpdateLead   endpoint.Endpoint
	DeleteLead   endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		// Lead
		FindLead:     lead.MakeFindEndpoint(s),
		FindAllLeads: lead.MakeFindAllEndpoint(s),
		CreateLead:   lead.MakeCreateEndpoint(s),
		UpdateLead:   lead.MakeUpdateEndpoint(s),
		DeleteLead:   lead.MakeDeleteEndpoint(s),
	}
}
