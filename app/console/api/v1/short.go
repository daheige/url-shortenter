package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// CreateShortReq is the request struct for the CreateShort endpoint.
type CreateShortReq struct {
	g.Meta `path:"/url" tags:"Account Service" method:"Post" summary:"create a short url"`
	*model.CreateShortInput
}

// CreateShortRes is the response struct for the CreateShort endpoint.
type CreateShortRes struct {
	*model.CreateShortOutput
}

// QueryShortReq is the request struct for the QueryShort endpoint.
type QueryShortReq struct {
	g.Meta `path:"/url/:shortUrl" tags:"Account Service" method:"Get" summary:"query a short url"`
	*model.QueryShortInput
}

// QueryShortRes is the response struct for the QueryShort endpoint.
type QueryShortRes struct {
	*model.QueryShortOutput
}

// ModifyShortReq is the request struct for the ModifyShort endpoint.
type ModifyShortReq struct {
	g.Meta `path:"/url/:shortUrl/change_state" tags:"Account Service" method:"Post" summary:"modify a short url"`
	*model.ModifyShortInput
}

// ModifyShortRes is the response struct for the ModifyShort endpoint.
type ModifyShortRes struct {
	*model.ModifyShortOutput
}

// QueryStatReq is the request struct for the QueryStat endpoint.
type QueryStatReq struct {
	g.Meta `path:"/url/:shortUrl/stat" tags:"Account Service" method:"Get" summary:"query a short url stat"`
	*model.QueryStatInput
}

// QueryStatRes is the response struct for the QueryStat endpoint.
type QueryStatRes struct {
	*model.QueryStatOutput
}