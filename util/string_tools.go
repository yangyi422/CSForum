package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/beinan/fastid"
)

// GetStrRequestID 获取随机字符串ID
func GetStrRequestID() string {
	var uuid = fastid.CommonConfig.GenInt64ID()
	return strconv.FormatInt(uuid, 10)
}

// AddSalt 密码加盐
func AddSalt(salt, password string) (string, string) {
	// 创建md5对象
	md5_1 := md5.New()
	// 对盐进行md5加密
	md5_1.Write([]byte(salt))
	// 将m5的hash转成[]byte格式
	m5_salt := hex.EncodeToString(md5_1.Sum(nil))
	password = password + m5_salt
	// 创建md5对象
	md5_2 := md5.New()
	// 对盐进行md5加密
	md5_2.Write([]byte(password))
	// 将m5的hash转成[]byte格式
	return salt, hex.EncodeToString(md5_2.Sum(nil))
}
