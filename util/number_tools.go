package util

import "github.com/beinan/fastid"

// GetRequestID 获取int64随机ID
func GetRequestID() int64 {
	return fastid.CommonConfig.GenInt64ID()
}
