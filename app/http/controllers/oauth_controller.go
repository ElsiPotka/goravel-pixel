package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/markbates/goth/gothic"
	"slices"

	"pixel/app/http/resources"
	"pixel/app/services"
)

type OAuthController struct {
	oAuthService *services.OAuthService
}

func NewOAuthController() *OAuthController {
	return &OAuthController{
		oAuthService: services.NewOAuthService(),
	}
}

func (c *OAuthController) OAuthLogin(ctx http.Context) http.Response {
	provider := ctx.Request().Route("provider")
	if provider == "" {
		response := resources.NewErrorResponse("Provider parameter is required", nil)
		return ctx.Response().Json(http.StatusBadRequest, response)
	}

	if !c.isProviderSupported(provider) {
		response := resources.NewErrorResponse("Unsupported OAuth provider", nil)
		return ctx.Response().Json(http.StatusBadRequest, response)
	}

	req := ctx.Request().Origin()
	req.URL.Path = "/oauth/" + provider

	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	authURL, err := gothic.GetAuthURL(ctx.Response().Writer(), req)
	if err != nil {
		response := resources.NewErrorResponse("Failed to get OAuth URL", nil)
		return ctx.Response().Json(http.StatusInternalServerError, response)
	}

	data := map[string]string{
		"auth_url": authURL,
		"provider": provider,
	}

	response := resources.NewSuccessResponse("OAuth URL generated successfully", data)
	return ctx.Response().Json(http.StatusOK, response)
}

func (c *OAuthController) OAuthCallback(ctx http.Context) http.Response {
	provider := ctx.Request().Route("provider")
	if provider == "" {
		response := resources.NewErrorResponse("Provider parameter is required", nil)
		return ctx.Response().Json(http.StatusBadRequest, response)
	}

	if !c.isProviderSupported(provider) {
		response := resources.NewErrorResponse("Unsupported OAuth provider", nil)
		return ctx.Response().Json(http.StatusBadRequest, response)
	}

	req := ctx.Request().Origin()
	req.URL.Path = "/oauth/" + provider + "/callback"

	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	gothUser, err := gothic.CompleteUserAuth(ctx.Response().Writer(), req)
	if err != nil {
		response := resources.NewErrorResponse("Failed to complete OAuth authentication", nil)
		return ctx.Response().Json(http.StatusUnauthorized, response)
	}

	oauthData := services.OAuthData{
		Provider:     gothUser.Provider,
		ProviderID:   gothUser.UserID,
		Email:        gothUser.Email,
		Name:         gothUser.Name,
		FirstName:    gothUser.FirstName,
		LastName:     gothUser.LastName,
		AvatarURL:    gothUser.AvatarURL,
		AccessToken:  gothUser.AccessToken,
		RefreshToken: gothUser.RefreshToken,
	}

	user, accessToken, refreshToken, expiresIn, err := c.oAuthService.ProcessOAuthUser(ctx, oauthData)
	if err != nil {
		response := resources.NewErrorResponse("Failed to process OAuth user", nil)
		return ctx.Response().Json(http.StatusInternalServerError, response)
	}

	jwtResp := resources.NewJWTResponse(accessToken, refreshToken, expiresIn)
	userResource := resources.NewUserResource(user)
	authResponse := resources.NewAuthResponse(*jwtResp, *userResource)

	response := resources.NewSuccessResponse("OAuth authentication successful", authResponse)
	return ctx.Response().Json(http.StatusOK, response)
}

func (c *OAuthController) isProviderSupported(provider string) bool {
	supportedProviders := []string{
		"google",
		"facebook",
	}

	return slices.Contains(supportedProviders, provider)
}
