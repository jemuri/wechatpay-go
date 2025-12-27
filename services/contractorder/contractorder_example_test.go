// Copyright 2021 Tencent Inc. All rights reserved.

package contractorder_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func ExampleContractOrderApiService_HandleContractNotify() {
	// 注意：这是一个 v2 接口的示例，需要提供 v2 的 APIKey
	var (
		appID  string = "wxcbda96de0b165486" // 应用ID
		mchID  string = "1200009811"         // 商户号
		apiKey string = "your_api_key"       // v2 API密钥
	)

	svc := contractorder.NewContractOrderApiService(appID, mchID, apiKey)

	// 模拟 HTTP 请求 (在实际使用中，从 HTTP 服务器接收)
	// 这里只是示例，实际应从 http.Request 解析
	xmlBody := `<xml>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <sign><![CDATA[C380BEC2BFD727A4B6845133519F3AD6]]></sign>
  <mch_id>10010404</mch_id>
  <contract_code>100001256</contract_code>
  <openid><![CDATA[onqOjjmM1tad-3ROpncN-yUfa6ua]]></openid>
  <plan_id><![CDATA[123]]></plan_id>
  <change_type><![CDATA[ADD]]></change_type>
  <operate_time><![CDATA[2015-07-01 10:00:00]]></operate_time>
  <contract_id><![CDATA[Wx15463511252015071056489715]]></contract_id>
  <request_serial>1695</request_serial>
</xml>`

	// 在实际使用中，使用 http.NewRequest 或从服务器接收的请求
	req, err := http.NewRequest("POST", "https://yoursite.com/notify", strings.NewReader(xmlBody))
	if err != nil {
		log.Printf("create request failed: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/xml")

	ctx := context.Background()
	notifyReq, notifyResp, err := svc.HandleContractNotify(ctx, req)
	if err != nil {
		log.Printf("handle notify failed: %v", err)
		return
	}

	fmt.Printf("Notify ChangeType: %s\n", notifyReq.ChangeType)
	fmt.Printf("Response ReturnCode: %s\n", notifyResp.ReturnCode)

	// 根据 change_type 处理业务逻辑
	if notifyReq.ChangeType == "ADD" {
		// 处理签约成功
		fmt.Println("Contract signed successfully")
	} else if notifyReq.ChangeType == "DELETE" {
		// 处理解约成功
		fmt.Println("Contract terminated successfully")
	}
}
