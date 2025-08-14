package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

type OauthServiceProvider struct {
}

func (receiver *OauthServiceProvider) Register(app foundation.Application) {
	goth.UseProviders(
		google.New(
			facades.Config().GetString("services.google.client_id"),
			facades.Config().GetString("services.google.client_secret"),
			facades.Config().GetString("services.google.redirect_url"),
			"email", "profile",
		),
		facebook.New(
			facades.Config().GetString("services.facebook.client_id"),
			facades.Config().GetString("services.facebook.client_secret"),
			facades.Config().GetString("services.facebook.redirect_url"),
			"email",
		),
	)
}

func (receiver *OauthServiceProvider) Boot(app foundation.Application) {
}
