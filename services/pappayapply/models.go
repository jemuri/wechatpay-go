// Copyright 2021 Tencent Inc. All rights reserved.

package pappayapply

// PapPayApplyRequest 申请扣款请求
type PapPayApplyRequest struct {
	AppID          string `xml:"appid"`                      // 应用ID
	MchID          string `xml:"mch_id"`                     // 商户号
	NonceStr       string `xml:"nonce_str"`                  // 随机字符串
	Sign           string `xml:"sign"`                       // 签名
	Body           string `xml:"body"`                       // 商品描述
	Detail         string `xml:"detail,omitempty"`           // 商品详情
	Attach         string `xml:"attach,omitempty"`           // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`               // 商户订单号
	TotalFee       int    `xml:"total_fee"`                  // 总金额
	FeeType        string `xml:"fee_type,omitempty"`         // 货币类型
	SpbillCreateIP string `xml:"spbill_create_ip,omitempty"` // 终端IP
	GoodsTag       string `xml:"goods_tag,omitempty"`        // 商品标记
	NotifyURL      string `xml:"notify_url"`                 // 回调通知url
	TradeType      string `xml:"trade_type"`                 // 交易类型
	ContractID     string `xml:"contract_id"`                // 委托代扣协议id
}

// PapPayApplyResponse 申请扣款响应
type PapPayApplyResponse struct {
	ReturnCode string `xml:"return_code"`            // 返回状态码
	ReturnMsg  string `xml:"return_msg,omitempty"`   // 返回信息
	MchID      string `xml:"mch_id,omitempty"`       // 商户号
	AppID      string `xml:"appid,omitempty"`        // 应用ID
	NonceStr   string `xml:"nonce_str,omitempty"`    // 随机字符串
	Sign       string `xml:"sign,omitempty"`         // 签名
	ResultCode string `xml:"result_code,omitempty"`  // 业务结果
	ErrCode    string `xml:"err_code,omitempty"`     // 错误代码
	ErrCodeDes string `xml:"err_code_des,omitempty"` // 错误代码描述
}
