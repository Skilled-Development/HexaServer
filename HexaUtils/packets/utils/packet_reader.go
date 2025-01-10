package utils

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
	// Buffer temporal para evitar asignaciones constantes.
	temp []byte
}

// NewPacketReader crea un nuevo lector de paquetes.
func NewPacketReader(data []byte) *PacketReader {
	// Inicializamos un temp con un tamaño suficiente para la lectura más grande (16 bytes para UUID)
	return &PacketReader{buffer: bytes.NewBuffer(data), temp: make([]byte, 16)}
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
	if r.buffer.Len() < 2 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un unsigned short")
	}

	// Lee directamente al buffer temporal
	_, err := r.buffer.Read(r.temp[:2])
	if err != nil {
		return 0, err
	}

	value := binary.BigEndian.Uint16(r.temp[:2])
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
	if r.buffer.Len() < 8 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un long")
	}

	_, err := r.buffer.Read(r.temp[:8])
	if err != nil {
		return 0, err
	}
	value := int64(binary.BigEndian.Uint64(r.temp[:8]))
	return value, nil
}

// ReadUUID lee un UUID (128 bits, 16 bytes) del buffer.
func (r *PacketReader) ReadUUID() (uuid.UUID, error) {
	if r.buffer.Len() < 16 {
		return uuid.Nil, errors.New("no hay suficientes bytes en el buffer para leer un UUID")
	}

	// Lee los 16 bytes al buffer temporal
	_, err := r.buffer.Read(r.temp)
	if err != nil {
		return uuid.Nil, err
	}

	msb := r.temp[:8] // 64 bits más significativos
	lsb := r.temp[8:] // 64 bits menos significativos

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

	// Usar el buffer subyacente directamente para evitar copiar la string
	buf := r.buffer.Next(int(length))

	if len(buf) != int(length) {
		return "", errors.New("no hay suficientes bytes en el buffer para leer la string")
	}

	return string(buf), nil
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

	if r.buffer.Len() < length {
		return nil, errors.New("no hay suficientes bytes en el buffer")
	}

	// Usar el buffer subyacente directamente
	bytes := r.buffer.Next(length)
	if len(bytes) != length {
		return nil, errors.New("no hay suficientes bytes en el buffer")
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
	if r.buffer.Len() < 2 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un short")
	}

	_, err := r.buffer.Read(r.temp[:2])
	if err != nil {
		return 0, err
	}
	value := int16(binary.BigEndian.Uint16(r.temp[:2]))
	return value, nil
}

// ReadInt lee un Int (int32) del buffer.
func (r *PacketReader) ReadInt() (int32, error) {
	if r.buffer.Len() < 4 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un int")
	}

	_, err := r.buffer.Read(r.temp[:4])
	if err != nil {
		return 0, err
	}

	value := int32(binary.BigEndian.Uint32(r.temp[:4]))
	return value, nil
}

// ReadFloat lee un Float (float32) del buffer.
func (r *PacketReader) ReadFloat() (float32, error) {
	if r.buffer.Len() < 4 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un float")
	}

	_, err := r.buffer.Read(r.temp[:4])
	if err != nil {
		return 0, err
	}

	bits := binary.BigEndian.Uint32(r.temp[:4])
	value := math.Float32frombits(bits)
	return value, nil
}

// ReadDouble lee un Double (float64) del buffer.
func (r *PacketReader) ReadDouble() (float64, error) {
	if r.buffer.Len() < 8 {
		return 0, errors.New("no hay suficientes bytes en el buffer para leer un double")
	}

	_, err := r.buffer.Read(r.temp[:8])
	if err != nil {
		return 0, err
	}
	bits := binary.BigEndian.Uint64(r.temp[:8])
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
	byteLength := int(math.Ceil(float64(length) / 8.0))
	bitSet, err := r.ReadBytes(byteLength)
	if err != nil {
		return nil, err
	}
	return bitSet, nil
}

// ReadOptional lee un valor opcional de tipo T.
func (r *PacketReader) ReadOptional(readFunc func() (interface{}, error)) (interface{}, error) {
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
