package controllers

import (
	"pixel/app/http/requests"
	"pixel/app/http/resources"
	"pixel/app/services"

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
		if err.Error() == "user with provided email is not found" || err.Error() == "incorrect password" {
			status = http.StatusUnauthorized
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
