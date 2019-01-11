// package utils provider tools for developers
// auth: kunlun
// date: 2019-01-11
// description:  单元测试
package utils

import (
	"fmt"
	"testing"
)

// test PASS
func TestCrc32(t *testing.T) {
	val := Crc32([]byte("hello"))
	fmt.Println("crc32 val: ", val)
}

func BenchmarkCrc32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Crc32([]byte("hello"))

	}
}
