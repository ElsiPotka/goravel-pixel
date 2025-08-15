package routes

import (
	"pixel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"pixel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	authController := controllers.NewAuthController()
	roleController := controllers.NewRoleController()
	oAuthController := controllers.NewOAuthController()

	facades.Route().Prefix("api/v1").Group(func(router route.Router) {

		router.Get("/users/{id}", userController.Show)

		router.Prefix("auth").Group(func(router route.Router) {

			router.Post("/register", authController.Register)
			router.Post("/login", authController.Login)
			router.Post("/logout", authController.Logout)
			router.Get("/oauth/{provider}", oAuthController.OAuthLogin)
			router.Get("/oauth/{provider}/callback", oAuthController.OAuthCallback)

		})

		router.Prefix("roles").
			Middleware(middleware.JwtAuth(), middleware.SuperAdminGuard()).
			Group(func(r route.Router) {
				r.Get("/", roleController.Index)
				r.Get("/{id}", roleController.Show)
				r.Post("/", roleController.Store)
				r.Put("/{id}", roleController.Update)
				r.Delete("/{id}", roleController.Destroy)
			})

		router.Prefix("goku").
			Middleware(middleware.JwtAuth(), middleware.SuperAdminGuard()).
			Group(func(r route.Router) {
				r.Get("/{id}", authController.Impersonate)
			})
	})
}
