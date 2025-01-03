package nbt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type NbtCompound struct {
	ChildTags map[string]NbtTag
}

func NewNbtCompound() *NbtCompound {
	return &NbtCompound{
		ChildTags: make(map[string]NbtTag),
	}
}

func (c *NbtCompound) Put(name string, value NbtTag) {
	c.ChildTags[name] = value
}

func (c *NbtCompound) serializeContent() ([]byte, error) {
	buf := new(bytes.Buffer)
	for name, tag := range c.ChildTags {
		if _, err := buf.Write([]byte{tag.TypeID()}); err != nil {
			return nil, fmt.Errorf("NbtCompound: error serializing tag type: %w", err)
		}

		nameData, err := StringTag{Value: name}.serializeData()
		if err != nil {
			return nil, fmt.Errorf("NbtCompound: error serializing name: %w", err)
		}

		if _, err := buf.Write(nameData); err != nil {
			return nil, fmt.Errorf("NbtCompound: error writing serialized name: %w", err)
		}

		tagData, err := tag.serializeData()
		if err != nil {
			return nil, fmt.Errorf("NbtCompound: error serializing tag data: %w", err)
		}

		if _, err := buf.Write(tagData); err != nil {
			return nil, fmt.Errorf("NbtCompound: error writing tag data: %w", err)
		}
	}
	_, err := buf.Write([]byte{END_ID})
	if err != nil {
		return nil, fmt.Errorf("NbtCompound: error serializing end tag: %w", err)
	}
	return buf.Bytes(), nil
}

func deserializeCompoundContent(r io.Reader) (*NbtCompound, error) {
	compound := NewNbtCompound()
	for {
		var tagId uint8
		if err := binary.Read(r, binary.BigEndian, &tagId); err != nil {
			if err == io.EOF {
				fmt.Println("deserializeCompoundContent: EOF")
				break
			}
			return nil, fmt.Errorf("deserializeCompoundContent: error reading tag ID: %w", err)
		}

		if tagId == END_ID {
			break
		}

		name, err := GetNbtString(r)
		if err != nil {
			return nil, fmt.Errorf("deserializeCompoundContent: error reading tag name: %w", err)
		}

		tag, err := deserializeTagData(r, tagId)

		if err != nil {
			return nil, fmt.Errorf("deserializeCompoundContent: error reading tag data: %w", err)
		}
		compound.Put(name, tag)
	}
	return compound, nil
}

func (c *NbtCompound) GetByte(name string) (int8, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractByte()
	}
	return 0, false
}

func (c *NbtCompound) GetShort(name string) (int16, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractShort()
	}
	return 0, false
}

func (c *NbtCompound) GetInt(name string) (int32, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractInt()
	}
	return 0, false
}

func (c *NbtCompound) GetLong(name string) (int64, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractLong()
	}
	return 0, false
}

func (c *NbtCompound) GetFloat(name string) (float32, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractFloat()
	}
	return 0, false
}

func (c *NbtCompound) GetDouble(name string) (float64, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractDouble()
	}
	return 0, false
}

func (c *NbtCompound) GetBool(name string) (bool, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractBool()
	}
	return false, false
}

func (c *NbtCompound) GetString(name string) (string, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractString()
	}
	return "", false
}

func (c *NbtCompound) GetList(name string) ([]NbtTag, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractList()
	}
	return nil, false
}

func (c *NbtCompound) GetCompound(name string) (*NbtCompound, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractCompound()
	}
	return nil, false
}

func (c *NbtCompound) GetIntArray(name string) ([]int32, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractIntArray()
	}
	return nil, false
}

func (c *NbtCompound) GetLongArray(name string) ([]int64, bool) {
	if tag, ok := c.ChildTags[name]; ok {
		return tag.extractLongArray()
	}
	return nil, false
}
