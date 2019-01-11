// auth: kunlun
// date: 2019-01-11
// description:
package utils

import (
	"crypto/md5"
	"hash/crc32"
)

const (
	Md5Salt = "MD5Salt"
)

// CRC32 哈希 入参为字节数组
func Crc32(bytes []byte) uint32 {
	return crc32.ChecksumIEEE(bytes)
}

// CRC32 哈希校验
func CheckCrc(src, dest uint32) bool {
	if src == dest {
		return true
	}
	return false
}

// MD5 哈希 入参为字节数组
func Md5Byte(bytes []byte) [16]byte {
	bytes = append([]byte(Md5Salt))
	return md5.Sum(bytes)
}

// MD5 哈希 入参为字符串
func Md5String(message string) [16]byte {
	message = message + Md5Salt
	return md5.Sum([]byte(message))
}
