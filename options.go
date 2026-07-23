package nga

import (
	"github.com/go-resty/resty/v2"
)

type Context struct {
	Req *resty.Request
}
type Option interface {
	Apply(*Context)
}

func ApplyOptions(opts ...Option) *Context {
	c := Context{}
	for _, opt := range opts {
		opt.Apply(&c)
	}
	return &c
}

type OptionFn func(ctx *Context)

func (o OptionFn) Apply(ctx *Context) {
	o(ctx)
}

type ReqOptionFn func(request *resty.Request) *resty.Request

func (o ReqOptionFn) Apply(c *Context) {
	c.Req = o(c.Req)
}

var WithQueryParam = func(param, value string) ReqOptionFn {
	return func(request *resty.Request) *resty.Request {
		return request.SetQueryParam(param, value)
	}
}

var WithQueryParams = func(params map[string]string) ReqOptionFn {
	return func(request *resty.Request) *resty.Request {
		return request.SetQueryParams(params)
	}
}

var WithReq = func(request *resty.Request) Option {
	return OptionFn(func(ctx *Context) {
		ctx.Req = request
	})
}
