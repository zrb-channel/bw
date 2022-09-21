package bw

import (
	"context"
	"errors"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	"net/http"
)

// GetToken 百望 请求获取token
// @param ctx
// @date 2022-09-21 00:58:55
func GetToken(ctx context.Context, conf *Config) (*TokenResponse, error) {

	body := &TokenRequest{
		ClientSecret: conf.AppSecret,
		Username:     conf.Username,
		Password:     conf.Password,
	}

	params := map[string]string{
		"method":     MethodToken,
		"client_id":  conf.AppKey,
		"grant_type": "password",
		"version":    ApiVersion,
		"timestamp":  utils.Timestamp(),
	}

	resp, err := Request(ctx).SetQueryParams(params).SetBody(body).Post(BaseAddr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("请求失败")
	}

	res := &BaseResponse[*TokenResponse]{}
	if err = json.Unmarshal(resp.Body(), res); err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, errors.New(res.Error.Message)
	}

	return res.Data, nil
}
