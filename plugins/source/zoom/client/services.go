package client

import (
	"github.com/zoom-lib-golang/zoom-lib-golang"
)

type ZoomServices struct {
	Meeting MeetingService
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_meeting.go . MeetingService
type MeetingService interface {
	GetMeeting(opts zoom.GetMeetingOptions) (zoom.Meeting, error)
	ListMeetings(opts zoom.ListMeetingsOptions) (zoom.ListMeetingsResponse, error)
}
