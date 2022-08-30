package role

import "github.com/himalayan-institute/zoom-lib-golang"

type Member struct {
	RoleID string `json:"role_id"`
	zoom.Member
}
