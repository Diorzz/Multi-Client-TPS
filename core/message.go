package core

import (
	"bytes"
	"errors"
	"github.com/izqui/helpers"
	"fmt"
)

type Message struct {
	Identifier byte
	TotalLength int32
	Options    []byte
	Data       []byte

	Reply chan Message
}

func NewMessage(id byte) *Message {

	return &Message{Identifier: id}
}

func (m *Message) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)

	options := helpers.FitBytesInto(m.Options, MESSAGE_OPTIONS_SIZE)
	m.TotalLength = int32(5) + int32(len(options)) + int32(len(m.Data))


	buf.WriteByte(m.Identifier)
	buf.Write(Int32ToBytes(m.TotalLength))
	buf.Write(options)
	buf.Write(m.Data)

	fmt.Println("Send Block Data Length: ", len(m.Data))
	fmt.Println("Send Total Data length: ", len(buf.Bytes()))
	return buf.Bytes(), nil

}

func (m *Message) UnmarshalBinary(d []byte) error {

	buf := bytes.NewBuffer(d)

	if len(d) < MESSAGE_OPTIONS_SIZE+MESSAGE_TYPE_SIZE {
		return errors.New("Insuficient message size")
	}
	m.Identifier = buf.Next(1)[0]
	m.TotalLength = BytesToInt(buf.Next(4))
	m.Options = helpers.StripByte(buf.Next(MESSAGE_OPTIONS_SIZE), 0)
	m.Data = buf.Next(helpers.MaxInt)

	fmt.Println("Receive Block Data Length: ", len(m.Data))
	fmt.Println("Receive Total Data length: ", len(d))
	return nil
}
