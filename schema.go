package bw

import "time"

type (
	// Config 百望配置参数
	Config struct {
		AppKey string `json:"app_key"`

		AppSecret string `json:"app_secret"`

		Username string `json:"username"`

		Password string `json:"password"`
	}

	ErrorResponse struct {
		Code int `json:"code"`

		Message string `json:"message"`

		SubCode string `json:"subCode"`

		SubMessage string `json:"subMessage"`
	}

	BaseRequest struct {
		// 给开发者颁发的 key
		AppKey string `json:"appKey" url:"appKey"`

		Format string `json:"format" url:"format"`

		// 接口名称
		Method string `json:"method" url:"method"`

		// Unix 时间戳
		Timestamp string `json:"timestamp" url:"timestamp"`

		// token 用于调用接口前授权检查。
		Token string `json:"token" url:"token"`

		Type string `json:"type" url:"type"`

		// Version API 协议版本，当前版本号 3.0。
		Version string `json:"version" url:"version"`

		// Sign 请求签名，使用 MD5，用于安全控制。
		Sign string `json:"sign,omitempty" url:"sign,omitempty"`
	}

	BaseResponse[T any] struct {
		Error     *ErrorResponse `json:"errorResponse"`
		Method    string         `json:"method,omitempty"`
		RequestId string         `json:"requestId,omitempty"`
		Data      T              `json:"response,omitempty"`
	}

	ApplyRequest struct {
		// OrgNo 机构编号（银税通开放平台注册成功后，在开发者信息页面可查。）
		OrgNo string `json:"orgNo"`

		// ChannelUserPhone 营销人员手机号（需要与银税通开放平台的 开发者注册手机号保持一致。）
		ChannelUserPhone string `json:"channelUserPhone"`

		// ProductId 产品ID
		ProductId string `json:"productId"`

		// ChannelOrderId 渠道订单编号
		ChannelOrderId string `json:"channelOrderId"`

		// ApplyName 姓名
		ApplyName string `json:"applyName"`

		// ApplyId 身份证号
		ApplyId string `json:"applyId"`

		// ApplyPhone 手机号
		ApplyPhone string `json:"applyPhone"`

		// CorpName 企业名称 企业产品必填，个人产品非必填
		CorpName string `json:"corpName"`

		// TaxNo 企业税号
		TaxNo string `json:"taxNo"`

		// ApplyType 申请人类型 1-法人 2-股东 企业产品必填，个人产品非必填
		ApplyType string `json:"applyType,omitempty"`
	}

	TokenRequest struct {
		ClientSecret string `json:"client_secret"`
		Username     string `json:"username"`
		Password     string `json:"password"`
	}

	TokenResponse struct {
		AccessToken  string        `json:"access_token"`
		TokenType    string        `json:"token_type"`
		RefreshToken string        `json:"refresh_token"`
		ExpiresIn    time.Duration `json:"expires_in"`
		Scope        string        `json:"scope"`
	}

	ApplyResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Url     string `json:"url"`
	}

	QueryRequest struct {
		OrgNo              string   `json:"orgNo"`
		ChannelOrderIdList []string `json:"channelOrderIdList"`
	}
)
