package bw

import "time"

type Config struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type ErrorResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	SubCode    string `json:"subCode"`
	SubMessage string `json:"subMessage"`
}

type BaseRequest struct {
	AppKey    string `json:"appKey" url:"appKey"` // 给开发者颁发的 key
	Format    string `json:"format" url:"format"`
	Method    string `json:"method" url:"method"`       // 接口名称
	Timestamp string `json:"timestamp" url:"timestamp"` // Unix 时间戳
	Token     string `json:"token" url:"token"`         // token 用于调用接口前授权检查。
	Type      string `json:"type" url:"type"`
	Version   string `json:"version" url:"version"`               // API 协议版本，当前版本号 3.0。
	Sign      string `json:"sign,omitempty" url:"sign,omitempty"` // 请求签名，使用 MD5，用于安全控制。
}

type BaseResponse[T any] struct {
	Error     *ErrorResponse `json:"errorResponse"`
	Method    string         `json:"method,omitempty"`
	RequestId string         `json:"requestId,omitempty"`
	Data      T              `json:"response,omitempty"`
}

type ApplyRequest struct {
	OrgNo            string `json:"orgNo"`               // 机构编号（银税通开放平台注册成功后，在开发者信息页面可查。）
	ChannelUserPhone string `json:"channelUserPhone"`    // 营销人员手机号（需要与银税通开放平台的 开发者注册手机号保持一致。）
	ProductId        string `json:"productId"`           // 产品 ID
	ChannelOrderId   string `json:"channelOrderId"`      // 渠道订单编号
	ApplyName        string `json:"applyName"`           // 姓名
	ApplyId          string `json:"applyId"`             // 身份证号
	ApplyPhone       string `json:"applyPhone"`          // 手机号
	CorpName         string `json:"corpName"`            // 企业名称 企业产品必填，个人产品非必填
	TaxNo            string `json:"taxNo"`               // 企业税号
	ApplyType        string `json:"applyType,omitempty"` // 申请人类型 1-法人 2-股东 企业产品必填，个人产品非必填
}

type TokenRequest struct {
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	Scope        string        `json:"scope"`
}

type ApplyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

type QueryRequest struct {
	OrgNo              string   `json:"orgNo"`
	ChannelOrderIdList []string `json:"channelOrderIdList"`
}
