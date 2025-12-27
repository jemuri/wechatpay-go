// Copyright 2021 Tencent Inc. All rights reserved.

package contractorder

// ContractOrderRequest 支付中签约请求
type ContractOrderRequest struct {
	AppID                  string `xml:"appid"`                    // 应用ID
	MchID                  string `xml:"mch_id"`                   // 商户号
	ContractMchID          string `xml:"contract_mchid"`           // 签约商户号
	ContractAppID          string `xml:"contract_appid"`           // 签约appid
	OutTradeNo             string `xml:"out_trade_no"`             // 商户订单号
	DeviceInfo             string `xml:"device_info,omitempty"`    // 设备号
	NonceStr               string `xml:"nonce_str"`                // 随机字符串
	Body                   string `xml:"body"`                     // 商品描述
	Detail                 string `xml:"detail,omitempty"`         // 商品详情
	Attach                 string `xml:"attach,omitempty"`         // 附加数据
	NotifyURL              string `xml:"notify_url"`               // 回调通知url
	TotalFee               int    `xml:"total_fee"`                // 总金额
	SpbillCreateIP         string `xml:"spbill_create_ip"`         // 终端IP
	TimeStart              string `xml:"time_start,omitempty"`     // 交易起始时间
	TimeExpire             string `xml:"time_expire,omitempty"`    // 交易结束时间
	GoodsTag               string `xml:"goods_tag,omitempty"`      // 商品标记
	TradeType              string `xml:"trade_type"`               // 交易类型
	ProductID              string `xml:"product_id,omitempty"`     // 商品ID
	OpenID                 string `xml:"openid,omitempty"`         // 用户标识
	PlanID                 int    `xml:"plan_id"`                  // 模板id
	ContractCode           string `xml:"contract_code"`            // 签约协议号
	RequestSerial          int64  `xml:"request_serial"`           // 请求序列号
	ContractDisplayAccount string `xml:"contract_display_account"` // 用户账户展示名称
	ContractNotifyURL      string `xml:"contract_notify_url"`      // 签约信息通知url
	Sign                   string `xml:"sign"`                     // 签名
}

// ContractOrderResponse 支付中签约响应
type ContractOrderResponse struct {
	ReturnCode             string `xml:"return_code"`                        // 返回状态码
	ReturnMsg              string `xml:"return_msg,omitempty"`               // 返回信息
	ResultCode             string `xml:"result_code,omitempty"`              // 业务结果
	AppID                  string `xml:"appid,omitempty"`                    // 应用ID
	MchID                  string `xml:"mch_id,omitempty"`                   // 商户号
	NonceStr               string `xml:"nonce_str,omitempty"`                // 随机字符串
	Sign                   string `xml:"sign,omitempty"`                     // 签名
	ErrCode                string `xml:"err_code,omitempty"`                 // 错误代码
	ErrCodeDes             string `xml:"err_code_des,omitempty"`             // 错误代码描述
	ContractResultCode     string `xml:"contract_result_code,omitempty"`     // 预签约结果
	ContractErrCode        string `xml:"contract_err_code,omitempty"`        // 预签约错误代码
	ContractErrCodeDes     string `xml:"contract_err_code_des,omitempty"`    // 预签约错误描述
	PrepayID               string `xml:"prepay_id,omitempty"`                // 预支付交易会话标识
	TradeType              string `xml:"trade_type,omitempty"`               // 交易类型
	CodeURL                string `xml:"code_url,omitempty"`                 // 二维码链接
	PlanID                 int    `xml:"plan_id,omitempty"`                  // 模板id
	RequestSerial          int64  `xml:"request_serial,omitempty"`           // 请求序列号
	ContractCode           string `xml:"contract_code,omitempty"`            // 签约协议号
	ContractDisplayAccount string `xml:"contract_display_account,omitempty"` // 用户账户展示名称
	MwebURL                string `xml:"mweb_url,omitempty"`                 // 支付跳转链接
	OutTradeNo             string `xml:"out_trade_no,omitempty"`             // 商户订单号
}

// ContractNotifyRequest 签约、解约结果通知请求
type ContractNotifyRequest struct {
	ReturnCode              string `xml:"return_code"`                         // 返回状态码
	ReturnMsg               string `xml:"return_msg,omitempty"`                // 返回信息
	ResultCode              string `xml:"result_code,omitempty"`               // 业务结果
	MchID                   string `xml:"mch_id,omitempty"`                    // 商户号
	ContractCode            string `xml:"contract_code,omitempty"`             // 签约协议号
	PlanID                  string `xml:"plan_id,omitempty"`                   // 模板id
	OpenID                  string `xml:"openid,omitempty"`                    // 用户标识
	Sign                    string `xml:"sign,omitempty"`                      // 签名
	ChangeType              string `xml:"change_type,omitempty"`               // 变更类型 ADD:签约 DELETE:解约
	OperateTime             string `xml:"operate_time,omitempty"`              // 操作时间
	ContractID              string `xml:"contract_id,omitempty"`               // 委托代扣协议id
	ContractExpiredTime     string `xml:"contract_expired_time,omitempty"`     // 协议到期时间
	ContractTerminationMode int    `xml:"contract_termination_mode,omitempty"` // 协议解约方式
	RequestSerial           int64  `xml:"request_serial,omitempty"`            // 请求序列号
}

// ContractNotifyResponse 签约、解约结果通知响应
type ContractNotifyResponse struct {
	ReturnCode string `xml:"return_code"`          // 返回状态码
	ReturnMsg  string `xml:"return_msg,omitempty"` // 返回信息
}
