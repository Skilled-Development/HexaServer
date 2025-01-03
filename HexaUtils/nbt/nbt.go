package nbt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Nbt struct {
	Name    string
	RootTag *NbtCompound
}

func (n Nbt) GetName() string {
	return n.Name
}

func NewNbt(name string, tag *NbtCompound) *Nbt {
	return &Nbt{
		Name:    name,
		RootTag: tag,
	}
}

func (n *Nbt) Write() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, COMPOUND_ID); err != nil {
		return nil, fmt.Errorf("Nbt: error writing compound ID: %w", err)
	}

	nameData, err := StringTag{Value: n.Name}.serializeData()
	if err != nil {
		return nil, fmt.Errorf("Nbt: error serializing name: %w", err)
	}
	if _, err := buf.Write(nameData); err != nil {
		return nil, fmt.Errorf("Nbt: error writing name data: %w", err)
	}

	content, err := n.RootTag.serializeContent()
	if err != nil {
		return nil, fmt.Errorf("Nbt: error serializing content: %w", err)
	}

	if _, err := buf.Write(content); err != nil {
		return nil, fmt.Errorf("Nbt: error writing content: %w", err)
	}

	return buf.Bytes(), nil
}

func (n *Nbt) WriteUnnamed() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, COMPOUND_ID); err != nil {
		return nil, fmt.Errorf("Nbt: error writing compound ID: %w", err)
	}

	content, err := n.RootTag.serializeContent()
	if err != nil {
		return nil, fmt.Errorf("Nbt: error serializing content: %w", err)
	}
	if _, err := buf.Write(content); err != nil {
		return nil, fmt.Errorf("Nbt: error writing content: %w", err)
	}

	return buf.Bytes(), nil
}

func DeserializeNbt(r io.Reader) (*Nbt, error) {
	var tagType uint8
	if err := binary.Read(r, binary.BigEndian, &tagType); err != nil {
		return nil, fmt.Errorf("DeserializeNbt: error reading root tag type: %w", err)
	}
	if tagType != COMPOUND_ID {
		return nil, NoRootCompound(tagType)
	}

	name, err := GetNbtString(r)
	if err != nil {
		return nil, fmt.Errorf("DeserializeNbt: error reading nbt name: %w", err)
	}

	rootTag, err := deserializeCompoundContent(r)
	if err != nil {
		return nil, fmt.Errorf("DeserializeNbt: error deserializing root tag: %w", err)
	}
	return &Nbt{Name: name, RootTag: rootTag}, nil
}

func DeserializeUnnamedNbt(r io.Reader) (*Nbt, error) {
	var tagType uint8
	if err := binary.Read(r, binary.BigEndian, &tagType); err != nil {
		return nil, fmt.Errorf("DeserializeUnnamedNbt: error reading root tag type: %w", err)
	}
	if tagType != COMPOUND_ID {
		return nil, NoRootCompound(tagType)
	}

	buf := make([]byte, 100)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, fmt.Errorf("DeserializeUnnamedNbt: error reading first 100 bytes: %w", err)
	}

	// Restauramos los primeros 100 bytes en el reader
	r = io.MultiReader(bytes.NewReader(buf), r)

	// Consume the first two bytes
	discardBuffer := make([]byte, 2)
	_, err2 := io.ReadFull(r, discardBuffer)
	if err2 != nil {
		// Handle the error appropriately based on whether it should be fatal or not.
		return nil, fmt.Errorf("error discarding first two bytes: %w", err2)
	}

	rootTag, err := deserializeCompoundContent(r)
	if err != nil {
		return nil, fmt.Errorf("DeserializeUnnamedNbt: error deserializing root tag: %w", err)
	}
	return &Nbt{Name: "", RootTag: rootTag}, nil
}

func (n *Nbt) GetByte(name string) (int8, bool) {
	return n.RootTag.GetByte(name)
}

func (n *Nbt) GetShort(name string) (int16, bool) {
	return n.RootTag.GetShort(name)
}

func (n *Nbt) GetInt(name string) (int32, bool) {
	return n.RootTag.GetInt(name)
}

func (n *Nbt) GetLong(name string) (int64, bool) {
	return n.RootTag.GetLong(name)
}

func (n *Nbt) GetFloat(name string) (float32, bool) {
	return n.RootTag.GetFloat(name)
}

func (n *Nbt) GetDouble(name string) (float64, bool) {
	return n.RootTag.GetDouble(name)
}

func (n *Nbt) GetBool(name string) (bool, bool) {
	return n.RootTag.GetBool(name)
}

func (n *Nbt) GetString(name string) (string, bool) {
	return n.RootTag.GetString(name)
}

func (n *Nbt) GetList(name string) ([]NbtTag, bool) {
	return n.RootTag.GetList(name)
}

func (n *Nbt) GetCompound(name string) (*NbtCompound, bool) {
	return n.RootTag.GetCompound(name)
}

func (n *Nbt) GetIntArray(name string) ([]int32, bool) {
	return n.RootTag.GetIntArray(name)
}

func (n *Nbt) GetLongArray(name string) ([]int64, bool) {
	return n.RootTag.GetLongArray(name)
}

// Helper function to create an NbtCompound from a map
func NbtCompoundFromMap(data map[string]NbtTag) *NbtCompound {
	compound := NewNbtCompound()
	for key, value := range data {
		compound.Put(key, value)
	}
	return compound
}

func NbtCompoundFromInterfaceMap(data map[string]interface{}) *NbtCompound {
	compound := NewNbtCompound()
	for key, value := range data {
		var tag NbtTag
		switch v := value.(type) {
		case int8:
			tag = ByteTag{Value: v}
		case int16:
			tag = ShortTag{Value: v}
		case int32:
			tag = IntTag{Value: v}
		case int64:
			tag = LongTag{Value: v}
		case float32:
			tag = FloatTag{Value: v}
		case float64:
			tag = DoubleTag{Value: v}
		case string:
			tag = StringTag{Value: v}
		case []byte:
			tag = ByteArrayTag{Value: v}
		case []NbtTag:
			tag = ListTag{Value: v}
		case map[string]interface{}:
			tag = CompoundTag{Value: NbtCompoundFromInterfaceMap(v)}
		case []int32:
			tag = IntArrayTag{Value: v}
		case []int64:
			tag = LongArrayTag{Value: v}
		case bool:
			tag = ByteTag{Value: 0}
			if v {
				tag = ByteTag{Value: 1}
			}
		case int:
			tag = IntTag{Value: int32(v)}
		case []string:
			listTag := ListTag{Value: make([]NbtTag, 0, len(v))}
			for _, str := range v {
				listTag.Value = append(listTag.Value, StringTag{Value: str})
			}
			tag = listTag
		case []map[string]interface{}:
			listTag := ListTag{Value: make([]NbtTag, 0, len(v))}
			for _, mapValue := range v {
				listTag.Value = append(listTag.Value, CompoundTag{Value: NbtCompoundFromInterfaceMap(mapValue)})
			}
			tag = listTag
		case StringTag:
			tag = v
		case LongTag:
			tag = v
		case IntTag:
			tag = v
		case ByteTag:
			tag = v
		case CompoundTag:
			tag = v
		case ListTag:
			tag = v
		case ShortTag:
			tag = v
		case FloatTag:
			tag = v
		case DoubleTag:
			tag = v
		case ByteArrayTag:
			tag = v
		case IntArrayTag:
			tag = v
		case LongArrayTag:
			tag = v
		default:
			panic(fmt.Sprintf("NbtCompoundFromInterfaceMap: Unsupported type: %T", v))
		}
		compound.Put(key, tag)
	}
	return compound
}

func deserializeTagData(r io.Reader, tagId uint8) (NbtTag, error) {
	switch tagId {
	case END_ID:
		return EndTag{}, nil
	case BYTE_ID:
		var value int8
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading byte: %w", err)
		}
		return ByteTag{Value: value}, nil
	case SHORT_ID:
		var value int16
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading short: %w", err)
		}
		return ShortTag{Value: value}, nil
	case INT_ID:
		var value int32
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading int: %w", err)
		}
		return IntTag{Value: value}, nil
	case LONG_ID:
		var value int64
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading long: %w", err)
		}
		return LongTag{Value: value}, nil
	case FLOAT_ID:
		var value float32
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading float: %w", err)
		}
		return FloatTag{Value: value}, nil
	case DOUBLE_ID:
		var value float64
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading double: %w", err)
		}
		return DoubleTag{Value: value}, nil
	case BYTE_ARRAY_ID:
		var length int32
		if err := binary.Read(r, binary.BigEndian, &length); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading byte array length: %w", err)
		}
		bytes := make([]byte, length)
		_, err := io.ReadFull(r, bytes)
		if err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading byte array: %w", err)
		}
		return ByteArrayTag{Value: bytes}, nil
	case STRING_ID:
		str, err := GetNbtString(r)
		if err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading string: %w", err)
		}
		return StringTag{Value: str}, nil
	case LIST_ID:
		var tagType uint8
		if err := binary.Read(r, binary.BigEndian, &tagType); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading list tag type: %w", err)
		}
		var length int32
		if err := binary.Read(r, binary.BigEndian, &length); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading list length: %w", err)
		}
		list := make([]NbtTag, 0, length)
		for i := 0; i < int(length); i++ {
			tag, err := deserializeTagData(r, tagType)
			if err != nil {
				return nil, fmt.Errorf("deserializeTagData: error reading list element: %w", err)
			}
			list = append(list, tag)
		}
		return ListTag{Value: list}, nil

	case COMPOUND_ID:
		compound, err := deserializeCompoundContent(r)
		if err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading compound: %w", err)
		}
		return CompoundTag{Value: compound}, nil
	case INT_ARRAY_ID:
		var length int32
		if err := binary.Read(r, binary.BigEndian, &length); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading int array length: %w", err)
		}
		intArray := make([]int32, 0, length)
		for i := 0; i < int(length); i++ {
			var val int32
			if err := binary.Read(r, binary.BigEndian, &val); err != nil {
				return nil, fmt.Errorf("deserializeTagData: error reading int array element: %w", err)
			}
			intArray = append(intArray, val)
		}
		return IntArrayTag{Value: intArray}, nil
	case LONG_ARRAY_ID:
		var length int32
		if err := binary.Read(r, binary.BigEndian, &length); err != nil {
			return nil, fmt.Errorf("deserializeTagData: error reading long array length: %w", err)
		}
		longArray := make([]int64, 0, length)
		for i := 0; i < int(length); i++ {
			var val int64
			if err := binary.Read(r, binary.BigEndian, &val); err != nil {
				return nil, fmt.Errorf("deserializeTagData: error reading long array element: %w", err)
			}
			longArray = append(longArray, val)
		}
		return LongArrayTag{Value: longArray}, nil
	default:
		return nil, UnknownTagId(tagId)
	}
}
