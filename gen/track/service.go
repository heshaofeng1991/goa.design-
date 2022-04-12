// Code generated by goa v3.6.2, DO NOT EDIT.
//
// track service
//
// Command:
// $ goa gen goa/design -o ./

package track

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The track service performs operations on track order status
type Service interface {
	// BatchQueryTrackInfo implements batch_query_track_info.
	BatchQueryTrackInfo(context.Context, *BatchQueryTrackPayload) (res *QueryTrackRsp, err error)
	// GetTrack implements get_track.
	GetTrack(context.Context, *QueryTrackPayload) (res *TrackRsp, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "track"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"batch_query_track_info", "get_track"}

// BatchQueryTrackPayload is the payload type of the track service
// batch_query_track_info method.
type BatchQueryTrackPayload struct {
	// tracking number
	TrackingNumbers []string
}

// QueryTrackPayload is the payload type of the track service get_track method.
type QueryTrackPayload struct {
	// tracking number
	TrackingNumber string
}

// QueryTrackRsp is the result type of the track service batch_query_track_info
// method.
type QueryTrackRsp struct {
	// data
	Data *TrackInfo
	// code
	Code int
	// message
	Message string
}

type Track struct {
	// tracking number of order
	TrackingNumber string
	// tracking url
	TrackingURL string
	// tracking details
	Details []*TrackItem
	// status
	Status int
}

type TrackInfo struct {
	// tracks
	List []*Track
}

type TrackItem struct {
	// tracking description
	Content *string
	// tracking timestamp
	Timestamp *string
}

// TrackRsp is the result type of the track service get_track method.
type TrackRsp struct {
	// data
	Data *Track
	// code
	Code int
	// message
	Message string
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "Unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
