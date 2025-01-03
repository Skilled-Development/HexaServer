package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// PacketReader es un lector de paquetes que usa un buffer.
type PacketReader struct {
	buffer *bytes.Buffer
}

// NewPacketReader crea un nuevo lector de paquetes.
func NewPacketReader(data []byte) *PacketReader {
	return &PacketReader{buffer: bytes.NewBuffer(data)}
}

func (r *PacketReader) SetBuffer(data []byte) {
	r.buffer = bytes.NewBuffer(data)
}

// ReadByte lee un byte del buffer.
func (r *PacketReader) ReadByte() (byte, error) {
	b, err := r.buffer.ReadByte()
	if err != nil {
		return 0, err
	}
	return b, nil
}

// ReadVarInt lee un VarInt (int32) del buffer.
func (r *PacketReader) ReadVarInt() (int32, error) {
	var value int32
	var position uint

	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		value |= int32(b&0x7F) << position
		position += 7

		if b&0x80 == 0 {
			break
		}

		if position >= 32 {
			return 0, errors.New("VarInt demasiado grande")
		}
	}

	return value, nil
}

// ReadUnsignedShort lee un unsigned short (u16) del buffer.
func (r *PacketReader) ReadUnsignedShort() (uint16, error) {
	// Verifica si hay suficientes bytes en el buffer (2 bytes)
	if r.buffer.Len() < 2 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un unsigned short")
	}

	// Lee 2 bytes del buffer
	bytes := make([]byte, 2)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convierte los 2 bytes a un unsigned short (uint16)
	// Se asume que el orden es Big Endian (como el ejemplo de Rust)
	// bytes[0] es el byte más significativo y bytes[1] es el menos significativo
	value := binary.BigEndian.Uint16(bytes)

	return value, nil
}

// ReadVarLong lee un VarLong (int64) del buffer.
func (r *PacketReader) ReadVarLong() (int64, error) {
	var value int64
	var position uint

	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		value |= int64(b&0x7F) << position
		position += 7

		if b&0x80 == 0 {
			break
		}

		if position >= 64 {
			return 0, errors.New("VarLong demasiado grande")
		}
	}

	return value, nil
}

// ReadLong lee un Long (int64) del buffer.
func (r *PacketReader) ReadLong() (int64, error) {
	// Verifica si hay suficientes bytes en el buffer (8 bytes)
	if r.buffer.Len() < 8 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un long")
	}

	// Lee 8 bytes del buffer
	bytes := make([]byte, 8)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convierte los 8 bytes a un long (int64)
	// Suponemos que el orden es Big Endian
	value := int64(binary.BigEndian.Uint64(bytes))

	return value, nil
}

// ReadUUID lee un UUID (128 bits, 16 bytes) del buffer.
func (r *PacketReader) ReadUUID() (uuid.UUID, error) {
	// Verificamos si hay suficientes bytes en el buffer (16 bytes)
	if r.buffer.Len() < 16 {
		return uuid.Nil, errors.New("no hay suficientes bytes en el buffer para leer un UUID")
	}

	// Leemos los 16 bytes del buffer
	bytes := make([]byte, 16)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return uuid.Nil, err
	}

	// Convertimos los 16 bytes a un formato UUID
	// Los primeros 8 bytes corresponden a los 64 bits más significativos, y los 8 siguientes a los menos significativos
	msb := bytes[:8] // 64 bits más significativos
	lsb := bytes[8:] // 64 bits menos significativos

	// Devolvemos el UUID en el formato estándar de string
	uuidString := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		binary.BigEndian.Uint32(msb[0:4]),
		binary.BigEndian.Uint16(msb[4:6]),
		binary.BigEndian.Uint16(msb[6:8]),
		binary.BigEndian.Uint16(lsb[0:2]),
		uint64(lsb[2])<<40|uint64(lsb[3])<<32|uint64(lsb[4])<<24|uint64(lsb[5])<<16|uint64(lsb[6])<<8|uint64(lsb[7]),
	)

	return uuid.Parse(uuidString)
}

// ReadString lee un String prefijado con su tamaño como VarInt.
func (r *PacketReader) ReadString() (string, error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return "", err
	}

	bytes := make([]byte, length)
	_, err = r.buffer.Read(bytes)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Regex para validar Identifier
var (
	namespaceRegex = regexp.MustCompile(`^[a-z0-9._-]+$`)
	valueRegex     = regexp.MustCompile(`^[a-z0-9._/-]+$`)
)

// ReadIdentifier lee un Identifier del buffer.
func (r *PacketReader) ReadIdentifier() (string, error) {
	// Leer el String con prefijo de tamaño como VarInt
	identifier, err := r.ReadString()
	if err != nil {
		return "", err
	}

	// Separar el namespace y el valor
	parts := strings.SplitN(identifier, ":", 2)
	var namespace, value string

	if len(parts) == 1 {
		// Si no se especifica el namespace, asumir "minecraft"
		namespace = "minecraft"
		value = parts[0]
	} else {
		namespace = parts[0]
		value = parts[1]
	}

	// Validar namespace y value con regex
	if !namespaceRegex.MatchString(namespace) {
		return "", fmt.Errorf("namespace inválido: %s", namespace)
	}
	if !valueRegex.MatchString(value) {
		return "", fmt.Errorf("valor inválido: %s", value)
	}

	// Devolver el Identifier en formato completo
	return fmt.Sprintf("%s:%s", namespace, value), nil
}

func (r *PacketReader) AvailableBytes() int {
	return r.buffer.Len()
}

// ReadBytes lee una cantidad específica de bytes del buffer.
func (r *PacketReader) ReadBytes(length int) ([]byte, error) {
	if length < 0 {
		return nil, errors.New("la longitud no puede ser negativa")
	}

	// Verificar si hay suficientes bytes en el buffer
	if r.buffer.Len() < length {
		return nil, errors.New("no hay suficientes bytes en el buffer")
	}

	// Leer los bytes
	bytes := make([]byte, length)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// ReadBoolean lee un booleano del buffer.
func (r *PacketReader) ReadBoolean() (bool, error) {
	b, err := r.ReadByte()
	if err != nil {
		return false, err
	}
	if b == 0x01 {
		return true, nil
	} else if b == 0x00 {
		return false, nil
	}
	return false, errors.New("valor booleano inválido")
}

// ReadUnsignedByte lee un unsigned byte (uint8) del buffer.
func (r *PacketReader) ReadUnsignedByte() (uint8, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint8(b), nil
}

// ReadShort lee un Short (int16) del buffer.
func (r *PacketReader) ReadShort() (int16, error) {
	// Verifica si hay suficientes bytes en el buffer (2 bytes)
	if r.buffer.Len() < 2 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un short")
	}

	// Lee 2 bytes del buffer
	bytes := make([]byte, 2)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convierte los 2 bytes a un short (int16)
	value := int16(binary.BigEndian.Uint16(bytes))
	return value, nil
}

// ReadInt lee un Int (int32) del buffer.
func (r *PacketReader) ReadInt() (int32, error) {
	// Verifica si hay suficientes bytes en el buffer (4 bytes)
	if r.buffer.Len() < 4 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un int")
	}

	// Lee 4 bytes del buffer
	bytes := make([]byte, 4)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convierte los 4 bytes a un int (int32)
	value := int32(binary.BigEndian.Uint32(bytes))
	return value, nil
}

// ReadFloat lee un Float (float32) del buffer.
func (r *PacketReader) ReadFloat() (float32, error) {
	// Verifica si hay suficientes bytes en el buffer (4 bytes)
	if r.buffer.Len() < 4 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un float")
	}

	// Lee 4 bytes del buffer
	bytes := make([]byte, 4)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convierte los 4 bytes a un float (float32)
	bits := binary.BigEndian.Uint32(bytes)
	value := math.Float32frombits(bits)
	return value, nil
}

// ReadDouble lee un Double (float64) del buffer.
func (r *PacketReader) ReadDouble() (float64, error) {
	// Verifica si hay suficientes bytes en el buffer (8 bytes)
	if r.buffer.Len() < 8 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un double")
	}

	// Lee 8 bytes del buffer
	bytes := make([]byte, 8)
	_, err := r.buffer.Read(bytes)
	if err != nil {
		return 0, err
	}

	// Convierte los 8 bytes a un double (float64)
	bits := binary.BigEndian.Uint64(bytes)
	value := math.Float64frombits(bits)
	return value, nil
}

// ReadPosition lee una posición (x, y, z) del buffer.
func (r *PacketReader) ReadPosition() (x, y, z int, err error) {
	val, err := r.ReadLong()
	if err != nil {
		return 0, 0, 0, err
	}

	x = int(val >> 38)
	y = int(val << 52 >> 52)
	z = int(val << 26 >> 38)
	// if x >= 1<<25 {
	// 	x -= 1 << 26
	// }
	// if y >= 1<<11 {
	// 	y -= 1 << 12
	// }
	// if z >= 1<<25 {
	// 	z -= 1 << 26
	// }

	return x, y, z, nil
}

// ReadAngle lee un ángulo del buffer.
func (r *PacketReader) ReadAngle() (byte, error) {
	angle, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	return angle, nil
}

// ReadBitSet lee un BitSet del buffer.
func (r *PacketReader) ReadBitSet() ([]uint64, error) {
	length, err := r.ReadVarInt()
	if err != nil {
		return nil, err
	}

	bitSet := make([]uint64, length)
	for i := 0; i < int(length); i++ {
		longVal, err := r.ReadLong()
		if err != nil {
			return nil, err
		}
		bitSet[i] = uint64(longVal)
	}
	return bitSet, nil
}

// ReadFixedBitSet lee un Fixed BitSet del buffer.
func (r *PacketReader) ReadFixedBitSet(length int) ([]byte, error) {
	// Calcula la longitud en bytes
	byteLength := int(math.Ceil(float64(length) / 8.0))
	// Lee los bytes
	bitSet, err := r.ReadBytes(byteLength)
	if err != nil {
		return nil, err
	}
	return bitSet, nil
}

// ReadOptional lee un valor opcional de tipo T.
func (r *PacketReader) ReadOptional(readFunc func() (interface{}, error)) (interface{}, error) {
	// La presencia o ausencia del valor debe ser conocida por el contexto del paquete.
	// Por lo tanto, esta función simplemente lee si el contexto lo requiere.
	return readFunc()
}

// ReadArray lee un array de tipo T.
func (r *PacketReader) ReadArray(readFunc func() (interface{}, error), length int) ([]interface{}, error) {
	array := make([]interface{}, length)
	for i := 0; i < length; i++ {
		value, err := readFunc()
		if err != nil {
			return nil, err
		}
		array[i] = value
	}
	return array, nil
}

// ReadEnum lee un valor de enumeración.
func (r *PacketReader) ReadEnum(readFunc func() (int, error)) (int, error) {
	return readFunc()
}

// ReadIDOrX lee un ID o un valor de tipo X.
func (r *PacketReader) ReadIDOrX(readXFunc func() (interface{}, error)) (id int, value interface{}, err error) {
	idInt, err := r.ReadVarInt()
	if err != nil {
		return 0, nil, err
	}

	id = int(idInt)
	if id == 0 {
		value, err = readXFunc()
		return id, value, err
	}

	return id - 1, nil, nil
}

// ReadIDSet lee un conjunto de IDs.
func (r *PacketReader) ReadIDSet() (typeID int, tagName string, ids []int, err error) {
	typeIDInt, err := r.ReadVarInt()
	if err != nil {
		return 0, "", nil, err
	}
	typeID = int(typeIDInt)

	if typeID == 0 {
		tagName, err = r.ReadIdentifier()
		return typeID, tagName, nil, err
	} else {
		ids = make([]int, typeID-1)
		for i := 0; i < typeID-1; i++ {
			idInt, err := r.ReadVarInt()
			if err != nil {
				return 0, "", nil, err
			}
			ids[i] = int(idInt)
		}
		return typeID, "", ids, nil
	}
}

/*
func (r *PacketReader) ReadNBT() (nbt.Nbt, error) {
	// Read all remaining bytes from the buffer
	remainingBytes := r.buffer.Bytes()

	// Create a new reader from the remaining bytes
	nbtReader := bytes.NewReader(remainingBytes)

	// Create a new NBT reader using the io.Reader
	nbtData, err := nbt.ReadUnnamed(nbtReader)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("error reading NBT: %w", err)
	}

	// Consume all remaining bytes from the buffer (because nbt.ReadNamed may not consume all bytes)
	r.buffer.Next(len(remainingBytes))

	return nbtData, nil
}*/

func (r *PacketReader) ReadJson() (string, error) {
	jsonString, err := r.ReadString()
	if err != nil {
		return "", fmt.Errorf("error reading JSON string: %w", err)
	}
	return jsonString, nil
}
