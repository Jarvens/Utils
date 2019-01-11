// auth: kunlun
// date: 2019-01-11
// description:
package utils

import "testing"

func BenchmarkNewDefaultProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDefaultProto("发送消息")
	}
}
