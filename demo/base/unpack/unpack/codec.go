package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_Header = "12345678"

// 自定义 tcp 协议 header + 4个字节存放body的长度 + body data
func Encode(bytesBuffer io.Writer, content string) error {

	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

func Decode(byteBuffer io.Reader) (bodyBuf []byte, err error) {
	MagicBuf := make([]byte, len(Msg_Header))
	if _, err = io.ReadFull(byteBuffer, MagicBuf); err != nil {
		return nil, err
	}
	if string(MagicBuf) != Msg_Header {
		return nil, errors.New("msg_header error")
	}
	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(byteBuffer, lengthBuf); err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)
	bodyBuf = make([]byte, length)
	if _, err = io.ReadFull(byteBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, err

}
