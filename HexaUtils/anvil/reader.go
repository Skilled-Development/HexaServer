package anvil

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	//"github.com/klauspost/compress/lz4"
)

const SECTOR_SIZE = 4096

// McaError representa los errores posibles durante la manipulación de regiones.
type McaError struct {
	Message string
	Code    string
}

func (e *McaError) Error() string {
	return fmt.Sprintf("MCA error %s: %s", e.Code, e.Message)
}

// NewMcaError crea un nuevo error MCA
func NewMcaError(message string, code string) *McaError {
	return &McaError{
		Message: message,
		Code:    code,
	}
}

// NotGenerated error
func NotGenerated() *McaError {
	return &McaError{
		Message: "Chunk hasn't been generated yet",
		Code:    "NotGenerated",
	}
}

// MissingHeader error
func MissingHeader() *McaError {
	return &McaError{
		Message: "No region header",
		Code:    "MissingHeader",
	}
}

// InvalidChunkPayload error
func InvalidChunkPayload(message string) *McaError {
	return &McaError{
		Message: fmt.Sprintf("Invalid chunk: %s", message),
		Code:    "InvalidChunkPayload",
	}
}

// OutOfBoundsByte error
func OutOfBoundsByte() *McaError {
	return &McaError{
		Message: "Out of bounds byte access",
		Code:    "OutOfBoundsByte",
	}
}

// IoError error
func IoError(err error) *McaError {
	return &McaError{
		Message: fmt.Sprintf("Io failed: %s", err.Error()),
		Code:    "IoError",
	}
}

// ZLibError error
func ZLibError(err error) *McaError {
	return &McaError{
		Message: fmt.Sprintf("Zlib Decompression failed: %s", err.Error()),
		Code:    "ZLibError",
	}
}

// CompressionType representa los tipos de compresión utilizados en los chunks.
type CompressionType uint8

const (
	GZip         CompressionType = 1
	Zlib         CompressionType = 2
	Uncompressed CompressionType = 3
	LZ4          CompressionType = 4
	Custom       CompressionType = 127
)

// FromUint8 convierte un u8 a un CompressionType.
func (c CompressionType) FromUint8(value uint8) CompressionType {
	switch value {
	case 1:
		return GZip
	case 2:
		return Zlib
	case 3:
		return Uncompressed
	case 4:
		return LZ4
	case 127:
		return Custom
	default:
		panic(fmt.Sprintf("Invalid compression type: %d", value))
	}
}

// ToUint8 convierte un CompressionType a un u8.
func (c CompressionType) ToUint8() uint8 {
	switch c {
	case GZip:
		return 1
	case Zlib:
		return 2
	case Uncompressed:
		return 3
	case LZ4:
		return 4
	case Custom:
		return 127
	default:
		panic("Invalid compression type")
	}
}

func (c CompressionType) String() string {
	switch c {
	case GZip:
		return "gzip"
	case Zlib:
		return "zlib"
	case Uncompressed:
		return "uncompressed"
	case LZ4:
		return "lz4"
	case Custom:
		return "custom"
	default:
		return "unknown"
	}
}

// Compress toma un slice de bytes y usa el tipo de compresión actual para comprimir los datos.
func (c CompressionType) Compress(data []byte) ([]byte, error) {
	switch c {
	case Zlib:
		compressed, err := compressZlib(data)
		if err != nil {
			return nil, fmt.Errorf("failed to compress zlib data: %w", err)
		}
		return compressed, nil

	case Uncompressed:
		return data, nil
	/*case LZ4:
	compressed, err := compressLZ4(data)
	if err != nil {
		return nil, fmt.Errorf("failed to compress lz4 data: %w", err)
	}
	return compressed, nil*/
	case GZip:
		return nil, NewMcaError("This is unused in practice and if you somehow need this, make an issue on github and i'll add it <3", "Unimplemented")
	case Custom:
		return nil, NewMcaError("Haven't implemented this and i don't personally need this but make an issue on github and i'll fix it <3", "Unimplemented")

	default:
		return nil, NewMcaError("Invalid compression type", "InvalidCompression")
	}
}

// Decompress toma un slice de bytes y usa el tipo de compresión actual para descomprimir los datos.
func (c CompressionType) Decompress(data []byte) ([]byte, error) {
	switch c {
	case Zlib:
		decompressed, err := decompressZlib(data)
		if err != nil {
			return nil, fmt.Errorf("failed to decompress zlib data: %w", err)
		}
		return decompressed, nil
	case Uncompressed:
		return data, nil
	/*case LZ4:
	decompressed, err := decompressLZ4(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress lz4 data: %w", err)
	}
	return decompressed, nil*/
	case GZip:
		return nil, NewMcaError("This is unused in practice and if you somehow need this, make an issue on github and i'll add it <3", "Unimplemented")
	case Custom:
		return nil, NewMcaError("Haven't implemented this and i don't personally need this but make an issue on github and i'll fix it <3", "Unimplemented")
	default:
		return nil, NewMcaError("Invalid compression type", "InvalidCompression")
	}
}

func compressZlib(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decompressZlib(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	decompressedData, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}

/*
func compressLZ4(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := lz4.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decompressLZ4(data []byte) ([]byte, error) {
	r := lz4.NewReader(bytes.NewReader(data))

	decompressedData, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return decompressedData, nil
}*/

// RawChunk representa un chunk comprimido.
type RawChunk struct {
	RawData         []byte
	compressionType CompressionType
}

// NewRawChunk crea un nuevo RawChunk.
func NewRawChunk(data []byte, compression CompressionType) RawChunk {
	return RawChunk{
		RawData:         data,
		compressionType: compression,
	}
}

// Decompress descomprime los datos del chunk.
func (rc *RawChunk) Decompress() ([]byte, error) {
	return rc.compressionType.Decompress(rc.RawData)
}

// GetCompressionType obtiene el tipo de compresión del chunk.
func (rc *RawChunk) GetCompressionType() CompressionType {
	return rc.compressionType
}

// PendingChunk representa un chunk pendiente de escritura.
type PendingChunk struct {
	CompressedData []byte
	Compression    CompressionType
	Timestamp      uint32
	Coordinate     [2]uint8
}

// NewPendingChunk crea un nuevo PendingChunk.
func NewPendingChunk(
	rawData []byte,
	compression CompressionType,
	timestamp uint32,
	coordinate [2]uint8,
) (*PendingChunk, error) {
	if coordinate[0] >= 32 || coordinate[1] >= 32 {
		return nil, errors.New("coordinates must be within 0-31")
	}
	compressedData, err := compression.Compress(rawData)
	if err != nil {
		return nil, err
	}

	return &PendingChunk{
		CompressedData: compressedData,
		Compression:    compression,
		Timestamp:      timestamp,
		Coordinate:     coordinate,
	}, nil
}

// RegionReader proporciona métodos para obtener slices de bytes de chunks y datos de cabecera.
type RegionReader struct {
	data []byte
}

// NewRegionReader inicializa una nueva región.
func NewRegionReader(data []byte) (*RegionReader, error) {
	if len(data) < (SECTOR_SIZE * 2) {
		return nil, MissingHeader()
	}
	return &RegionReader{data: data}, nil
}

// Inner obtiene los datos internos de la región.
func (r *RegionReader) Inner() []byte {
	return r.data
}

// chunkOffset obtiene el offset dependiendo de las coordenadas del chunk.
func (r *RegionReader) chunkOffset(x, z int) int {
	if x >= 32 {
		panic("x must be less than 32")
	}
	if z >= 32 {
		panic("z must be less than 32")
	}

	return 4 * ((x & 31) + (z&31)*32)
}

// GetChunk obtiene un RawChunk basado en sus coordenadas relativas a la región.
func (r *RegionReader) GetChunk(x, z int) (*RawChunk, error) {
	dataLen := len(r.data)
	offset := r.chunkOffset(x, z)

	location, err := r.getLocation(offset)
	if err != nil {
		return nil, err
	}
	if location == nil {
		return nil, nil // Chunk not generated
	}

	payloadOffset := int(binary.BigEndian.Uint32([]byte{0, location[0], location[1], location[2]})) * SECTOR_SIZE

	if dataLen < (payloadOffset + 4) {
		return nil, InvalidChunkPayload("Not enough data for chunk payload")
	}

	byteLength := int(binary.BigEndian.Uint32(r.data[payloadOffset : payloadOffset+4]))

	if dataLen < payloadOffset+byteLength {
		return nil, InvalidChunkPayload("Not enough data for chunk bytes")
	}

	payloadOffset += 4

	compressionType := CompressionType(r.data[payloadOffset])

	rawData := r.data[payloadOffset+1 : payloadOffset+byteLength]

	return &RawChunk{RawData: rawData, compressionType: compressionType}, nil

}

// getLocation obtiene la ubicación del payload del chunk basándose en los offsets de bytes de las coordenadas del chunk.
func (r *RegionReader) getLocation(offset int) ([]byte, error) {
	bytes := r.data[offset : offset+4]
	if len(bytes) < 4 {
		return nil, OutOfBoundsByte()
	}

	if bytes[0] == 0 && bytes[3] == 0 {
		return nil, nil
	}

	return bytes, nil

}

// getTimestamp obtiene los bytes big endian de la marca de tiempo del chunk.
func (r *RegionReader) getTimestamp(offset int) ([]byte, error) {
	offset = SECTOR_SIZE + offset
	bytes := r.data[offset : offset+4]
	if len(bytes) < 4 {
		return nil, OutOfBoundsByte()
	}
	return bytes, nil
}

// GetU32Timestamp convierte los bytes de la marca de tiempo a segundos unix epoch u32.
func (r *RegionReader) GetU32Timestamp(timestampBytes []byte) uint32 {
	return binary.BigEndian.Uint32(timestampBytes)
}

// RegionIter implementa un iterador para chunks dentro de una región.
type RegionIter struct {
	region *RegionReader
	index  int
}

// Iter devuelve un nuevo iterador de región.
func (r *RegionReader) Iter() *RegionIter {
	return &RegionIter{
		region: r,
		index:  0,
	}
}

// MAX constante del tamaño máximo de chunks dentro de una región
const MAX_CHUNK_SIZE int = 32 * 32

// getChunkCoordinate obtiene las coordenadas de un chunk basado en el índice del iterador.
func (ri *RegionIter) getChunkCoordinate(index int) (int, int) {
	return index % 32, index / 32
}

// Next implementa la iteracion
func (ri *RegionIter) Next() (*RawChunk, error) {
	if ri.index >= MAX_CHUNK_SIZE {
		return nil, nil
	}

	x, z := ri.getChunkCoordinate(ri.index)
	ri.index++
	println("Chunk at", x, z)
	chunk, err := ri.region.GetChunk(x, z)
	if err != nil {
		return nil, err
	}
	return chunk, nil
}
