package bw

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/zrb-channel/utils"
	log "github.com/zrb-channel/utils/logger"
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
	if err := ctx.Err(); err != nil {
		return nil, err
	}

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
		log.WithError(err).Error("[百望]-获取联登地址，请求失败", log.Fields(map[string]any{"conf": conf, "body": body}))
		return nil, err
	}

	res := &BaseResponse[*ApplyResponse]{}
	if err = json.Unmarshal(resp.Body(), res); err != nil {
		log.WithError(err).Error("[百望]-获取联登地址，返回数据解析失败", log.Fields(map[string]any{"conf": conf, "body": body, "resp": resp.String()}))
		return nil, err
	}

	if res.Error != nil {
		log.WithError(err).Error("[百望]-获取联登地址，返回数据有误", log.Fields(map[string]any{"conf": conf, "body": body, "params": params, "resp": resp.String(), "result": res}))
		if v := res.Error.SubMessage; v != "" {
			return nil, errors.New(v)
		}
		return nil, errors.New(res.Error.Message)
	}

	return res.Data, nil
}
