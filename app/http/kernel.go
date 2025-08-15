package http

import (
	"pixel/app/http/middleware"

	"github.com/goravel/framework/contracts/http"
)

type Kernel struct {
}

// The application's global HTTP middleware stack.
// These middleware are run during every request to your application.
func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{}
}

func (kernel Kernel) RouteMiddleware() map[string]http.Middleware {
	return map[string]http.Middleware{
		"jwt_auth":             middleware.JwtAuth(),
		"super_admin":          middleware.SuperAdminGuard(),
		"admin":                middleware.AdminGuard(),
		"admin_or_super_admin": middleware.AdminOrSuperAdminGuard(),
		"manager":              middleware.ManagerGuard(),
		"partner":              middleware.PartnerGuard(),
		"client":               middleware.ClientGuard(),
		"staff":                middleware.StaffGuard(),
		"business":             middleware.BusinessGuard(),
	}
}
