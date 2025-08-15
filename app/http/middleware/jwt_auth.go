package middleware

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func JwtAuth() http.Middleware {
	return func(ctx http.Context) {
		facades.Log().Info("[JWT] Middleware called")

		raw := ctx.Request().Header("Authorization", "")
		token := strings.TrimSpace(strings.TrimPrefix(raw, "Bearer"))
		if token == "" {
			facades.Log().Info("[JWT] Missing token")
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Missing token"})
			return
		}

		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			facades.Log().Info("[JWT] Invalid or expired token: ", err)
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Invalid or expired token"})
			return
		}

		facades.Log().Info("[JWT] Token parsed successfully")
		ctx.Request().Next()
	}
}
