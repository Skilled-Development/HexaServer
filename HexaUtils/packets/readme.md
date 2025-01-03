# 📦 Packet Reader & Writer 📝

This document provides a detailed explanation of the `PacketReader` and `PacketWriter` Go packages. These packages are designed to handle reading and writing data types commonly used in network protocols, especially for Minecraft packets.

## 📚 Packet Reader (`packets/packet_reader.go`)

The `PacketReader` struct is used to read data from a byte buffer. It provides methods for reading various data types, handling errors, and ensuring data integrity.

### ⚙️ Structure

```go
type PacketReader struct {
    buffer *bytes.Buffer
}
```
- `buffer`: A pointer to a `bytes.Buffer` that stores the data to be read.

### 🛠️ Methods

#### `NewPacketReader(data []byte) *PacketReader`
- **Purpose**: 🔨 Creates a new `PacketReader` instance from a byte slice.
- **Input**: `data` ([]byte): The byte slice containing the data to be read.
- **Output**: `*PacketReader`: A pointer to the new `PacketReader` instance.
- **Example**:
```go
data := []byte{0x01, 0x02, 0x03}
reader := NewPacketReader(data)
```

#### `ReadByte() (byte, error)`
- **Purpose**: 📖 Reads a single byte from the buffer.
- **Output**: `byte`: The read byte.
- **Output**: `error`: An error if reading fails.
- **Example**:
```go
b, err := reader.ReadByte()
```

#### `ReadVarInt() (int32, error)`
- **Purpose**: 🔢 Reads a variable-length integer (VarInt) from the buffer.
- **Output**: `int32`: The read integer.
- **Output**: `error`: An error if reading fails (e.g., invalid VarInt format).
- **Details**: VarInts are integers that use a variable number of bytes to save space. Smaller numbers take fewer bytes to store.
- **Example**:
```go
value, err := reader.ReadVarInt()
```

#### `ReadUnsignedShort() (uint16, error)`
- **Purpose**: 🧮 Reads an unsigned 16-bit integer (u16) from the buffer.
- **Output**: `uint16`: The read unsigned short.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 2 bytes and interprets them as a big-endian unsigned short.
- **Example**:
```go
value, err := reader.ReadUnsignedShort()
```

#### `ReadVarLong() (int64, error)`
- **Purpose**: 🔢 Reads a variable-length long integer (VarLong) from the buffer.
- **Output**: `int64`: The read long integer.
- **Output**: `error`: An error if reading fails.
- **Details**: Similar to VarInt, but for 64-bit integers.
- **Example**:
```go
value, err := reader.ReadVarLong()
```

#### `ReadLong() (int64, error)`
- **Purpose**: 🧮 Reads a 64-bit integer (int64) from the buffer.
- **Output**: `int64`: The read long integer.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 8 bytes and interprets them as a big-endian long integer.
- **Example**:
```go
value, err := reader.ReadLong()
```

#### `ReadUUID() (uuid.UUID, error)`
- **Purpose**: 🆔 Reads a UUID (Universally Unique Identifier) from the buffer.
- **Output**: `uuid.UUID`: The read UUID.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 16 bytes representing a UUID.
- **Example**:
```go
uuidValue, err := reader.ReadUUID()
```

#### `ReadString() (string, error)`
- **Purpose**: 📝 Reads a string from the buffer.
- **Output**: `string`: The read string.
- **Output**: `error`: An error if reading fails.
- **Details**: The string is prefixed by a VarInt indicating its length.
- **Example**:
```go
str, err := reader.ReadString()
```

#### `ReadIdentifier() (string, error)`
- **Purpose**: 🏷️ Reads a Minecraft Identifier (e.g., `minecraft:stone`).
- **Output**: `string`: The read identifier in `namespace:value` format.
- **Output**: `error`: An error if reading fails, including invalid namespace or value.
- **Details**: Validates the identifier format using regular expressions.
- **Example**:
```go
identifier, err := reader.ReadIdentifier()
```

#### `AvailableBytes() int`
- **Purpose**: 📏 Returns the number of bytes left in the buffer.
- **Output**: `int`: The remaining bytes.
- **Example**:
```go
available := reader.AvailableBytes()
```

#### `ReadBytes(length int) ([]byte, error)`
- **Purpose**: 📑 Reads a specific number of bytes from the buffer.
- **Input**: `length` (int): The number of bytes to read.
- **Output**: `[]byte`: The read byte slice.
- **Output**: `error`: An error if reading fails (e.g., insufficient bytes).
- **Example**:
```go
bytes, err := reader.ReadBytes(10)
```

#### `ReadBoolean() (bool, error)`
- **Purpose**: ✅ Reads a boolean value from the buffer.
- **Output**: `bool`: The read boolean value.
- **Output**: `error`: An error if reading fails (invalid boolean value).
- **Details**: Reads a single byte, where `0x01` is `true` and `0x00` is `false`.
- **Example**:
```go
boolean, err := reader.ReadBoolean()
```

#### `ReadUnsignedByte() (uint8, error)`
- **Purpose**: 🧮 Reads an unsigned 8-bit integer (uint8) from the buffer.
- **Output**: `uint8`: The read unsigned byte.
- **Output**: `error`: An error if reading fails.
- **Example**:
```go
value, err := reader.ReadUnsignedByte()
```

#### `ReadShort() (int16, error)`
- **Purpose**: 🧮 Reads a signed 16-bit integer (int16) from the buffer.
- **Output**: `int16`: The read short.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 2 bytes and interprets them as a big-endian signed short.
- **Example**:
```go
value, err := reader.ReadShort()
```

#### `ReadInt() (int32, error)`
- **Purpose**: 🧮 Reads a signed 32-bit integer (int32) from the buffer.
- **Output**: `int32`: The read int.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 4 bytes and interprets them as a big-endian signed integer.
- **Example**:
```go
value, err := reader.ReadInt()
```

#### `ReadFloat() (float32, error)`
- **Purpose**: 🔢 Reads a single-precision floating-point number (float32) from the buffer.
- **Output**: `float32`: The read float.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 4 bytes and interprets them as a big-endian float.
- **Example**:
```go
value, err := reader.ReadFloat()
```

#### `ReadDouble() (float64, error)`
- **Purpose**: 🔢 Reads a double-precision floating-point number (float64) from the buffer.
- **Output**: `float64`: The read double.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads 8 bytes and interprets them as a big-endian double.
- **Example**:
```go
value, err := reader.ReadDouble()
```

#### `ReadPosition() (x, y, z int, err error)`
- **Purpose**: 📍 Reads a 3D position (x, y, z coordinates) from the buffer.
- **Output**: `x, y, z int`: The read coordinates.
- **Output**: `err error`: An error if reading fails.
- **Details**: The position is packed into a single 64-bit integer.
- **Example**:
```go
x, y, z, err := reader.ReadPosition()
```

#### `ReadAngle() (byte, error)`
- **Purpose**: 📐 Reads an angle (byte) from the buffer.
- **Output**: `byte`: The read angle.
- **Output**: `error`: An error if reading fails.
- **Example**:
```go
angle, err := reader.ReadAngle()
```

#### `ReadBitSet() ([]uint64, error)`
- **Purpose**: 🧮 Reads a variable-length BitSet from the buffer.
- **Output**: `[]uint64`: The read BitSet (array of 64-bit integers).
- **Output**: `error`: An error if reading fails.
- **Details**: The length of the BitSet is prefixed as a VarInt.
- **Example**:
```go
bitset, err := reader.ReadBitSet()
```

#### `ReadFixedBitSet(length int) ([]byte, error)`
- **Purpose**: 🧮 Reads a fixed-length BitSet from the buffer.
- **Input**: `length` (int): The length of the BitSet in bits.
- **Output**: `[]byte`: The read BitSet.
- **Output**: `error`: An error if reading fails.
- **Details**: The length is provided as input.
- **Example**:
```go
fixedBitSet, err := reader.ReadFixedBitSet(20)
```

#### `ReadOptional(readFunc func() (interface{}, error)) (interface{}, error)`
- **Purpose**: ❓ Reads an optional value using a provided reading function.
- **Input**: `readFunc`: Function that reads the value.
- **Output**: `interface{}`: The read optional value.
- **Output**: `error`: An error if reading fails.
- **Details**: Reads a value if the context requires it, otherwise returns `nil` and no error.
- **Example**:
```go
optionalValue, err := reader.ReadOptional(func() (interface{}, error) {
    // example with ReadString
	return reader.ReadString()
})
```

#### `ReadArray(readFunc func() (interface{}, error), length int) ([]interface{}, error)`
- **Purpose**: 🗂️ Reads an array of values using a provided reading function.
- **Input**: `readFunc`: Function that reads a single element.
- **Input**: `length` (int): The number of elements in the array.
- **Output**: `[]interface{}`: The read array.
- **Output**: `error`: An error if reading fails.
- **Example**:
```go
array, err := reader.ReadArray(func() (interface{}, error) {
	return reader.ReadVarInt()
}, 10)
```

#### `ReadEnum(readFunc func() (int, error)) (int, error)`
- **Purpose**: 🚦 Reads an enum value using a provided reading function.
- **Input**: `readFunc`: Function that reads the enum value.
- **Output**: `int`: The read enum value.
- **Output**: `error`: An error if reading fails.
- **Example**:
```go
enumValue, err := reader.ReadEnum(func() (int, error) {
	return reader.ReadVarInt()
})
```

#### `ReadIDOrX(readXFunc func() (interface{}, error)) (id int, value interface{}, err error)`
- **Purpose**: 🆔 Reads either an ID or a value of type X.
- **Input**: `readXFunc`: Function that reads a value of type X.
- **Output**: `id int`: The read id.
- **Output**: `value interface{}`: The read value if ID is 0.
- **Output**: `err error`: An error if reading fails.
- **Details**: If id read from buffer is 0, it continues reading and returns a value, if not it returns nil as the value.
- **Example**:
```go
id, value, err := reader.ReadIDOrX(func() (interface{}, error){
	return reader.ReadString()
})
```

#### `ReadIDSet() (typeID int, tagName string, ids []int, err error)`
- **Purpose**: 🆔 Reads a set of IDs or a tag name from the buffer.
- **Output**: `typeID int`: The type of the id set.
- **Output**: `tagName string`: The read tag name (if typeID is 0).
- **Output**: `ids []int`: The read ids (if typeID is > 0).
- **Output**: `err error`: An error if reading fails.
- **Details**: If the typeID is 0, then a tag name is read, otherwise it reads a list of ids.
- **Example**:
```go
typeID, tagName, ids, err := reader.ReadIDSet()
```

#### `ReadJson() (string, error)`
- **Purpose**: 📜 Reads a JSON formatted string from the buffer.
- **Output**: `string`: The read JSON string.
- **Output**: `error`: An error if reading fails.
- **Details**: The JSON string is read as a prefixed string, limited to 32767 bytes.
- **Example**:
```go
jsonString, err := reader.ReadJson()
```


## ✍️ Packet Writer (`packets/packet_writer.go`)

The `PacketWriter` struct is used to write data into a byte buffer. It provides methods for writing various data types, handling memory allocation, and preparing the packet for transmission.

### ⚙️ Structure

```go
type PacketWriter struct {
	buffer []byte
}
```
- `buffer`: A slice of bytes that stores the data being written.

### 🛠️ Methods

#### `NewPacketWriter() *PacketWriter`
- **Purpose**: 🔨 Creates a new `PacketWriter` instance with an empty buffer.
- **Output**: `*PacketWriter`: A pointer to the new `PacketWriter` instance.
- **Example**:
```go
writer := NewPacketWriter()
```

#### `NewPacketWriterFromBuffer(buffer []byte) *PacketWriter`
- **Purpose**: 🔨 Creates a new `PacketWriter` instance from an existing byte slice.
- **Input**: `buffer` ([]byte): The byte slice to initialize the writer with.
- **Output**: `*PacketWriter`: A pointer to the new `PacketWriter` instance.
- **Example**:
```go
initialBuffer := []byte{0x01, 0x02, 0x03}
writer := NewPacketWriterFromBuffer(initialBuffer)
```

#### `GetAsPacket() *PacketWriter`
- **Purpose**: 📦 Prepares the packet for sending.
- **Output**: `*PacketWriter`: A new `PacketWriter` containing the packet length as a VarInt followed by the packet data itself.
- **Details**: Used to prepend the packet length as a VarInt at the beginning of the data.
- **Example**:
```go
packet := writer.GetAsPacket()
```

#### `GetPacketBuffer() []byte`
- **Purpose**: 📑 Returns the current buffer as a byte slice.
- **Output**: `[]byte`: The current state of the buffer.
- **Example**:
```go
buffer := writer.GetPacketBuffer()
```

#### `WriteUUID(uuid uuid.UUID) error`
- **Purpose**: 🆔 Writes a UUID to the buffer.
- **Input**: `uuid` (uuid.UUID): The UUID to write.
- **Output**: `error`: An error if the UUID parsing fails.
- **Details**: Writes the 16 bytes of a UUID to the buffer in big-endian order.
- **Example**:
```go
err := writer.WriteUUID(myUUID)
```

#### `WriteByte(b byte)`
- **Purpose**: 📝 Writes a single byte to the buffer.
- **Input**: `b` (byte): The byte to write.
- **Example**:
```go
writer.WriteByte(0x05)
```

#### `WriteVarInt(value int32)`
- **Purpose**: 🔢 Writes a VarInt (variable-length integer) to the buffer.
- **Input**: `value` (int32): The integer to write.
- **Details**:  VarInts are integers that use a variable number of bytes to save space. Smaller numbers take fewer bytes to store.
- **Example**:
```go
writer.WriteVarInt(150)
```

#### `WriteVarLong(value int64)`
- **Purpose**: 🔢 Writes a VarLong (variable-length long integer) to the buffer.
- **Input**: `value` (int64): The long integer to write.
- **Details**: Similar to VarInt but for 64-bit integers.
- **Example**:
```go
writer.WriteVarLong(1234567890)
```

#### `WriteString(value string) *PacketWriter`
- **Purpose**: 📝 Writes a string to the buffer, prefixed by its length as a VarInt.
- **Input**: `value` (string): The string to write.
- **Output**: `*PacketWriter`: Returns the same `PacketWriter` instance for method chaining.
- **Details**:  Converts the string to bytes and writes the bytes with length prefix.
- **Example**:
```go
writer.WriteString("Hello, World!")
```

#### `WriteIdentifier(identifier string) *PacketWriter`
- **Purpose**: 🏷️ Writes a Minecraft Identifier (e.g., `minecraft:stone`).
- **Input**: `identifier` (string): The identifier to write.
- **Output**: `*PacketWriter`: Returns the same `PacketWriter` instance for method chaining.
- **Details**: Writes the identifier string prefixed by its length as VarInt.
- **Example**:
```go
writer.WriteIdentifier("minecraft:stone")
```

#### `WriteIdentifierWithoutLength(identifier string)`
- **Purpose**: 🏷️ Writes a Minecraft Identifier without its length.
- **Input**: `identifier` (string): The identifier to write.
- **Details**: Writes the identifier string, but without its length
- **Example**:
```go
writer.WriteIdentifierWithoutLength("minecraft:stone")
```

#### `WriteByteArray(data []byte)`
- **Purpose**: 📑 Writes a byte array to the buffer, prefixed by its length as a VarInt.
- **Input**: `data` ([]byte): The byte array to write.
- **Example**:
```go
writer.WriteByteArray([]byte{0x01, 0x02, 0x03})
```

#### `AppendByteArray(data []byte)`
- **Purpose**: 📑 Appends a byte array to the buffer without writing any length information.
- **Input**: `data` ([]byte): The byte array to append.
- **Example**:
```go
writer.AppendByteArray([]byte{0x01, 0x02, 0x03})
```

#### `WriteLong(value int64)`
- **Purpose**: 🧮 Writes a 64-bit integer (int64) to the buffer.
- **Input**: `value` (int64): The long integer to write.
- **Details**: Writes the 8 bytes of a long to the buffer in big-endian order.
- **Example**:
```go
writer.WriteLong(1234567890)
```

#### `WriteIdentifierArray(identifiers []string)`
- **Purpose**: 🏷️ Writes an array of Minecraft Identifiers.
- **Input**: `identifiers` ([]string): The array of identifiers to write.
- **Details**:  Writes the length of the identifier array as a VarInt and then writes the identifiers individually, each one prefixed by its length as a VarInt.
- **Example**:
```go
writer.WriteIdentifierArray([]string{"minecraft:stone", "minecraft:dirt"})
```

#### `WriteNBT(nbt nbt.Nbt)`
- **Purpose**: 📦 Writes NBT (Named Binary Tag) data to the buffer.
- **Input**: `nbt` (nbt.Nbt): The NBT data to write.
- **Details**:  Serializes the NBT data to a byte array and appends it to the buffer without length information.
- **Example**:
```go
writer.WriteNBT(nbtData)
```

#### `WriteJson(json string) error`
- **Purpose**: 📜 Writes a JSON formatted string to the buffer.
- **Input**: `json` (string): The JSON string to write.
- **Output**: `error`: An error if the JSON string length exceeds the limit.
- **Details**: Writes the string to the buffer prefixed with its length, but limited to 32767 characters as specified by Notchian server from 1.20.3 version.
- **Example**:
```go
err := writer.WriteJson(`{"key": "value"}`)
```

#### `WriteUnsignedByte(value uint8)`
- **Purpose**: 🧮 Writes an unsigned 8-bit integer (uint8) to the buffer.
- **Input**: `value` (uint8): The unsigned byte to write.
- **Example**:
```go
writer.WriteUnsignedByte(255)
```

#### `WriteBoolean(value bool)`
- **Purpose**: ✅ Writes a boolean value to the buffer.
- **Input**: `value` (bool): The boolean value to write.
- **Details**: Writes `0x01` for `true` and `0x00` for `false`.
- **Example**:
```go
writer.WriteBoolean(true)
```

#### `WriteShort(value int16)`
- **Purpose**: 🧮 Writes a 16-bit integer (int16) to the buffer.
- **Input**: `value` (int16): The short to write.
- **Details**: Writes the 2 bytes of a short to the buffer in big-endian order.
- **Example**:
```go
writer.WriteShort(1000)
```

#### `WriteUnsignedShort(value uint16)`
- **Purpose**: 🧮 Writes an unsigned 16-bit integer (uint16) to the buffer.
- **Input**: `value` (uint16): The unsigned short to write.
- **Details**: Writes the 2 bytes of an unsigned short to the buffer in big-endian order.
- **Example**:
```go
writer.WriteUnsignedShort(65000)
```

#### `WriteInt(value int32)`
- **Purpose**: 🧮 Writes a 32-bit integer (int32) to the buffer.
- **Input**: `value` (int32): The int to write.
- **Details**: Writes the 4 bytes of an int to the buffer in big-endian order.
- **Example**:
```go
writer.WriteInt(123456)
```

#### `WriteFloat(value float32)`
- **Purpose**: 🔢 Writes a single-precision floating-point number (float32) to the buffer.
- **Input**: `value` (float32): The float to write.
- **Details**: Writes the 4 bytes of a float to the buffer in big-endian order.
- **Example**:
```go
writer.WriteFloat(3.14)
```

#### `WriteDouble(value float64)`
- **Purpose**: 🔢 Writes a double-precision floating-point number (float64) to the buffer.
- **Input**: `value` (float64): The double to write.
- **Details**: Writes the 8 bytes of a double to the buffer in big-endian order.
- **Example**:
```go
writer.WriteDouble(3.14159)
```

#### `WritePosition(x, y, z int)`
- **Purpose**: 📍 Writes a 3D position (x, y, z coordinates) to the buffer.
- **Input**: `x, y, z` (int): The x, y, z coordinates.
- **Details**: Packs the x, y, and z values into a single long and writes the long to the buffer.
- **Example**:
```go
writer.WritePosition(10, 20, 30)
```

#### `WriteAngle(value byte)`
- **Purpose**: 📐 Writes an angle (byte) to the buffer.
- **Input**: `value` (byte): The angle to write.
- **Example**:
```go
writer.WriteAngle(90)
```

#### `WriteBitSet(bitSet []uint64)`
- **Purpose**: 🧮 Writes a variable-length BitSet to the buffer.
- **Input**: `bitSet` ([]uint64): The BitSet (array of 64-bit integers) to write.
- **Details**:  Writes the length of the BitSet as a VarInt, and then writes the BitSet data as longs.
- **Example**:
```go
writer.WriteBitSet([]uint64{0x01, 0x02, 0x03})
```

#### `WriteFixedBitSet(bitSet []byte)`
- **Purpose**: 🧮 Writes a fixed-length BitSet to the buffer.
- **Input**: `bitSet` ([]byte): The BitSet to write.
- **Details**: Writes the byte array as a fixed-length BitSet.
- **Example**:
```go
writer.WriteFixedBitSet([]byte{0x01, 0x02, 0x03})
```

#### `WriteOptional(value interface{}, writeFunc func(interface{}))`
- **Purpose**: ❓ Writes an optional value.
- **Input**: `value` (interface{}): The optional value to write.
- **Input**: `writeFunc`: Function that writes a single optional value.
- **Details**: Writes a value if present using a provided writing function, otherwise, it does not write anything.
- **Example**:
```go
writer.WriteOptional("Test", func(value interface{}){
	writer.WriteString(value.(string))
})
```

#### `WriteArray(array []interface{}, writeFunc func(interface{}))`
- **Purpose**: 🗂️ Writes an array of values.
- **Input**: `array` ([]interface{}): The array of values to write.
- **Input**: `writeFunc`: Function that writes a single element of the array.
- **Details**: Iterates through the array and writes each element using the provided writing function.
- **Example**:
```go
writer.WriteArray([]interface{}{"a", "b"}, func(value interface{}){
	writer.WriteString(value.(string))
})
```

#### `WriteEnum(value int, writeFunc func(int))`
- **Purpose**: 🚦 Writes an enum value.
- **Input**: `value` (int): The enum value to write.
- **Input**: `writeFunc`: Function that writes the enum value.
- **Details**: Writes the value using the provided function.
- **Example**:
```go
writer.WriteEnum(5, func(value int){
	writer.WriteVarInt(int32(value))
})
```

#### `WriteIDOrX(id int, value interface{}, writeValueFunc func(interface{}))`
- **Purpose**: 🆔 Writes either an ID or a value of type X.
- **Input**: `id` (int): The ID or 0 to write the value.
- **Input**: `value` (interface{}): The value to write if id is 0.
- **Input**: `writeValueFunc`: Function that writes the value of type X.
- **Details**: If ID is 0, it writes 0 followed by a value of type X using the function to write a value. If the ID is greater than 0, it writes ID + 1.
- **Example**:
```go
writer.WriteIDOrX(0, "test", func(value interface{}){
	writer.WriteString(value.(string))
})
```

#### `WriteIDSet(typeID int, tagName string, ids []int)`
- **Purpose**: 🆔 Writes a set of IDs or a tag name.
- **Input**: `typeID` (int): The type of the id set.
- **Input**: `tagName` (string): The tag name to write (if typeID is 0).
- **Input**: `ids` ([]int): The list of ids to write (if typeID is > 0).
- **Details**: Writes a typeID. If the typeID is 0, then it writes the tag name. If not, it writes the list of ids.
- **Example**:
```go
writer.WriteIDSet(0, "test", nil) // Tag name
writer.WriteIDSet(3, "", []int{1,2}) // List of ids
```

#### `BuildPacket() []byte`
- **Purpose**: 📦 Builds the final packet including the packet length as a VarInt.
- **Output**: `[]byte`: The final packet data to be sent over the network.
- **Details**: Prepends the packet data length as a VarInt to the current buffer, then returns it as a byte array.
- **Example**:
```go
packet := writer.BuildPacket()
```
