// auth: kunlun
// date: 2019-01-11
// description:
package utils

import "fmt"

type BaseProto struct {
	Magic string
	Len   uint16
	Data  string
	Crc32 uint32
}

// 对baseProto 进行编码
func (proto *BaseProto) Encoder() []byte {
	encodeBytes := append(append(append([]byte(proto.Magic),
		Uint16ToByte(proto.Len)...),
		proto.Data...),
		Uint32ToByte(proto.Crc32)...)
	return encodeBytes
}

// 解码器
func Decoder(buffer []byte, ch chan []byte) []byte {
	var i int
	len := len(buffer)
	var msgLen uint16
	for i = 0; i < len; i++ {
		if len < QuoteHeaderLen {
			fmt.Printf("decoder faild, cause protocol length less than minLength ")
			break
		}
		header := string(buffer[i:QuoteHeaderLen])
		if header == QuoteHeader {
			msgLen = ByteToUint16(buffer[i+QuoteHeaderLen : i+QuoteHeaderLen+2])
			if len <= i+QuoteHeaderLen+2+int(msgLen)+4 {
				break
			}
			data := buffer[i+QuoteHeaderLen : i+QuoteHeaderLen+int(msgLen)]
			ch <- data
			i += QuoteHeaderLen + int(msgLen) - 1
		}
	}
	if i == len {
		return make([]byte, 0)
	}
	return buffer[i+QuoteHeaderLen+2 : i+QuoteHeaderLen+2+int(msgLen)]
}

// 默认协议构造
func NewDefaultProto(message string) *BaseProto {
	crc32Val := Crc32([]byte(message))
	//crc32Val := Crc32String(message)
	return &BaseProto{Magic: QuoteHeader, Len: uint16(len([]byte(message))), Data: message, Crc32: crc32Val}
}

//指定协议头构造
func NewProto(magic, message string) *BaseProto {
	crc32Val := Crc32([]byte(message))
	return &BaseProto{Magic: magic, Len: uint16(len([]byte(message))), Data: message, Crc32: crc32Val}
}
