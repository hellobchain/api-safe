package apisafe

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/hellobchain/api-safe/cache"
	"github.com/hellobchain/api-safe/constant"
)

// API接口参数校验
// 1. 时间校验
func checkTimestamp(timestamp int64, timeError int64) error {
	if timeError == 0 {
		timeError = constant.DEFAULT_TIME_ERR_SECOND
	}
	nowTimestamp := time.Now().UTC().UnixMilli()
	if nowTimestamp-timestamp > (timeError * 1000) {
		return fmt.Errorf("timestamp 时间差超出允许的范围，请求无效")
	}
	return nil
}

// 2. nonce 是否存在
func isExistedNonce(nonce string) bool {
	_, found := cache.GetCache().Get("nonce_" + nonce)
	return found
}

// 3. para1+para2+para3+app_secret
func sign(params string) ([]byte, error) {
	sum := sha256.Sum256([]byte(params))
	return sum[:], nil
}

// 4. 设置nonce 并有过期时间
func setNonce(nonce string, nonceExpireTime int64) {
	if nonceExpireTime == 0 {
		nonceExpireTime = constant.DEFAULT_NONCE_EXPIRE_SECOND
	}
	cache.GetCache().Set("nonce_"+nonce, nonce, nonceExpireTime)
}
