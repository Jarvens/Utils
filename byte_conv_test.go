// auth: kunlun
// date: 2019-01-11
// description:  单元测试
package utils

import "testing"

// test PASS
func TestIntToBytes(t *testing.T) {
	IntToBytes(10)
}

func TestUint32ToByte(t *testing.T) {
	Uint32ToByte(10)
}

func TestByteToHex(t *testing.T) {
	ByteToHex([]byte{0, 0, 0, 0, 1})
}
