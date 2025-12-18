// Copyright 2021 Tencent Inc. All rights reserved.

package contractorder_test

import (
	"context"
	"fmt"
	"log"

	"github.com/jemuri/wechatpay-go/services/contractorder"
)

func ExampleContractOrderApiService_ContractOrder() {
	// 注意：这是一个 v2 接口的示例，需要提供 v2 的 APIKey
	var (
		appID  string = "wxcbda96de0b165486" // 应用ID
		mchID  string = "1200009811"         // 商户号
		apiKey string = "your_api_key"       // v2 API密钥
	)

	svc := contractorder.NewContractOrderApiService(appID, mchID, apiKey)

	req := &contractorder.ContractOrderRequest{
		OutTradeNo:             "123456",
		NonceStr:               "5K8264ILTKCH16CQ2502SI8ZNMTM67VS",
		Body:                   "Ipad mini 16G 白色",
		NotifyURL:              "https://weixin.qq.com",
		TotalFee:               888,
		SpbillCreateIP:         "123.12.12.123",
		TradeType:              "JSAPI",
		PlanID:                 123,
		ContractCode:           "100001256",
		RequestSerial:          1000,
		ContractDisplayAccount: "微信代扣",
		ContractNotifyURL:      "https://yoursite.com",
		// 其他可选字段
	}

	ctx := context.Background()
	resp, err := svc.ContractOrder(ctx, req)
	if err != nil {
		log.Printf("contract order failed: %v", err)
		return
	}

	fmt.Printf("ReturnCode: %s\n", resp.ReturnCode)
	if resp.ReturnCode == "SUCCESS" && resp.ResultCode == "SUCCESS" {
		fmt.Printf("PrepayID: %s\n", resp.PrepayID)
		fmt.Printf("ContractResultCode: %s\n", resp.ContractResultCode)
	}
}
