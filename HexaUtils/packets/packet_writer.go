package packets

import (
	"HexaUtils/nbt"
	"fmt"
	"math"

	"github.com/google/uuid"
)

// PacketWriter es un escritor de paquetes que usa un arreglo de bytes.
type PacketWriter struct {
	buffer []byte
}

func NewPacketWriterFromBuffer(buffer []byte) *PacketWriter {
	return &PacketWriter{buffer: buffer}
}

func (w *PacketWriter) GetAsPacket() *PacketWriter {
	// Obtener el buffer del paquete que contiene los datos a enviar
	buffer := w.GetPacketBuffer()
	// Crear un nuevo buffer que contiene el paquete de longitud
	packetLength := len(buffer)
	// Usamos un buffer preasignado para optimizar la asignación de memoria
	otherWriter := NewPacketWriter()
	// Usar WriteVarInt directamente en el buffer de destino
	otherWriter.WriteVarInt(int32(packetLength))

	// Usar append para agregar el buffer de forma más eficiente
	otherWriter.buffer = append(otherWriter.buffer, buffer...)

	return otherWriter
}

// GetPacketBuffer devuelve el buffer como un arreglo de bytes.
func (w *PacketWriter) GetPacketBuffer() []byte {
	return w.buffer
}

func (w *PacketWriter) WriteUUID(uuid uuid.UUID) error {
	// Convertimos el UUID de la cadena a bytes
	var msb, lsb uint64

	// Usamos la función Scan para parsear el UUID
	_, err := fmt.Sscanf(uuid.String(), "%08x-%04x-%04x-%04x-%012x",
		&msb,
		&lsb,
		&lsb,
		&lsb,
		&lsb)
	if err != nil {
		return fmt.Errorf("error al parsear el UUID: %v", err)
	}

	// Escribimos los 16 bytes en el buffer
	for i := 7; i >= 0; i-- {
		w.WriteByte(byte(msb >> (8 * i)))
	}
	for i := 7; i >= 0; i-- {
		w.WriteByte(byte(lsb >> (8 * i)))
	}

	return nil
}

// NewPacketWriter crea un nuevo escritor de paquetes con un buffer vacío.
func NewPacketWriter() *PacketWriter {
	return &PacketWriter{buffer: []byte{}}
}

// WriteByte escribe un byte en el buffer.
func (w *PacketWriter) WriteByte(b byte) {
	w.buffer = append(w.buffer, b)
}

// WriteVarInt escribe un VarInt (int32) en el buffer.
func (w *PacketWriter) WriteVarInt(value int32) {
	for {
		b := byte(value & 0x7F)
		value >>= 7

		if (value != 0) || (b&0x80 != 0) {
			b |= 0x80
		}

		w.WriteByte(b)

		if value == 0 {
			break
		}
	}
}

// WriteVarLong escribe un VarLong (int64) en el buffer.
func (w *PacketWriter) WriteVarLong(value int64) {
	for {
		b := byte(value & 0x7F)
		value >>= 7

		if (value != 0) || (b&0x80 != 0) {
			b |= 0x80
		}

		w.WriteByte(b)

		if value == 0 {
			break
		}
	}
}

// WriteString escribe un String en el buffer, con su longitud prefijada como VarInt.
func (p *PacketWriter) WriteString(value string) *PacketWriter {

	// Convertimos el string a bytes UTF-8
	valueBytes := []byte(value)

	// Calculamos la longitud en bytes
	length := len(valueBytes)

	// Verificamos que la longitud no exceda el límite
	if length > 32767*3+3 {
		panic("String is too long to be encoded as UTF-8 with a VarInt length prefix")
	}
	// Escribimos la longitud como un VarInt
	p.WriteVarInt(int32(length))

	// Escribimos los bytes de la cadena
	p.buffer = append(p.buffer, valueBytes...)

	return p
}

func (w *PacketWriter) WriteIdentifier(identifier string) *PacketWriter {
	bytes := []byte(identifier)
	w.WriteVarInt(int32(len(bytes)))
	w.buffer = append(w.buffer, bytes...)
	return w
}

func (w *PacketWriter) WriteIdentifierWithoutLength(identifier string) {
	bytes := []byte(identifier)
	w.buffer = append(w.buffer, bytes...)
}

// WriteByteArray escribe un arreglo de bytes en el buffer, con su longitud prefijada como VarInt.
func (w *PacketWriter) WriteByteArray(data []byte) {
	// Escribimos la longitud del arreglo como un VarInt
	w.WriteVarInt(int32(len(data)))

	// Escribimos los bytes del arreglo
	w.buffer = append(w.buffer, data...)
}

func (w *PacketWriter) AppendByteArray(data []byte) {
	w.buffer = append(w.buffer, data...)
}

// En packets/packet_writer.go
func (w *PacketWriter) WriteLong(value int64) {
	// Se divide el valor en 8 bytes (long de 64 bits)
	for i := 7; i >= 0; i-- {
		// Se obtiene el byte correspondiente (el más alto en cada iteración)
		b := byte(value >> (i * 8))
		// Escribimos el byte en el buffer
		w.WriteByte(b)
	}
}

func (w *PacketWriter) WriteIdentifierArray(identifiers []string) {
	// Escribimos la longitud del arreglo como un VarInt
	w.WriteVarInt(int32(len(identifiers)))

	// Escribimos cada identificador en el buffer
	for _, identifier := range identifiers {
		w.WriteIdentifier(identifier)
	}
}

func (w *PacketWriter) WriteNBT(nbt nbt.Nbt) {
	// Escribimos el contenido del NBT sin el nombre del tag raiz
	compboundBuffer, err := nbt.WriteUnnamed() // Use the unnamed serialization
	if err != nil {
		panic(fmt.Sprintf("error writing NBT: %v", err))
	}
	w.buffer = append(w.buffer, compboundBuffer...)
}

func (w *PacketWriter) WriteJson(json string) error {
	const maxJsonLength = 32767 // Límite impuesto por el servidor Notchian desde la versión 1.20.3

	if len(json) > maxJsonLength {
		return fmt.Errorf("el contenido JSON excede el límite máximo de %d caracteres", maxJsonLength)
	}

	w.WriteString(json) // Reutiliza WriteString para escribir la longitud y el contenido
	return nil
}

// WriteUnsignedByte escribe un Unsigned Byte (uint8) en el buffer.
func (w *PacketWriter) WriteUnsignedByte(value uint8) {
	w.WriteByte(byte(value))
}

// WriteBoolean escribe un Boolean (bool) en el buffer.
func (w *PacketWriter) WriteBoolean(value bool) {
	if value {
		w.WriteByte(0x01)
	} else {
		w.WriteByte(0x00)
	}
}

// WriteShort escribe un Short (int16) en el buffer.
func (w *PacketWriter) WriteShort(value int16) {
	w.WriteByte(byte(value >> 8))
	w.WriteByte(byte(value))
}

// WriteUnsignedShort escribe un Unsigned Short (uint16) en el buffer.
func (w *PacketWriter) WriteUnsignedShort(value uint16) {
	w.WriteByte(byte(value >> 8))
	w.WriteByte(byte(value))
}

// WriteInt escribe un Int (int32) en el buffer.
func (w *PacketWriter) WriteInt(value int32) {
	w.WriteByte(byte(value >> 24))
	w.WriteByte(byte(value >> 16))
	w.WriteByte(byte(value >> 8))
	w.WriteByte(byte(value))
}

// WriteFloat escribe un Float (float32) en el buffer.
func (w *PacketWriter) WriteFloat(value float32) {
	bits := math.Float32bits(value)
	w.WriteByte(byte(bits >> 24))
	w.WriteByte(byte(bits >> 16))
	w.WriteByte(byte(bits >> 8))
	w.WriteByte(byte(bits))
}

// WriteDouble escribe un Double (float64) en el buffer.
func (w *PacketWriter) WriteDouble(value float64) {
	bits := math.Float64bits(value)
	w.WriteByte(byte(bits >> 56))
	w.WriteByte(byte(bits >> 48))
	w.WriteByte(byte(bits >> 40))
	w.WriteByte(byte(bits >> 32))
	w.WriteByte(byte(bits >> 24))
	w.WriteByte(byte(bits >> 16))
	w.WriteByte(byte(bits >> 8))
	w.WriteByte(byte(bits))
}

// WritePosition escribe una posición (x, y, z) en el buffer.
func (w *PacketWriter) WritePosition(x, y, z int) {
	position := ((int64(x) & 0x3FFFFFF) << 38) | ((int64(z) & 0x3FFFFFF) << 12) | (int64(y) & 0xFFF)
	w.WriteLong(position)
}

// WriteAngle escribe un ángulo en el buffer.
func (w *PacketWriter) WriteAngle(value byte) {
	w.WriteByte(value)
}

// WriteBitSet escribe un BitSet en el buffer.
func (w *PacketWriter) WriteBitSet(bitSet []uint64) {
	w.WriteVarInt(int32(len(bitSet)))
	for _, long := range bitSet {
		w.WriteLong(int64(long))
	}
}

// WriteFixedBitSet escribe un Fixed BitSet en el buffer.
func (w *PacketWriter) WriteFixedBitSet(bitSet []byte) {
	w.buffer = append(w.buffer, bitSet...)
}

// WriteOptional escribe un valor opcional de tipo T.
func (w *PacketWriter) WriteOptional(value interface{}, writeFunc func(interface{})) {
	if value != nil {
		writeFunc(value)
	}
}

// WriteArray escribe un array de tipo T.
func (w *PacketWriter) WriteArray(array []interface{}, writeFunc func(interface{})) {
	for _, value := range array {
		writeFunc(value)
	}
}

// WriteEnum escribe un valor de enumeración.
func (w *PacketWriter) WriteEnum(value int, writeFunc func(int)) {
	writeFunc(value)
}

// WriteIDOrX escribe un ID o un valor de tipo X.
func (w *PacketWriter) WriteIDOrX(id int, value interface{}, writeValueFunc func(interface{})) {
	if id == 0 {
		w.WriteVarInt(int32(0))
		if value != nil {
			writeValueFunc(value)
		}
	} else {
		w.WriteVarInt(int32(id + 1))
	}
}

// WriteIDSet escribe un conjunto de IDs.
func (w *PacketWriter) WriteIDSet(typeID int, tagName string, ids []int) {
	w.WriteVarInt(int32(typeID))
	if typeID == 0 {
		w.WriteIdentifier(tagName)
	} else {
		for _, id := range ids {
			w.WriteVarInt(int32(id))
		}
	}
}

func (w *PacketWriter) BuildPacket() []byte {
	packetData := w.buffer
	packetLength := len(packetData)

	lengthWriter := NewPacketWriter()
	lengthWriter.WriteVarInt(int32(packetLength))

	finalBuffer := make([]byte, 0)
	finalBuffer = append(finalBuffer, lengthWriter.buffer...)
	finalBuffer = append(finalBuffer, packetData...)

	return finalBuffer
}
