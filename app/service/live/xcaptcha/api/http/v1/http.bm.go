// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api/http/v1/http.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api/http/v1/http.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

// =================
// Captcha Interface
// =================

// XCaptcha
type Captcha interface {
	// 验证码校验 `internal:"true"`
	Verify(ctx context.Context, req *XVerifyReq) (resp *XVerifyResp, err error)
}

var v1CaptchaSvc Captcha

// @params XVerifyReq
// @router GET /xlive/internal/xcaptcha/v1/captcha/verify
// @response XVerifyResp
func captchaVerify(c *bm.Context) {
	p := new(XVerifyReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1CaptchaSvc.Verify(c, p)
	c.JSON(resp, err)
}

// RegisterV1CaptchaService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1CaptchaService(e *bm.Engine, svc Captcha, midMap map[string]bm.HandlerFunc) {
	v1CaptchaSvc = svc
	e.GET("/xlive/internal/xcaptcha/v1/captcha/verify", captchaVerify)
}