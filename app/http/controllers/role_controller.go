package controllers

import (
	"pixel/app/http/requests"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/http"

	"pixel/app/http/resources"
	"pixel/app/models"
	"pixel/app/services"
)

type RoleController struct {
	service *services.RoleService
}

func NewRoleController() *RoleController {
	return &RoleController{
		service: services.NewRoleService(),
	}
}

func toResponse(role *models.Role) resources.RoleResponse {
	return resources.RoleResponse{
		ID:          role.ID.String(),
		Role:        string(role.Role),
		Description: role.Description,
		Active:      role.IsActive,
	}
}

func (c *RoleController) Index(ctx http.Context) http.Response {
	roles, err := c.service.GetAll()
	if err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{Status: "error", Message: "Failed to fetch roles", Data: nil})
	}

	var resp []resources.RoleResponse
	for _, r := range roles {
		resp = append(resp, toResponse(&r))
	}

	return ctx.Response().Json(200, resources.ApiResponse{Status: "success", Message: "Roles retrieved", Data: resp})
}

func (c *RoleController) Show(ctx http.Context) http.Response {
	idParam := ctx.Request().Route("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.Response().Json(400, resources.ApiResponse{Status: "error", Message: "Invalid UUID", Data: nil})
	}

	role, err := c.service.GetByID(id)
	if err != nil {
		return ctx.Response().Json(404, resources.ApiResponse{Status: "error", Message: "Role not found", Data: nil})
	}
	return ctx.Response().Json(200, resources.ApiResponse{Status: "success", Message: "Role retrieved", Data: toResponse(role)})
}

func (c *RoleController) Store(ctx http.Context) http.Response {
	var req requests.RoleRequest
	errors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{
			Status:  "error",
			Message: "Validation processing failed",
			Data:    nil,
		})
	}
	if errors != nil {
		return ctx.Response().Json(422, resources.ApiResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    errors.All(),
		})
	}

	input := models.Role{
		Role:        models.RoleType(req.Role),
		Description: req.Description,
		IsActive:    req.IsActive == "true",
	}

	if err := c.service.Create(input); err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{
			Status:  "error",
			Message: "Role creation failed",
			Data:    nil,
		})
	}

	return ctx.Response().Json(201, resources.ApiResponse{
		Status:  "success",
		Message: "Role created",
		Data:    nil,
	})
}

func (c *RoleController) Update(ctx http.Context) http.Response {
	idParam := ctx.Request().Route("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.Response().Json(400, resources.ApiResponse{
			Status:  "error",
			Message: "Invalid UUID",
			Data:    nil,
		})
	}

	var req requests.RoleRequest
	errors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{
			Status:  "error",
			Message: "Validation processing failed",
			Data:    nil,
		})
	}
	if errors != nil {
		return ctx.Response().Json(422, resources.ApiResponse{
			Status:  "error",
			Message: "Validation failed",
			Data:    errors.All(),
		})
	}

	updated := models.Role{
		Role:        models.RoleType(req.Role),
		Description: req.Description,
		IsActive:    req.IsActive == "true",
	}

	if err := c.service.Update(id, updated); err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{
			Status:  "error",
			Message: "Role update failed",
			Data:    nil,
		})
	}

	return ctx.Response().Json(200, resources.ApiResponse{
		Status:  "success",
		Message: "Role updated",
		Data:    nil,
	})
}

func (c *RoleController) Destroy(ctx http.Context) http.Response {
	idParam := ctx.Request().Route("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.Response().Json(400, resources.ApiResponse{Status: "error", Message: "Invalid UUID", Data: nil})
	}

	if err := c.service.Delete(id); err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{Status: "error", Message: "Role deletion failed", Data: nil})
	}
	return ctx.Response().Json(200, resources.ApiResponse{Status: "success", Message: "Role deleted", Data: nil})
}
