// Copyright 2021 Tencent Inc. All rights reserved.

package contractorder

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

// ContractOrderApiService 支付中签约服务
type ContractOrderApiService struct {
	AppID  string // 应用ID
	MchID  string // 商户号
	APIKey string // API密钥 (v2 签名密钥)
}

// NewContractOrderApiService 创建支付中签约服务
func NewContractOrderApiService(appID, mchID, apiKey string) *ContractOrderApiService {
	return &ContractOrderApiService{
		AppID:  appID,
		MchID:  mchID,
		APIKey: apiKey,
	}
}

// ContractOrder 支付中签约
func (s *ContractOrderApiService) ContractOrder(ctx context.Context, req *ContractOrderRequest) (*ContractOrderResponse, error) {
	// 设置必填字段
	req.AppID = s.AppID
	req.MchID = s.MchID
	req.ContractMchID = s.MchID
	req.ContractAppID = s.AppID

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
	var response ContractOrderResponse
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
func (s *ContractOrderApiService) generateSign(req *ContractOrderRequest) (string, error) {
	// 将结构体转换为 map，只对非空参数参与签名
	params := make(map[string]string)

	if req.AppID != "" {
		params["appid"] = req.AppID
	}
	if req.MchID != "" {
		params["mch_id"] = req.MchID
	}
	if req.ContractMchID != "" {
		params["contract_mchid"] = req.ContractMchID
	}
	if req.ContractAppID != "" {
		params["contract_appid"] = req.ContractAppID
	}
	if req.OutTradeNo != "" {
		params["out_trade_no"] = req.OutTradeNo
	}
	if req.DeviceInfo != "" {
		params["device_info"] = req.DeviceInfo
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
	if req.NotifyURL != "" {
		params["notify_url"] = req.NotifyURL
	}
	if req.TotalFee != 0 {
		params["total_fee"] = fmt.Sprintf("%d", req.TotalFee)
	}
	if req.SpbillCreateIP != "" {
		params["spbill_create_ip"] = req.SpbillCreateIP
	}
	if req.TimeStart != "" {
		params["time_start"] = req.TimeStart
	}
	if req.TimeExpire != "" {
		params["time_expire"] = req.TimeExpire
	}
	if req.GoodsTag != "" {
		params["goods_tag"] = req.GoodsTag
	}
	if req.TradeType != "" {
		params["trade_type"] = req.TradeType
	}
	if req.ProductID != "" {
		params["product_id"] = req.ProductID
	}
	if req.OpenID != "" {
		params["openid"] = req.OpenID
	}
	if req.PlanID != 0 {
		params["plan_id"] = fmt.Sprintf("%d", req.PlanID)
	}
	if req.ContractCode != "" {
		params["contract_code"] = req.ContractCode
	}
	if req.RequestSerial != 0 {
		params["request_serial"] = fmt.Sprintf("%d", req.RequestSerial)
	}
	if req.ContractDisplayAccount != "" {
		params["contract_display_account"] = req.ContractDisplayAccount
	}
	if req.ContractNotifyURL != "" {
		params["contract_notify_url"] = req.ContractNotifyURL
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
func (s *ContractOrderApiService) verifySign(resp *ContractOrderResponse) error {
	params := make(map[string]string)

	if resp.ReturnCode != "" {
		params["return_code"] = resp.ReturnCode
	}
	if resp.ReturnMsg != "" {
		params["return_msg"] = resp.ReturnMsg
	}
	if resp.ResultCode != "" {
		params["result_code"] = resp.ResultCode
	}
	if resp.AppID != "" {
		params["appid"] = resp.AppID
	}
	if resp.MchID != "" {
		params["mch_id"] = resp.MchID
	}
	if resp.NonceStr != "" {
		params["nonce_str"] = resp.NonceStr
	}
	if resp.ErrCode != "" {
		params["err_code"] = resp.ErrCode
	}
	if resp.ErrCodeDes != "" {
		params["err_code_des"] = resp.ErrCodeDes
	}
	if resp.ContractResultCode != "" {
		params["contract_result_code"] = resp.ContractResultCode
	}
	if resp.ContractErrCode != "" {
		params["contract_err_code"] = resp.ContractErrCode
	}
	if resp.ContractErrCodeDes != "" {
		params["contract_err_code_des"] = resp.ContractErrCodeDes
	}
	if resp.PrepayID != "" {
		params["prepay_id"] = resp.PrepayID
	}
	if resp.TradeType != "" {
		params["trade_type"] = resp.TradeType
	}
	if resp.CodeURL != "" {
		params["code_url"] = resp.CodeURL
	}
	if resp.PlanID != 0 {
		params["plan_id"] = fmt.Sprintf("%d", resp.PlanID)
	}
	if resp.RequestSerial != 0 {
		params["request_serial"] = fmt.Sprintf("%d", resp.RequestSerial)
	}
	if resp.ContractCode != "" {
		params["contract_code"] = resp.ContractCode
	}
	if resp.ContractDisplayAccount != "" {
		params["contract_display_account"] = resp.ContractDisplayAccount
	}
	if resp.MwebURL != "" {
		params["mweb_url"] = resp.MwebURL
	}
	if resp.OutTradeNo != "" {
		params["out_trade_no"] = resp.OutTradeNo
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
func (s *ContractOrderApiService) doRequest(ctx context.Context, xmlStr string) (*http.Response, error) {
	url := "https://api.mch.weixin.qq.com/pay/contractorder"
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(xmlStr))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	return client.Do(req)
}

// HandleContractNotify 处理签约、解约结果通知
func (s *ContractOrderApiService) HandleContractNotify(ctx context.Context, req *http.Request) (*ContractNotifyRequest, *ContractNotifyResponse, error) {
	// 读取请求体
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("read request body failed: %w", err)
	}
	defer req.Body.Close()

	// 解析 XML
	var notifyReq ContractNotifyRequest
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
	response := &ContractNotifyResponse{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}

	return &notifyReq, response, nil
}

// verifyNotifySign 验证通知签名
func (s *ContractOrderApiService) verifyNotifySign(req *ContractNotifyRequest) error {
	// 将结构体转换为 map，只对非空参数参与签名
	params := make(map[string]string)

	if req.ReturnCode != "" {
		params["return_code"] = req.ReturnCode
	}
	if req.ReturnMsg != "" {
		params["return_msg"] = req.ReturnMsg
	}
	if req.ResultCode != "" {
		params["result_code"] = req.ResultCode
	}
	if req.MchID != "" {
		params["mch_id"] = req.MchID
	}
	if req.ContractCode != "" {
		params["contract_code"] = req.ContractCode
	}
	if req.PlanID != "" {
		params["plan_id"] = req.PlanID
	}
	if req.OpenID != "" {
		params["openid"] = req.OpenID
	}
	if req.ChangeType != "" {
		params["change_type"] = req.ChangeType
	}
	if req.OperateTime != "" {
		params["operate_time"] = req.OperateTime
	}
	if req.ContractID != "" {
		params["contract_id"] = req.ContractID
	}
	if req.ContractExpiredTime != "" {
		params["contract_expired_time"] = req.ContractExpiredTime
	}
	if req.ContractTerminationMode != 0 {
		params["contract_termination_mode"] = fmt.Sprintf("%d", req.ContractTerminationMode)
	}
	if req.RequestSerial != 0 {
		params["request_serial"] = fmt.Sprintf("%d", req.RequestSerial)
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
