package main

import (
	"fmt"
	"os"
	"strconv"

	zoom "github.com/zoom-lib-golang/zoom-lib-golang"
)

var (
	apiKey = os.Getenv("ZOOM_API_KEY")
	apiSecret = os.Getenv("ZOOM_API_SECRET")
	meetingID, _ = strconv.ParseInt(os.Getenv("ZOOM_MEETING_ID"),10, 32)
)

func main() {
	fmt.Println("hello")
	client := zoom.NewClient(apiKey, apiSecret)
	info, err := client.GetMeeting(zoom.GetMeetingOptions{
		MeetingID:    int(meetingID),
	})
	if err != nil {
		fmt.Printf("err: %#v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("info: %#v\n", info)	
}