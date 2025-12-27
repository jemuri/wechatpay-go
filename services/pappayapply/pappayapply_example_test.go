// Copyright 2021 Tencent Inc. All rights reserved.

package pappayapply_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func ExamplePapPayApplyApiService_HandlePapPayNotify() {
	// 注意：这是一个 v2 接口的示例，需要提供 v2 的 APIKey
	var (
		appID  string = "wxcbda96de0b165486" // 应用ID
		mchID  string = "10000098"           // 商户号
		apiKey string = "your_api_key"       // v2 API密钥
	)

	svc := pappayapply.NewPapPayApplyApiService(appID, mchID, apiKey)

	// 模拟 HTTP 请求 (在实际使用中，从 HTTP 服务器接收)
	// 这里只是示例，实际应从 http.Request 解析
	xmlBody := `<xml>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <attach><![CDATA[支付测试]]></attach>
  <bank_type><![CDATA[CFT]]></bank_type>
  <fee_type><![CDATA[CNY]]></fee_type>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
  <openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
  <out_trade_no><![CDATA[1409811653]]></out_trade_no>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
  <time_end><![CDATA[20140903131540]]></time_end>
  <total_fee>1</total_fee>
  <transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
  <contract_id><![CDATA[Wx15463511252015071056489715]]></contract_id>
</xml>`

	// 在实际使用中，使用 http.NewRequest 或从服务器接收的请求
	req, err := http.NewRequest("POST", "https://yoursite.com/notify", strings.NewReader(xmlBody))
	if err != nil {
		log.Printf("create request failed: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/xml")

	ctx := context.Background()
	notifyReq, notifyResp, err := svc.HandlePapPayNotify(ctx, req)
	if err != nil {
		log.Printf("handle notify failed: %v", err)
		return
	}

	fmt.Printf("Notify TradeState: %s\n", notifyReq.TradeState)
	fmt.Printf("Response ReturnCode: %s\n", notifyResp.ReturnCode)

	// 根据 trade_state 处理业务逻辑
	if notifyReq.TradeState == "SUCCESS" {
		// 扣款成功处理
		fmt.Println("Payment successful")
	} else {
		// 扣款失败处理
		fmt.Printf("Payment failed: %s\n", notifyReq.ErrCode)
	}
}
