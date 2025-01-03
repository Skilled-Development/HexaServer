package nbt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	END_ID        uint8 = 0
	BYTE_ID       uint8 = 1
	SHORT_ID      uint8 = 2
	INT_ID        uint8 = 3
	LONG_ID       uint8 = 4
	FLOAT_ID      uint8 = 5
	DOUBLE_ID     uint8 = 6
	BYTE_ARRAY_ID uint8 = 7
	STRING_ID     uint8 = 8
	LIST_ID       uint8 = 9
	COMPOUND_ID   uint8 = 10
	INT_ARRAY_ID  uint8 = 11
	LONG_ARRAY_ID uint8 = 12
)

func GetNbtString(r io.Reader) (string, error) {
	var length uint16
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return "", fmt.Errorf("GetNbtString: error reading length: %w", err)
	}

	stringBytes := make([]byte, length)
	if _, err := io.ReadFull(r, stringBytes); err != nil {
		return "", fmt.Errorf("GetNbtString: error reading string: %w", err)
	}
	string, err := fromJavaCESU8(stringBytes)
	if err != nil {
		return "", fmt.Errorf("GetNbtString: error decoding CESU8: %w", err)
	}

	return string, nil
}

func WriteVarInt(buffer *bytes.Buffer, value int32) {

	for {

		temp := byte(value & 0x7F)

		value >>= 7

		if value != 0 {

			temp |= 0x80

		}

		buffer.WriteByte(temp)

		if value == 0 {

			break

		}

	}

}
