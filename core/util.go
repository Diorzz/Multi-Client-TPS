package core

import (
	"bytes"
	"encoding/binary"
)

func Int32ToBytes(x int32) []byte  {

	bf := bytes.NewBuffer([]byte{})
	binary.Write(bf, binary.BigEndian, x)
	return bf.Bytes()
}

func BytesToInt(b []byte) int32  {

	var x int32
	bf := bytes.NewBuffer(b)
	binary.Read(bf, binary.BigEndian, &x)

	return x
}