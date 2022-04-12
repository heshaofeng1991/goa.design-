// Code generated by goa v3.6.2, DO NOT EDIT.
//
// HTTP request path constructors for the track service.
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"fmt"
)

// BatchQueryTrackInfoTrackPath returns the URL path to the track service batch_query_track_info HTTP endpoint.
func BatchQueryTrackInfoTrackPath() string {
	return "/v1/tracks"
}

// GetTrackTrackPath returns the URL path to the track service get_track HTTP endpoint.
func GetTrackTrackPath(trackingNumber string) string {
	return fmt.Sprintf("/v1/tracks/%v", trackingNumber)
}
