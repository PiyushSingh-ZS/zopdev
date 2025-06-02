package resourcegroup

import (
	"strconv"

	"gofr.dev/pkg/gofr"
	gofrHttp "gofr.dev/pkg/gofr/http"

	"github.com/zopdev/zopdev/api/resources/models"
)

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetAllResourceGroups(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"id"}}
	}

	accID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"id"}}
	}

	res, err := h.svc.GetAllResourceGroups(ctx, accID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) GetResourceGroup(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"id"}}
	}

	accID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"id"}}
	}

	idStr := ctx.PathParam("rgID")
	if idStr == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"rgId"}}
	}

	rgID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"rgId"}}
	}

	res, err := h.svc.GetResourceGroupByID(ctx, accID, rgID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) CreateResourceGroup(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"id"}}
	}

	accID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"id"}}
	}

	var rg models.RGCreate

	err = ctx.Bind(&rg)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"body"}}
	}

	rg.CloudAccountID = accID

	res, err := h.svc.CreateResourceGroup(ctx, &rg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) UpdateResourceGroup(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"id"}}
	}

	accID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"id"}}
	}

	grpID := ctx.PathParam("rgID")
	if grpID == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"rgId"}}
	}

	groupID, err := strconv.ParseInt(grpID, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"rgId"}}
	}

	var rg models.RGUpdate

	err = ctx.Bind(&rg)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"body"}}
	}

	rg.CloudAccountID = accID
	rg.ID = groupID

	res, err := h.svc.UpdateResourceGroup(ctx, &rg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *Handler) DeleteResourceGroup(ctx *gofr.Context) (any, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"id"}}
	}

	accID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"id"}}
	}

	idStr := ctx.PathParam("rgID")
	if idStr == "" {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"rgId"}}
	}

	rgID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, gofrHttp.ErrorInvalidParam{Params: []string{"rgId"}}
	}

	err = h.svc.DeleteResourceGroup(ctx, accID, rgID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
