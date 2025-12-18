// Copyright 2021 Tencent Inc. All rights reserved.

package pappayapply_test

import (
	"context"
	"fmt"
	"log"

	"github.com/jemuri/wechatpay-go/services/pappayapply"
)

func ExamplePapPayApplyApiService_PapPayApply() {
	// 注意：这是一个 v2 接口的示例，需要提供 v2 的 APIKey
	var (
		appID  string = "wxcbda96de0b165486" // 应用ID
		mchID  string = "10000098"           // 商户号
		apiKey string = "your_api_key"       // v2 API密钥
	)

	svc := pappayapply.NewPapPayApplyApiService(appID, mchID, apiKey)

	req := &pappayapply.PapPayApplyRequest{
		NonceStr:       "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		Body:           "水电代扣",
		OutTradeNo:     "1217752501201407033233368018",
		TotalFee:       888,
		SpbillCreateIP: "8.8.8.8",
		NotifyURL:      "http://yoursite.com/wxpay.html",
		ContractID:     "Wx15463511252015071056489715",
		// 其他可选字段
	}

	ctx := context.Background()
	resp, err := svc.PapPayApply(ctx, req)
	if err != nil {
		log.Printf("pap pay apply failed: %v", err)
		return
	}

	fmt.Printf("ReturnCode: %s\n", resp.ReturnCode)
	if resp.ReturnCode == "SUCCESS" && resp.ResultCode == "SUCCESS" {
		fmt.Printf("扣款申请成功\n")
	} else {
		fmt.Printf("ErrCode: %s, ErrCodeDes: %s\n", resp.ErrCode, resp.ErrCodeDes)
	}
}
