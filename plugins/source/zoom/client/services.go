package client

import (
	"github.com/himalayan-institute/zoom-lib-golang"
)

type ZoomServices interface {
	AccountService
	GroupService
	MeetingService
	RecordingService
	RoleService
	UserService
}

var _ ZoomServices = (*zoom.Client)(nil)

//go:generate mockgen -package=mocks -destination=./mocks/mock_account.go . AccountService
type AccountService interface {
	GetAccountLockSettings(opts zoom.GetAccountLockSettingsOpts) (zoom.AccountLockSettings, error)
	GetAccountSettings(opts zoom.GetAccountSettingsOpts) (zoom.AccountSettings, error)
	GetAccountManagedDomains(opts zoom.GetAccountManagedDomainsOpts) (zoom.AccountManagedDomains, error)
	GetAccountTrustedDomains(opts zoom.GetAccountTrustedDomainsOpts) (zoom.AccountTrustedDomains, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_group.go . GroupService
type GroupService interface {
	ListGroups() (zoom.ListGroupsResponse, error)
	ListGroupMembers(options zoom.ListGroupMembersOptions) (zoom.ListGroupMembersResponse, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_meeting.go . MeetingService
type MeetingService interface {
	GetMeeting(opts zoom.GetMeetingOptions) (zoom.Meeting, error)
	ListMeetings(opts zoom.ListMeetingsOptions) (zoom.ListMeetingsResponse, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_recording.go . RecordingService
type RecordingService interface {
	ListAllRecordings(opts zoom.ListAllRecordingsOptions) (zoom.ListAllRecordingsResponse, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_role.go . RoleService
type RoleService interface {
	GetRole(opts zoom.GetRoleOpts) (zoom.Role, error)
	ListRoles() (zoom.ListRolesResponse, error)
	ListRoleMembers(options zoom.ListRoleMembersOptions) (zoom.ListRoleMembersResponse, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_user.go . UserService
type UserService interface {
	GetUser(opts zoom.GetUserOpts) (zoom.User, error)
	ListUsers(opts zoom.ListUsersOptions) (zoom.ListUsersResponse, error)
}
