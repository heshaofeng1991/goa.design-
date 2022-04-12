// Code generated by goa v3.6.2, DO NOT EDIT.
//
// track HTTP server types
//
// Command:
// $ goa gen goa/design -o ./

package server

import (
	track "goa/gen/track"

	goa "goa.design/goa/v3/pkg"
)

// BatchQueryTrackInfoResponseBody is the type of the "track" service
// "batch_query_track_info" endpoint HTTP response body.
type BatchQueryTrackInfoResponseBody struct {
	// data
	Data *TrackInfoResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
}

// GetTrackResponseBody is the type of the "track" service "get_track" endpoint
// HTTP response body.
type GetTrackResponseBody struct {
	// data
	Data *TrackResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// code
	Code int `form:"code" json:"code" xml:"code"`
	// message
	Message string `form:"message" json:"message" xml:"message"`
}

// BatchQueryTrackInfoUnauthorizedResponseBody is the type of the "track"
// service "batch_query_track_info" endpoint HTTP response body for the
// "Unauthorized" error.
type BatchQueryTrackInfoUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetTrackUnauthorizedResponseBody is the type of the "track" service
// "get_track" endpoint HTTP response body for the "Unauthorized" error.
type GetTrackUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// TrackInfoResponseBody is used to define fields on response body types.
type TrackInfoResponseBody struct {
	// tracks
	List []*TrackResponseBody `form:"list" json:"list" xml:"list"`
}

// TrackResponseBody is used to define fields on response body types.
type TrackResponseBody struct {
	// tracking number of order
	TrackingNumber string `form:"tracking_number" json:"tracking_number" xml:"tracking_number"`
	// tracking url
	TrackingURL string `form:"tracking_url" json:"tracking_url" xml:"tracking_url"`
	// tracking details
	Details []*TrackItemResponseBody `form:"details" json:"details" xml:"details"`
	// status
	Status int `form:"status" json:"status" xml:"status"`
}

// TrackItemResponseBody is used to define fields on response body types.
type TrackItemResponseBody struct {
	// tracking description
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// tracking timestamp
	Timestamp *string `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

// NewBatchQueryTrackInfoResponseBody builds the HTTP response body from the
// result of the "batch_query_track_info" endpoint of the "track" service.
func NewBatchQueryTrackInfoResponseBody(res *track.QueryTrackRsp) *BatchQueryTrackInfoResponseBody {
	body := &BatchQueryTrackInfoResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = marshalTrackTrackInfoToTrackInfoResponseBody(res.Data)
	}
	return body
}

// NewGetTrackResponseBody builds the HTTP response body from the result of the
// "get_track" endpoint of the "track" service.
func NewGetTrackResponseBody(res *track.TrackRsp) *GetTrackResponseBody {
	body := &GetTrackResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = marshalTrackTrackToTrackResponseBody(res.Data)
	}
	return body
}

// NewBatchQueryTrackInfoUnauthorizedResponseBody builds the HTTP response body
// from the result of the "batch_query_track_info" endpoint of the "track"
// service.
func NewBatchQueryTrackInfoUnauthorizedResponseBody(res *goa.ServiceError) *BatchQueryTrackInfoUnauthorizedResponseBody {
	body := &BatchQueryTrackInfoUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetTrackUnauthorizedResponseBody builds the HTTP response body from the
// result of the "get_track" endpoint of the "track" service.
func NewGetTrackUnauthorizedResponseBody(res *goa.ServiceError) *GetTrackUnauthorizedResponseBody {
	body := &GetTrackUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewBatchQueryTrackInfoBatchQueryTrackPayload builds a track service
// batch_query_track_info endpoint payload.
func NewBatchQueryTrackInfoBatchQueryTrackPayload(trackingNumbers []string) *track.BatchQueryTrackPayload {
	v := &track.BatchQueryTrackPayload{}
	v.TrackingNumbers = trackingNumbers

	return v
}

// NewGetTrackQueryTrackPayload builds a track service get_track endpoint
// payload.
func NewGetTrackQueryTrackPayload(trackingNumber string) *track.QueryTrackPayload {
	v := &track.QueryTrackPayload{}
	v.TrackingNumber = trackingNumber

	return v
}
