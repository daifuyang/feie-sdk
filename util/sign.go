/**
** @创建时间: 2020/9/7 9:36 上午
** @作者　　: return
** @描述　　:
 */
package util

import (
	"crypto/sha1"
	"encoding/hex"
)

// sha1加密
func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1:= hex.EncodeToString(s[:])
	return strsha1
}