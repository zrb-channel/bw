package bw

import (
	"context"
	"errors"
	"github.com/google/go-querystring/query"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	log "github.com/zrb-channel/utils/logger"
	"net/url"
)

// Query
// @param ctx
// @param conf
// @param req
// @date 2022-09-21 16:01:33
func Query(ctx context.Context, conf *Config, token string, req *QueryRequest) (*QueryResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	timestamp := utils.Timestamp()
	params := &BaseRequest{
		AppKey:    conf.AppKey,
		Format:    "json",
		Method:    MethodQuery,
		Timestamp: timestamp,
		Token:     token,
		Type:      "sync",
		Version:   ApiVersion,
	}

	bodyStr, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var queryValue url.Values
	if queryValue, err = query.Values(params); err != nil {
		return nil, err
	}

	sign := Sign(conf.AppSecret, string(bodyStr), queryValue)
	queryValue.Set("sign", sign)

	resp, err := Request(ctx).SetQueryString(queryValue.Encode()).SetBody(req).Post(BaseAddr)
	if err != nil {
		return nil, err
	}

	res := &BaseResponse[[]*QueryResponse]{}
	if err = json.Unmarshal(resp.Body(), res); err != nil {
		log.WithError(err).Error("[百望]-查询订单状态，返回数据解析失败", log.Fields(map[string]any{"conf": conf, "body": req, "resp": resp.String()}))
		return nil, err
	}

	if len(res.Data) > 0 {
		return res.Data[0], nil
	}

	return nil, errors.New("查询失败")
}
