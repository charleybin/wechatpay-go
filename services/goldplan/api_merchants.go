// Copyright 2021 Tencent Inc. All rights reserved.
//
// 点金计划对外API
//
// 特约商户点金计划管理API
//
// API version: 0.2.3

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package goldplan

import (
	"context"
	nethttp "net/http"
	neturl "net/url"

	"github.com/charleybin/wechatpay-go/core"
	"github.com/charleybin/wechatpay-go/core/consts"
	"github.com/charleybin/wechatpay-go/services"
)

type MerchantsApiService services.Service

// CloseAdvertisingShow 关闭广告展示
//
// 使用此接口为特约商户的点金计划页面关闭广告展示功能
func (a *MerchantsApiService) CloseAdvertisingShow(ctx context.Context, req CloseAdvertisingShowRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/goldplan/merchants/close-advertising-show"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// OpenAdvertisingShow 开通广告展示
//
// 此接口为特约商户的点金计划页面开通广告展示功能，可同时配置同业过滤标签，防止特约商户支付后出现同行业的广告内容。最多传入3个同业过滤标签值
func (a *MerchantsApiService) OpenAdvertisingShow(ctx context.Context, req OpenAdvertisingShowRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/goldplan/merchants/open-advertising-show"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SetAdvertisingIndustryFilter 同业过滤标签管理
//
// 服务商帮助特约商户设置点金计划同业过滤信息，最多传入3个同业过滤标签值
func (a *MerchantsApiService) SetAdvertisingIndustryFilter(ctx context.Context, req SetAdvertisingIndustryFilterRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/goldplan/merchants/set-advertising-industry-filter"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}
