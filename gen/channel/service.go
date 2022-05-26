// Code generated by goa v3.6.2, DO NOT EDIT.
//
// channel service
//
// Command:
// $ goa gen goa/design -o ./

package channel

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The quote service performs operations on quotation
type Service interface {
	// UpdateChannelCostStatus implements UpdateChannelCostStatus.
	UpdateChannelCostStatus(context.Context, *UpdateChannelCostStatusReq) (res *UpdateChannelCostStatusRsp, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "channel"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"UpdateChannelCostStatus"}

// UpdateChannelCostStatusReq is the payload type of the channel service
// UpdateChannelCostStatus method.
type UpdateChannelCostStatusReq struct {
	// 渠道ID
	Ids []int64
	// 排除国际二字码
	CountryCodes []string
	// 状态（0 不启用 1 启用）
	Status bool
	// Authorization
	Authorization *string
	// JWT used for authentication
	Token *string
}

// UpdateChannelCostStatusRsp is the result type of the channel service
// UpdateChannelCostStatus method.
type UpdateChannelCostStatusRsp struct {
	// data
	Data *UpdateCustomerConfigData
	// code
	Code int
	// message
	Message string
}

type UpdateCustomerConfigData struct {
	// 状态（0 更新成功 1 更新失败）
	Status int32
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "Unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
