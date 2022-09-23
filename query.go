package bw

import (
	"context"
	"fmt"
)

// Query
// @param ctx
// @param conf
// @param req
// @date 2022-09-21 16:01:33
func Query(ctx context.Context, conf *Config, req *QueryRequest) {
	if err := ctx.Err(); err != nil {
		return
	}

	resp, err := Request(ctx).Post(BaseAddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp.String())
}
