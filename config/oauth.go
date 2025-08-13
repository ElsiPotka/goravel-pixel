package config

import "github.com/goravel/framework/facades"

func init() {
	config := facades.Config()
	config.Add("services", map[string]any{

		"google": map[string]any{
			"client_id":     config.Env("GOOGLE_CLIENT_ID", ""),
			"client_secret": config.Env("GOOGLE_CLIENT_SECRET", ""),
			"redirect_url":  config.Env("GOOGLE_REDIRECT_URL", ""),
		},

		"facebook": map[string]any{
			"client_id":     config.Env("FACEBOOK_CLIENT_ID", ""),
			"client_secret": config.Env("FACEBOOK_CLIENT_SECRET", ""),
			"redirect_url":  config.Env("FACEBOOK_REDIRECT_URL", ""),
		},
	})
}
