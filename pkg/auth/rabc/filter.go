package rabc

import "gocms/app/models/permission"

type FilterPermission struct {
	Url    string
	Method string
}

var FilterPermissions []FilterPermission

func init() {
	FilterPermissions = []FilterPermission{
		{
			"/",
			"GET",
		}, {
			"/api/login",
			"POST",
		}, {
			"/api/admin/register",
			"POST",
		}, {
			"/api/pwd",
			"GET",
		}, {
			"/api/me",
			"GET",
		}, {
			"/api/logout",
			"GET",
		},
	}
}

func Filter(permission *permission.Permission) bool {
	for _, p := range FilterPermissions {
		if p.Url == permission.Url && p.Method == permission.Method {
			return false
		}
	}
	return true
}
