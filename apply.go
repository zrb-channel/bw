package bw

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/zrb-channel/utils"
	"net/url"

	"github.com/google/go-querystring/query"

	json "github.com/json-iterator/go"
)

// Apply
// @param ctx
// @param token
// @param order
// @date 2022-09-21 02:46:17
func Apply(ctx context.Context, token string, conf *Config, body *ApplyRequest) (*ApplyResponse, error) {

	timestamp := utils.Timestamp()

	params := &BaseRequest{
		AppKey:    conf.AppKey,
		Format:    "json",
		Method:    MethodApply,
		Timestamp: timestamp,
		Token:     token,
		Type:      "sync",
		Version:   ApiVersion,
	}

	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var queryValue url.Values
	if queryValue, err = query.Values(params); err != nil {
		return nil, err
	}

	sign := Sign(conf.AppSecret, string(bodyStr), queryValue)
	queryValue.Set("sign", sign)

	var resp *resty.Response
	if resp, err = Request(ctx).SetQueryString(queryValue.Encode()).SetBody(body).Post(BaseAddr); err != nil {
		return nil, err
	}

	res := &BaseResponse[*ApplyResponse]{}
	if err = json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	return res.Data, nil
}
