package nbt

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type NbtTagType uint8

type NbtTag interface {
	TypeID() uint8
	Serialize() ([]byte, error)
	serializeData() ([]byte, error)

	extractByte() (int8, bool)
	extractShort() (int16, bool)
	extractInt() (int32, bool)
	extractLong() (int64, bool)
	extractFloat() (float32, bool)
	extractDouble() (float64, bool)
	extractBool() (bool, bool)
	extractByteArray() ([]byte, bool)
	extractString() (string, bool)
	extractList() ([]NbtTag, bool)
	extractCompound() (*NbtCompound, bool)
	extractIntArray() ([]int32, bool)
	extractLongArray() ([]int64, bool)
}

type EndTag struct{}

func (t EndTag) TypeID() uint8                  { return END_ID }
func (t EndTag) Serialize() ([]byte, error)     { return []byte{END_ID}, nil }
func (t EndTag) serializeData() ([]byte, error) { return []byte{}, nil }

func (t EndTag) extractByte() (int8, bool)             { return 0, false }
func (t EndTag) extractShort() (int16, bool)           { return 0, false }
func (t EndTag) extractInt() (int32, bool)             { return 0, false }
func (t EndTag) extractLong() (int64, bool)            { return 0, false }
func (t EndTag) extractFloat() (float32, bool)         { return 0, false }
func (t EndTag) extractDouble() (float64, bool)        { return 0, false }
func (t EndTag) extractBool() (bool, bool)             { return false, false }
func (t EndTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t EndTag) extractString() (string, bool)         { return "", false }
func (t EndTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t EndTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t EndTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t EndTag) extractLongArray() ([]int64, bool)     { return nil, false }

type ByteTag struct{ Value int8 }

func (t ByteTag) TypeID() uint8 { return BYTE_ID }
func (t ByteTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{BYTE_ID}, data...), nil
}
func (t ByteTag) serializeData() ([]byte, error) {
	return []byte{byte(t.Value)}, nil
}
func (t ByteTag) extractByte() (int8, bool)             { return t.Value, true }
func (t ByteTag) extractShort() (int16, bool)           { return 0, false }
func (t ByteTag) extractInt() (int32, bool)             { return 0, false }
func (t ByteTag) extractLong() (int64, bool)            { return 0, false }
func (t ByteTag) extractFloat() (float32, bool)         { return 0, false }
func (t ByteTag) extractDouble() (float64, bool)        { return 0, false }
func (t ByteTag) extractBool() (bool, bool)             { return t.Value != 0, true }
func (t ByteTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t ByteTag) extractString() (string, bool)         { return "", false }
func (t ByteTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t ByteTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t ByteTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t ByteTag) extractLongArray() ([]int64, bool)     { return nil, false }

type ShortTag struct{ Value int16 }

func (t ShortTag) TypeID() uint8 { return SHORT_ID }
func (t ShortTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{SHORT_ID}, data...), nil
}
func (t ShortTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.Value); err != nil {
		return nil, fmt.Errorf("ShortTag: error serializing data: %w", err)
	}
	return buf.Bytes(), nil
}
func (t ShortTag) extractByte() (int8, bool)             { return 0, false }
func (t ShortTag) extractShort() (int16, bool)           { return t.Value, true }
func (t ShortTag) extractInt() (int32, bool)             { return 0, false }
func (t ShortTag) extractLong() (int64, bool)            { return 0, false }
func (t ShortTag) extractFloat() (float32, bool)         { return 0, false }
func (t ShortTag) extractDouble() (float64, bool)        { return 0, false }
func (t ShortTag) extractBool() (bool, bool)             { return false, false }
func (t ShortTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t ShortTag) extractString() (string, bool)         { return "", false }
func (t ShortTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t ShortTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t ShortTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t ShortTag) extractLongArray() ([]int64, bool)     { return nil, false }

type IntTag struct{ Value int32 }

func (t IntTag) TypeID() uint8 { return INT_ID }
func (t IntTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{INT_ID}, data...), nil
}
func (t IntTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.Value); err != nil {
		return nil, fmt.Errorf("IntTag: error serializing data: %w", err)
	}
	return buf.Bytes(), nil
}
func (t IntTag) extractByte() (int8, bool)             { return 0, false }
func (t IntTag) extractShort() (int16, bool)           { return 0, false }
func (t IntTag) extractInt() (int32, bool)             { return t.Value, true }
func (t IntTag) extractLong() (int64, bool)            { return 0, false }
func (t IntTag) extractFloat() (float32, bool)         { return 0, false }
func (t IntTag) extractDouble() (float64, bool)        { return 0, false }
func (t IntTag) extractBool() (bool, bool)             { return false, false }
func (t IntTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t IntTag) extractString() (string, bool)         { return "", false }
func (t IntTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t IntTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t IntTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t IntTag) extractLongArray() ([]int64, bool)     { return nil, false }

type LongTag struct{ Value int64 }

func (t LongTag) TypeID() uint8 { return LONG_ID }
func (t LongTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{LONG_ID}, data...), nil
}
func (t LongTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.Value); err != nil {
		return nil, fmt.Errorf("LongTag: error serializing data: %w", err)
	}
	return buf.Bytes(), nil
}
func (t LongTag) extractByte() (int8, bool)             { return 0, false }
func (t LongTag) extractShort() (int16, bool)           { return 0, false }
func (t LongTag) extractInt() (int32, bool)             { return 0, false }
func (t LongTag) extractLong() (int64, bool)            { return t.Value, true }
func (t LongTag) extractFloat() (float32, bool)         { return 0, false }
func (t LongTag) extractDouble() (float64, bool)        { return 0, false }
func (t LongTag) extractBool() (bool, bool)             { return false, false }
func (t LongTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t LongTag) extractString() (string, bool)         { return "", false }
func (t LongTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t LongTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t LongTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t LongTag) extractLongArray() ([]int64, bool)     { return nil, false }

type FloatTag struct{ Value float32 }

func (t FloatTag) TypeID() uint8 { return FLOAT_ID }
func (t FloatTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{FLOAT_ID}, data...), nil
}
func (t FloatTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.Value); err != nil {
		return nil, fmt.Errorf("FloatTag: error serializing data: %w", err)
	}
	return buf.Bytes(), nil
}
func (t FloatTag) extractByte() (int8, bool)             { return 0, false }
func (t FloatTag) extractShort() (int16, bool)           { return 0, false }
func (t FloatTag) extractInt() (int32, bool)             { return 0, false }
func (t FloatTag) extractLong() (int64, bool)            { return 0, false }
func (t FloatTag) extractFloat() (float32, bool)         { return t.Value, true }
func (t FloatTag) extractDouble() (float64, bool)        { return 0, false }
func (t FloatTag) extractBool() (bool, bool)             { return false, false }
func (t FloatTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t FloatTag) extractString() (string, bool)         { return "", false }
func (t FloatTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t FloatTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t FloatTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t FloatTag) extractLongArray() ([]int64, bool)     { return nil, false }

type DoubleTag struct{ Value float64 }

func (t DoubleTag) TypeID() uint8 { return DOUBLE_ID }
func (t DoubleTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{DOUBLE_ID}, data...), nil
}
func (t DoubleTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, t.Value); err != nil {
		return nil, fmt.Errorf("DoubleTag: error serializing data: %w", err)
	}
	return buf.Bytes(), nil
}

func (t DoubleTag) extractByte() (int8, bool)             { return 0, false }
func (t DoubleTag) extractShort() (int16, bool)           { return 0, false }
func (t DoubleTag) extractInt() (int32, bool)             { return 0, false }
func (t DoubleTag) extractLong() (int64, bool)            { return 0, false }
func (t DoubleTag) extractFloat() (float32, bool)         { return 0, false }
func (t DoubleTag) extractDouble() (float64, bool)        { return t.Value, true }
func (t DoubleTag) extractBool() (bool, bool)             { return false, false }
func (t DoubleTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t DoubleTag) extractString() (string, bool)         { return "", false }
func (t DoubleTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t DoubleTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t DoubleTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t DoubleTag) extractLongArray() ([]int64, bool)     { return nil, false }

type ByteArrayTag struct{ Value []byte }

func (t ByteArrayTag) TypeID() uint8 { return BYTE_ARRAY_ID }
func (t ByteArrayTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{BYTE_ARRAY_ID}, data...), nil
}
func (t ByteArrayTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, int32(len(t.Value))); err != nil {
		return nil, fmt.Errorf("ByteArrayTag: error serializing length: %w", err)
	}
	_, err := buf.Write(t.Value)
	if err != nil {
		return nil, fmt.Errorf("ByteArrayTag: error serializing bytes: %w", err)
	}
	return buf.Bytes(), nil
}
func (t ByteArrayTag) extractByte() (int8, bool)             { return 0, false }
func (t ByteArrayTag) extractShort() (int16, bool)           { return 0, false }
func (t ByteArrayTag) extractInt() (int32, bool)             { return 0, false }
func (t ByteArrayTag) extractLong() (int64, bool)            { return 0, false }
func (t ByteArrayTag) extractFloat() (float32, bool)         { return 0, false }
func (t ByteArrayTag) extractDouble() (float64, bool)        { return 0, false }
func (t ByteArrayTag) extractBool() (bool, bool)             { return false, false }
func (t ByteArrayTag) extractByteArray() ([]byte, bool)      { return t.Value, true }
func (t ByteArrayTag) extractString() (string, bool)         { return "", false }
func (t ByteArrayTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t ByteArrayTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t ByteArrayTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t ByteArrayTag) extractLongArray() ([]int64, bool)     { return nil, false }

type StringTag struct{ Value string }

func (t StringTag) TypeID() uint8 { return STRING_ID }
func (t StringTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{STRING_ID}, data...), nil
}
func (t StringTag) serializeData() ([]byte, error) {
	javaString := toJavaCESU8(t.Value)
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, uint16(len(javaString))); err != nil {
		return nil, fmt.Errorf("StringTag: error serializing length: %w", err)
	}
	_, err := buf.Write(javaString)
	if err != nil {
		return nil, fmt.Errorf("StringTag: error serializing string: %w", err)
	}
	return buf.Bytes(), nil
}
func (t StringTag) extractByte() (int8, bool)             { return 0, false }
func (t StringTag) extractShort() (int16, bool)           { return 0, false }
func (t StringTag) extractInt() (int32, bool)             { return 0, false }
func (t StringTag) extractLong() (int64, bool)            { return 0, false }
func (t StringTag) extractFloat() (float32, bool)         { return 0, false }
func (t StringTag) extractDouble() (float64, bool)        { return 0, false }
func (t StringTag) extractBool() (bool, bool)             { return false, false }
func (t StringTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t StringTag) extractString() (string, bool)         { return t.Value, true }
func (t StringTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t StringTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t StringTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t StringTag) extractLongArray() ([]int64, bool)     { return nil, false }

type ListTag struct{ Value []NbtTag }

func (t ListTag) TypeID() uint8 { return LIST_ID }
func (t ListTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{LIST_ID}, data...), nil
}
func (t ListTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	var tagType uint8
	if len(t.Value) > 0 {
		tagType = t.Value[0].TypeID()
	}
	if err := binary.Write(buf, binary.BigEndian, tagType); err != nil {
		return nil, fmt.Errorf("ListTag: error serializing tag type: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, int32(len(t.Value))); err != nil {
		return nil, fmt.Errorf("ListTag: error serializing length: %w", err)
	}
	for _, tag := range t.Value {
		tagData, err := tag.serializeData()
		if err != nil {
			return nil, fmt.Errorf("ListTag: error serializing list element: %w", err)
		}
		_, err = buf.Write(tagData)
		if err != nil {
			return nil, fmt.Errorf("ListTag: error writing list element: %w", err)
		}
	}
	return buf.Bytes(), nil
}
func (t ListTag) extractByte() (int8, bool)             { return 0, false }
func (t ListTag) extractShort() (int16, bool)           { return 0, false }
func (t ListTag) extractInt() (int32, bool)             { return 0, false }
func (t ListTag) extractLong() (int64, bool)            { return 0, false }
func (t ListTag) extractFloat() (float32, bool)         { return 0, false }
func (t ListTag) extractDouble() (float64, bool)        { return 0, false }
func (t ListTag) extractBool() (bool, bool)             { return false, false }
func (t ListTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t ListTag) extractString() (string, bool)         { return "", false }
func (t ListTag) extractList() ([]NbtTag, bool)         { return t.Value, true }
func (t ListTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t ListTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t ListTag) extractLongArray() ([]int64, bool)     { return nil, false }

type CompoundTag struct{ Value *NbtCompound }

func (t CompoundTag) TypeID() uint8 { return COMPOUND_ID }
func (t CompoundTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{COMPOUND_ID}, data...), nil
}
func (t CompoundTag) serializeData() ([]byte, error) {
	return t.Value.serializeContent()
}

func (t CompoundTag) extractByte() (int8, bool)             { return 0, false }
func (t CompoundTag) extractShort() (int16, bool)           { return 0, false }
func (t CompoundTag) extractInt() (int32, bool)             { return 0, false }
func (t CompoundTag) extractLong() (int64, bool)            { return 0, false }
func (t CompoundTag) extractFloat() (float32, bool)         { return 0, false }
func (t CompoundTag) extractDouble() (float64, bool)        { return 0, false }
func (t CompoundTag) extractBool() (bool, bool)             { return false, false }
func (t CompoundTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t CompoundTag) extractString() (string, bool)         { return "", false }
func (t CompoundTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t CompoundTag) extractCompound() (*NbtCompound, bool) { return t.Value, true }
func (t CompoundTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t CompoundTag) extractLongArray() ([]int64, bool)     { return nil, false }

type IntArrayTag struct{ Value []int32 }

func (t IntArrayTag) TypeID() uint8 { return INT_ARRAY_ID }
func (t IntArrayTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{INT_ARRAY_ID}, data...), nil
}
func (t IntArrayTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, int32(len(t.Value))); err != nil {
		return nil, fmt.Errorf("IntArrayTag: error serializing length: %w", err)
	}

	for _, intVal := range t.Value {
		if err := binary.Write(buf, binary.BigEndian, intVal); err != nil {
			return nil, fmt.Errorf("IntArrayTag: error serializing int: %w", err)
		}
	}
	return buf.Bytes(), nil
}
func (t IntArrayTag) extractByte() (int8, bool)             { return 0, false }
func (t IntArrayTag) extractShort() (int16, bool)           { return 0, false }
func (t IntArrayTag) extractInt() (int32, bool)             { return 0, false }
func (t IntArrayTag) extractLong() (int64, bool)            { return 0, false }
func (t IntArrayTag) extractFloat() (float32, bool)         { return 0, false }
func (t IntArrayTag) extractDouble() (float64, bool)        { return 0, false }
func (t IntArrayTag) extractBool() (bool, bool)             { return false, false }
func (t IntArrayTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t IntArrayTag) extractString() (string, bool)         { return "", false }
func (t IntArrayTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t IntArrayTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t IntArrayTag) extractIntArray() ([]int32, bool)      { return t.Value, true }
func (t IntArrayTag) extractLongArray() ([]int64, bool)     { return nil, false }

type LongArrayTag struct{ Value []int64 }

func (t LongArrayTag) TypeID() uint8 { return LONG_ARRAY_ID }
func (t LongArrayTag) Serialize() ([]byte, error) {
	data, err := t.serializeData()
	if err != nil {
		return nil, err
	}
	return append([]byte{LONG_ARRAY_ID}, data...), nil
}
func (t LongArrayTag) serializeData() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, int32(len(t.Value))); err != nil {
		return nil, fmt.Errorf("LongArrayTag: error serializing length: %w", err)
	}
	for _, longVal := range t.Value {
		if err := binary.Write(buf, binary.BigEndian, longVal); err != nil {
			return nil, fmt.Errorf("LongArrayTag: error serializing long: %w", err)
		}
	}
	return buf.Bytes(), nil
}

func (t LongArrayTag) extractByte() (int8, bool)             { return 0, false }
func (t LongArrayTag) extractShort() (int16, bool)           { return 0, false }
func (t LongArrayTag) extractInt() (int32, bool)             { return 0, false }
func (t LongArrayTag) extractLong() (int64, bool)            { return 0, false }
func (t LongArrayTag) extractFloat() (float32, bool)         { return 0, false }
func (t LongArrayTag) extractDouble() (float64, bool)        { return 0, false }
func (t LongArrayTag) extractBool() (bool, bool)             { return false, false }
func (t LongArrayTag) extractByteArray() ([]byte, bool)      { return nil, false }
func (t LongArrayTag) extractString() (string, bool)         { return "", false }
func (t LongArrayTag) extractList() ([]NbtTag, bool)         { return nil, false }
func (t LongArrayTag) extractCompound() (*NbtCompound, bool) { return nil, false }
func (t LongArrayTag) extractIntArray() ([]int32, bool)      { return nil, false }
func (t LongArrayTag) extractLongArray() ([]int64, bool)     { return t.Value, true }
