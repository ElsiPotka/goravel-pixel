package controllers

import (
	"github.com/google/uuid"
	"pixel/app/http/requests"
	"pixel/app/http/resources"
	"pixel/app/services"
	"strings"

	"github.com/goravel/framework/contracts/http"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

func (c *AuthController) Register(ctx http.Context) http.Response {
	var req requests.RegisterUserRequest

	validationErrors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		response := resources.NewErrorResponse("An unexpected error occurred.", nil)
		return ctx.Response().Json(http.StatusInternalServerError, response)
	}
	if validationErrors != nil {
		response := resources.NewValidationErrorResponse(validationErrors)
		return ctx.Response().Json(http.StatusUnprocessableEntity, response)
	}

	data := services.RegisterData{
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email,
		Password: req.Password,
	}

	user, accessToken, refreshToken, expiresIn, err := c.authService.Register(ctx, data)
	if err != nil {
		response := resources.NewErrorResponse(err.Error(), nil)
		return ctx.Response().Json(http.StatusInternalServerError, response)
	}

	jwtResp := resources.NewJWTResponse(accessToken, refreshToken, expiresIn)
	userResource := resources.NewUserResource(user)
	authResponse := resources.NewAuthResponse(*jwtResp, *userResource)

	response := resources.NewSuccessResponse("User registered successfully", authResponse)
	return ctx.Response().Json(http.StatusCreated, response)
}

func (c *AuthController) Login(ctx http.Context) http.Response {
	var req requests.LoginUserRequest
	validationErrors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		response := resources.NewErrorResponse("An unexpected error occurred.", nil)
		return ctx.Response().Json(http.StatusInternalServerError, response)
	}
	if validationErrors != nil {
		response := resources.NewValidationErrorResponse(validationErrors)
		return ctx.Response().Json(http.StatusUnprocessableEntity, response)
	}

	data := services.LoginData{
		Email:    req.Email,
		Password: req.Password,
	}

	user, accessToken, refreshToken, expiresIn, err := c.authService.Login(ctx, data)
	if err != nil {
		status := http.StatusInternalServerError
		errorMsg := err.Error()

		switch {
		case errorMsg == "user with provided email not found":
			status = http.StatusUnauthorized
		case errorMsg == "incorrect password" || strings.Contains(errorMsg, "incorrect password"):
			status = http.StatusUnauthorized
		case strings.Contains(errorMsg, "account has been deactivated"):
			status = http.StatusForbidden
		case strings.Contains(errorMsg, "account setup incomplete"):
			status = http.StatusForbidden
		case strings.Contains(errorMsg, "created using") && strings.Contains(errorMsg, "Please use"):
			status = http.StatusBadRequest
		}

		response := resources.NewErrorResponse(err.Error(), nil)
		return ctx.Response().Json(status, response)
	}

	jwtResp := resources.NewJWTResponse(accessToken, refreshToken, expiresIn)
	userResource := resources.NewUserResource(user)
	authResponse := resources.NewAuthResponse(*jwtResp, *userResource)
	response := resources.NewSuccessResponse("User logged in successfully", authResponse)
	return ctx.Response().Json(http.StatusOK, response)
}

func (c *AuthController) Logout(ctx http.Context) http.Response {
	err := c.authService.Logout(ctx)
	if err != nil {
		response := resources.NewErrorResponse("Failed to logout", nil)
		return ctx.Response().Json(http.StatusInternalServerError, response)
	}

	response := resources.NewSuccessResponse("User logged out successfully", nil)
	return ctx.Response().Json(http.StatusOK, response)
}

func (c *AuthController) Impersonate(ctx http.Context) http.Response {
	idParam := ctx.Request().Route("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.Response().Json(400, resources.ApiResponse{
			Status:  "error",
			Message: "Invalid UUID",
			Data:    nil,
		})
	}

	user, accessToken, refreshToken, expiresIn, err := c.authService.Impersonate(ctx, id)
	if err != nil {
		return ctx.Response().Json(500, resources.ApiResponse{
			Status:  "error",
			Message: "Impersonation failed",
			Data:    nil,
		})
	}

	jwtResp := resources.NewJWTResponse(accessToken, refreshToken, expiresIn)
	userResource := resources.NewUserResource(user)
	authResponse := resources.NewAuthResponse(*jwtResp, *userResource)

	response := resources.NewSuccessResponse("Impersonation successful", authResponse)
	return ctx.Response().Json(200, response)
}
