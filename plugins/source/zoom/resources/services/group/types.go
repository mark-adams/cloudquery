package group

import "github.com/himalayan-institute/zoom-lib-golang"

type Member struct {
	GroupID string `json:"group_id"`
	zoom.Member
}
