package routes

import (
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

		//TODO use auth middleware for role routes
		router.Prefix("roles").Group(func(router route.Router) {
			router.Get("/", roleController.Index)
			router.Get("/{id}", roleController.Show)
			router.Post("/", roleController.Store)
			router.Put("/{id}", roleController.Update)
			router.Delete("/{id}", roleController.Destroy)
		})

		router.Prefix("auth").Group(func(router route.Router) {

			router.Post("/register", authController.Register)
			router.Post("/login", authController.Login)
			router.Post("/logout", authController.Logout)
			router.Get("/oauth/{provider}", oAuthController.OAuthLogin)
			router.Get("/oauth/{provider}/callback", oAuthController.OAuthCallback)

		})
	})
}
