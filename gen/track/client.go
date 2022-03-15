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
	GetEndpoint goa.Endpoint
}

// NewClient initializes a "track" service client given the endpoints.
func NewClient(get goa.Endpoint) *Client {
	return &Client{
		GetEndpoint: get,
	}
}

// Get calls the "get" endpoint of the "track" service.
func (c *Client) Get(ctx context.Context, p *GetTrack) (res []*Track, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Track), nil
}