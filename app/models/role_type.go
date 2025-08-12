package models

type RoleType string

const (
	RoleSuperAdmin RoleType = "super_admin"
	RoleAdmin      RoleType = "admin"
	RoleManager    RoleType = "manager"
	RolePartner    RoleType = "partner"
	RoleClient     RoleType = "client"
)
