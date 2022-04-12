// Code generated by goa v3.6.2, DO NOT EDIT.
//
// track client
//
// Command:
// $ goa gen goa/design -o ./

package track

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "track" service client.
type Client struct {
	BatchQueryTrackInfoEndpoint goa.Endpoint
	GetTrackEndpoint            goa.Endpoint
}

// NewClient initializes a "track" service client given the endpoints.
func NewClient(batchQueryTrackInfo, getTrack goa.Endpoint) *Client {
	return &Client{
		BatchQueryTrackInfoEndpoint: batchQueryTrackInfo,
		GetTrackEndpoint:            getTrack,
	}
}

// BatchQueryTrackInfo calls the "batch_query_track_info" endpoint of the
// "track" service.
func (c *Client) BatchQueryTrackInfo(ctx context.Context, p *BatchQueryTrackPayload) (res *QueryTrackRsp, err error) {
	var ires interface{}
	ires, err = c.BatchQueryTrackInfoEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*QueryTrackRsp), nil
}

// GetTrack calls the "get_track" endpoint of the "track" service.
func (c *Client) GetTrack(ctx context.Context, p *QueryTrackPayload) (res *TrackRsp, err error) {
	var ires interface{}
	ires, err = c.GetTrackEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TrackRsp), nil
}
