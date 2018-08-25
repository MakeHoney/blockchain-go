package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	// num을 big endian 순서로 buff에 binary 형식으로 쓴다.
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	// buff에 쓰인 binary 데이터를 byte형식(16진수 형식)으로 리턴
	// .Bytes() => returns type []byte
	return buff.Bytes()
}
