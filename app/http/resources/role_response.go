package resources

type RoleResponse struct {
	ID          string `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}
