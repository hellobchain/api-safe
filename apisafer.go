package apisafe

import (
	"bytes"
	"fmt"
)

type APICheckParam struct {
	Timestamp       int64  // 时间戳
	TimeError       int64  // 时间误差  在这个误差内都合理
	Nonce           string // 随机数
	NonceExpireTime int64  // nonce过期时间 秒为单位
	Params          string // 参数组合
	Sign            []byte // 签名
}

// new APICheckParam
func NewAPICheckParam(timestamp int64, timeError int64, nonce string, nonceExpireTime int64, params string, sign []byte) *APICheckParam {
	return &APICheckParam{
		Timestamp:       timestamp,
		TimeError:       timeError,
		Nonce:           nonce,
		NonceExpireTime: nonceExpireTime,
		Params:          params,
		Sign:            sign,
	}
}

// API参数校验
func (a APICheckParam) APICheck() error {
	err := checkTimestamp(a.Timestamp, a.TimeError)
	if err != nil {
		return fmt.Errorf("请求过期,错误:%v", err)
	}
	isExisted := isExistedNonce(a.Nonce)
	if isExisted {
		return fmt.Errorf("此nonce已被使用过了,请求无效")
	}
	sign, err := sign(a.Params)
	if err != nil {
		return fmt.Errorf("生成sign失败,错误:%v", err)
	}
	if !bytes.Equal(sign, a.Sign) {
		return fmt.Errorf("无效sign,无法响应请求")
	}
	setNonce(a.Nonce, a.NonceExpireTime)
	return nil
}
