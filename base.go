package bw

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"regexp"
	"github.com/zrb-channel/utils"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	ApiVersion = "3.0"
)

const (
	MethodToken = "baiwang.oauth.token"

	MethodApply = "winLending.finance.supermarket.productv2"

	MethodQuery = "winLending.finance.supermarket.queryChannelOrder"
)

var (
	replaceRegex = regexp.MustCompile(`[=&]`)
)

// Request
// @param ctx
// @date 2022-09-21 13:01:03
func Request(ctx context.Context) *resty.Request {
	return utils.Request(ctx).SetHeader("Content-Type", "application/json")
}

// Sign
// @param appSecret
// @param body
// @param params
// @date 2022-09-21 13:01:01
func Sign(appSecret, body string, params url.Values) string {
	params.Del("sign") // 删除签名

	paramsStr := replaceRegex.ReplaceAllString(params.Encode(), "")

	h := md5.New()
	_, _ = h.Write([]byte(appSecret + paramsStr + body + appSecret))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
