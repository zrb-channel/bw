package main

import (
	"context"
	"fmt"
	"github.com/zrb-channel/bw"
)

func main() {
	conf := &bw.Config{
		AppKey:    "10002050",
		AppSecret: "2ffb9263-c6b1-41a3-8ff4-74670e844742",
		Username:  "77793526",
		Password:  "fsd123",
	}

	ctx := context.Background()

	auth, err := getToken(ctx, conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	create(ctx, conf, auth)
}

func create(ctx context.Context, conf *bw.Config, token string) {
	req := &bw.ApplyRequest{
		OrgNo:            "46f097f48594b61590aa",
		ChannelUserPhone: "13058300624",
		ProductId:        "SPDB-FPD",
		ChannelOrderId:   "2022091421340",
		ApplyName:        "刘国琼",
		ApplyId:          "440881198506294828",
		ApplyPhone:       "13189664629",
		CorpName:         "广东粤省事智能科技有限公司",
		TaxNo:            "91440605MA4X6KJM2N",
		ApplyType:        "",
	}

	bw.Apply(ctx, token, conf, req)
}

func getToken(ctx context.Context, conf *bw.Config) (string, error) {
	resp, err := bw.GetToken(ctx, conf)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
