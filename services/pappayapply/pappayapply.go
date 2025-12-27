// Copyright 2021 Tencent Inc. All rights reserved.

package pappayapply

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

// PapPayApplyApiService 申请扣款服务
type PapPayApplyApiService struct {
	AppID  string // 应用ID
	MchID  string // 商户号
	APIKey string // API密钥 (v2 签名密钥)
}

// NewPapPayApplyApiService 创建申请扣款服务
func NewPapPayApplyApiService(appID, mchID, apiKey string) *PapPayApplyApiService {
	return &PapPayApplyApiService{
		AppID:  appID,
		MchID:  mchID,
		APIKey: apiKey,
	}
}

// PapPayApply 申请扣款
func (s *PapPayApplyApiService) PapPayApply(ctx context.Context, req *PapPayApplyRequest) (*PapPayApplyResponse, error) {
	// 设置必填字段
	req.AppID = s.AppID
	req.MchID = s.MchID
	req.TradeType = "PAP" // 固定为 PAP

	// 生成签名
	sign, err := s.generateSign(req)
	if err != nil {
		return nil, fmt.Errorf("generate sign failed: %w", err)
	}
	req.Sign = sign

	// 序列化为 XML
	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: %w", err)
	}
	xmlStr := string(xmlData)

	// 发送请求
	resp, err := s.doRequest(ctx, xmlStr)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %w", err)
	}

	// 解析 XML 响应
	var response PapPayApplyResponse
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	// 验证签名
	if response.ReturnCode == "SUCCESS" {
		if err := s.verifySign(&response); err != nil {
			return nil, fmt.Errorf("verify sign failed: %w", err)
		}
	}

	return &response, nil
}

// generateSign 生成 v2 签名
func (s *PapPayApplyApiService) generateSign(req *PapPayApplyRequest) (string, error) {
	// 将结构体转换为 map，只对非空参数参与签名
	params := make(map[string]string)

	if req.AppID != "" {
		params["appid"] = req.AppID
	}
	if req.MchID != "" {
		params["mch_id"] = req.MchID
	}
	if req.NonceStr != "" {
		params["nonce_str"] = req.NonceStr
	}
	if req.Body != "" {
		params["body"] = req.Body
	}
	if req.Detail != "" {
		params["detail"] = req.Detail
	}
	if req.Attach != "" {
		params["attach"] = req.Attach
	}
	if req.OutTradeNo != "" {
		params["out_trade_no"] = req.OutTradeNo
	}
	params["total_fee"] = fmt.Sprintf("%d", req.TotalFee)
	if req.FeeType != "" {
		params["fee_type"] = req.FeeType
	}
	if req.SpbillCreateIP != "" {
		params["spbill_create_ip"] = req.SpbillCreateIP
	}
	if req.GoodsTag != "" {
		params["goods_tag"] = req.GoodsTag
	}
	if req.NotifyURL != "" {
		params["notify_url"] = req.NotifyURL
	}
	if req.TradeType != "" {
		params["trade_type"] = req.TradeType
	}
	if req.ContractID != "" {
		params["contract_id"] = req.ContractID
	}

	// 排序 key
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接字符串
	var buf strings.Builder
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.WriteString("key=")
	buf.WriteString(s.APIKey)

	// MD5
	hash := md5.Sum([]byte(buf.String()))
	return strings.ToUpper(hex.EncodeToString(hash[:])), nil
}

// verifySign 验证响应签名
func (s *PapPayApplyApiService) verifySign(resp *PapPayApplyResponse) error {
	params := make(map[string]string)

	if resp.ReturnCode != "" {
		params["return_code"] = resp.ReturnCode
	}
	if resp.ReturnMsg != "" {
		params["return_msg"] = resp.ReturnMsg
	}
	if resp.MchID != "" {
		params["mch_id"] = resp.MchID
	}
	if resp.AppID != "" {
		params["appid"] = resp.AppID
	}
	if resp.NonceStr != "" {
		params["nonce_str"] = resp.NonceStr
	}
	if resp.ResultCode != "" {
		params["result_code"] = resp.ResultCode
	}
	if resp.ErrCode != "" {
		params["err_code"] = resp.ErrCode
	}
	if resp.ErrCodeDes != "" {
		params["err_code_des"] = resp.ErrCodeDes
	}

	// 排序 key
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接字符串
	var buf strings.Builder
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.WriteString("key=")
	buf.WriteString(s.APIKey)

	// MD5
	hash := md5.Sum([]byte(buf.String()))
	expectedSign := strings.ToUpper(hex.EncodeToString(hash[:]))

	if expectedSign != resp.Sign {
		return fmt.Errorf("sign verification failed")
	}
	return nil
}

// doRequest 发送HTTP请求
func (s *PapPayApplyApiService) doRequest(ctx context.Context, xmlStr string) (*http.Response, error) {
	url := "https://api.mch.weixin.qq.com/pay/pappayapply"
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(xmlStr))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	return client.Do(req)
}

// HandlePapPayNotify 处理扣款结果通知
func (s *PapPayApplyApiService) HandlePapPayNotify(ctx context.Context, req *http.Request) (*PapPayNotifyRequest, *PapPayNotifyResponse, error) {
	// 读取请求体
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read request body failed: %w", err)
	}
	defer req.Body.Close()

	// 解析 XML
	var notifyReq PapPayNotifyRequest
	err = xml.Unmarshal(body, &notifyReq)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshal request failed: %w", err)
	}

	// 验证签名
	if notifyReq.ReturnCode == "SUCCESS" && notifyReq.ResultCode == "SUCCESS" {
		if err := s.verifyNotifySign(&notifyReq); err != nil {
			return nil, nil, fmt.Errorf("verify sign failed: %w", err)
		}
	}

	// 返回成功响应
	response := &PapPayNotifyResponse{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}

	return &notifyReq, response, nil
}

// verifyNotifySign 验证通知签名
func (s *PapPayApplyApiService) verifyNotifySign(req *PapPayNotifyRequest) error {
	// 将结构体转换为 map，只对非空参数参与签名
	params := make(map[string]string)

	if req.ReturnCode != "" {
		params["return_code"] = req.ReturnCode
	}
	if req.ReturnMsg != "" {
		params["return_msg"] = req.ReturnMsg
	}
	if req.AppID != "" {
		params["appid"] = req.AppID
	}
	if req.MchID != "" {
		params["mch_id"] = req.MchID
	}
	if req.SubAppID != "" {
		params["sub_appid"] = req.SubAppID
	}
	if req.SubMchID != "" {
		params["sub_mch_id"] = req.SubMchID
	}
	if req.DeviceInfo != "" {
		params["device_info"] = req.DeviceInfo
	}
	if req.NonceStr != "" {
		params["nonce_str"] = req.NonceStr
	}
	if req.ResultCode != "" {
		params["result_code"] = req.ResultCode
	}
	if req.ErrCode != "" {
		params["err_code"] = req.ErrCode
	}
	if req.ErrCodeDes != "" {
		params["err_code_des"] = req.ErrCodeDes
	}
	if req.OpenID != "" {
		params["openid"] = req.OpenID
	}
	if req.SubOpenID != "" {
		params["sub_openid"] = req.SubOpenID
	}
	if req.IsSubscribe != "" {
		params["is_subscribe"] = req.IsSubscribe
	}
	if req.BankType != "" {
		params["bank_type"] = req.BankType
	}
	if req.TotalFee != 0 {
		params["total_fee"] = fmt.Sprintf("%d", req.TotalFee)
	}
	if req.FeeType != "" {
		params["fee_type"] = req.FeeType
	}
	if req.CashFee != 0 {
		params["cash_fee"] = fmt.Sprintf("%d", req.CashFee)
	}
	if req.CashFeeType != "" {
		params["cash_fee_type"] = req.CashFeeType
	}
	if req.TradeState != "" {
		params["trade_state"] = req.TradeState
	}
	if req.CouponFee != 0 {
		params["coupon_fee"] = fmt.Sprintf("%d", req.CouponFee)
	}
	if req.CouponCount != 0 {
		params["coupon_count"] = fmt.Sprintf("%d", req.CouponCount)
	}
	if req.CouponID != "" {
		params["coupon_id_0"] = req.CouponID
	}
	if req.CouponFeeN != 0 {
		params["coupon_fee_0"] = fmt.Sprintf("%d", req.CouponFeeN)
	}
	if req.TransactionID != "" {
		params["transaction_id"] = req.TransactionID
	}
	if req.OutTradeNo != "" {
		params["out_trade_no"] = req.OutTradeNo
	}
	if req.Attach != "" {
		params["attach"] = req.Attach
	}
	if req.TimeEnd != "" {
		params["time_end"] = req.TimeEnd
	}
	if req.ContractID != "" {
		params["contract_id"] = req.ContractID
	}

	// 排序 key
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接字符串
	var buf strings.Builder
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.WriteString("key=")
	buf.WriteString(s.APIKey)

	// MD5
	hash := md5.Sum([]byte(buf.String()))
	expectedSign := strings.ToUpper(hex.EncodeToString(hash[:]))

	if expectedSign != req.Sign {
		return fmt.Errorf("sign verification failed")
	}
	return nil
}
