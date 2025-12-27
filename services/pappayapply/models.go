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

// PapPayNotifyRequest 扣款结果通知请求
type PapPayNotifyRequest struct {
	ReturnCode    string `xml:"return_code"`              // 返回状态码
	ReturnMsg     string `xml:"return_msg,omitempty"`     // 返回信息
	AppID         string `xml:"appid,omitempty"`          // 应用ID
	MchID         string `xml:"mch_id,omitempty"`         // 商户号
	SubAppID      string `xml:"sub_appid,omitempty"`      // 子商户应用ID
	SubMchID      string `xml:"sub_mch_id,omitempty"`     // 子商户号
	DeviceInfo    string `xml:"device_info,omitempty"`    // 设备号
	NonceStr      string `xml:"nonce_str,omitempty"`      // 随机字符串
	Sign          string `xml:"sign,omitempty"`           // 签名
	ResultCode    string `xml:"result_code,omitempty"`    // 业务结果
	ErrCode       string `xml:"err_code,omitempty"`       // 错误代码
	ErrCodeDes    string `xml:"err_code_des,omitempty"`   // 错误代码描述
	OpenID        string `xml:"openid,omitempty"`         // 用户标识
	SubOpenID     string `xml:"sub_openid,omitempty"`     // 用户子标识
	IsSubscribe   string `xml:"is_subscribe,omitempty"`   // 是否关注公众账号
	BankType      string `xml:"bank_type,omitempty"`      // 付款银行
	TotalFee      int    `xml:"total_fee,omitempty"`      // 总金额
	FeeType       string `xml:"fee_type,omitempty"`       // 货币种类
	CashFee       int    `xml:"cash_fee,omitempty"`       // 现金支付金额
	CashFeeType   string `xml:"cash_fee_type,omitempty"`  // 现金支付货币类型
	TradeState    string `xml:"trade_state,omitempty"`    // 交易状态
	CouponFee     int    `xml:"coupon_fee,omitempty"`     // 代金券或立减优惠金额
	CouponCount   int    `xml:"coupon_count,omitempty"`   // 代金券或立减优惠使用数量
	CouponID      string `xml:"coupon_id_0,omitempty"`    // 代金券或立减优惠ID (简化，只处理第一个)
	CouponFeeN    int    `xml:"coupon_fee_0,omitempty"`   // 单个代金券或立减优惠支付金额 (简化)
	TransactionID string `xml:"transaction_id,omitempty"` // 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no,omitempty"`   // 商户订单号
	Attach        string `xml:"attach,omitempty"`         // 商家数据包
	TimeEnd       string `xml:"time_end,omitempty"`       // 支付完成时间
	ContractID    string `xml:"contract_id,omitempty"`    // 委托代扣协议id
}

// PapPayNotifyResponse 扣款结果通知响应
type PapPayNotifyResponse struct {
	ReturnCode string `xml:"return_code"`          // 返回状态码
	ReturnMsg  string `xml:"return_msg,omitempty"` // 返回信息
}
