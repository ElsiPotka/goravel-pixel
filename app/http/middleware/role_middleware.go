package middleware

import (
	"pixel/app/models"
	"pixel/app/services"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

var roleService = services.NewRoleService()

func SuperAdminGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		err := facades.Auth(ctx).User(&userModel)
		if err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"message": "Unauthenticated",
			}).Abort()
			return
		}

		isSuperAdmin, err := roleService.IsSuperAdmin(&userModel)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{
				"message": "Error checking user role",
			}).Abort()
			return
		}

		if !isSuperAdmin {
			ctx.Response().Json(http.StatusForbidden, http.Json{
				"message": "Access denied.",
			}).Abort()
			return
		}

		ctx.Request().Next()
	}
}
func AdminGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		isAdmin, err := roleService.IsAdmin(&userModel)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !isAdmin {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}

func AdminOrSuperAdminGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		adminRoles := []models.RoleType{models.RoleAdmin, models.RoleSuperAdmin}
		hasAdminRole, err := roleService.HasRoles(&userModel, adminRoles)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !hasAdminRole {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}

func ManagerGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		isManager, err := roleService.IsManager(&userModel)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !isManager {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}

func PartnerGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		isPartner, err := roleService.IsPartner(&userModel)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !isPartner {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}

func ClientGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		isClient, err := roleService.IsClient(&userModel)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !isClient {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}

func StaffGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		staffRoles := []models.RoleType{models.RoleSuperAdmin, models.RoleAdmin, models.RoleManager}
		hasStaffRole, err := roleService.HasRoles(&userModel, staffRoles)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !hasStaffRole {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}

func BusinessGuard() http.Middleware {
	return func(ctx http.Context) {
		var userModel models.User
		if err := facades.Auth(ctx).User(&userModel); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{"message": "Unauthenticated"}).Abort()
			return
		}

		businessRoles := []models.RoleType{models.RolePartner, models.RoleClient}
		hasBusinessRole, err := roleService.HasRoles(&userModel, businessRoles)
		if err != nil {
			ctx.Response().Json(http.StatusInternalServerError, http.Json{"message": "Error checking user role"}).Abort()
			return
		}

		if !hasBusinessRole {
			ctx.Response().Json(http.StatusForbidden, http.Json{"message": "Access denied."}).Abort()
			return
		}

		ctx.Request().Next()
	}
}
