package lead

import (
	"context"
	"encoding/json"
	"net/http"

	reqModel "github.com/vupham79/convertlead-be/model/request/lead"
	cErr "github.com/vupham79/convertlead-be/util/errors"
)

func FindLead(ctx context.Context, r *http.Request) (interface{}, error) {
	var req reqModel.Find
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, cErr.BodyNotAllowedError
	}

	return &req, nil
}

func FindAllLeads(ctx context.Context, r *http.Request) (interface{}, error) {
	var req reqModel.FindAll
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, cErr.BodyNotAllowedError
	}

	return &req, nil
}

func CreateLead(ctx context.Context, r *http.Request) (interface{}, error) {
	var req reqModel.Create
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, cErr.BodyNotAllowedError
	}

	return &req, nil
}

func UpdateLead(ctx context.Context, r *http.Request) (interface{}, error) {
	var req reqModel.Update
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, cErr.BodyNotAllowedError
	}

	return &req, nil
}

func DeleteLead(ctx context.Context, r *http.Request) (interface{}, error) {
	var req reqModel.Delete
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, cErr.BodyNotAllowedError
	}

	return &req, nil
}
