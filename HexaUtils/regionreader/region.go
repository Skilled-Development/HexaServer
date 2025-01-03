package regionreader

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	sectorSize = 4096
	headerSize = 8192
)

// Region represents a Minecraft region file
type Region struct {
	file   *os.File
	Header *Header
}

// OpenRegion opens and parses a region file.
func OpenRegion(filePath string) (*Region, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening region file: %w", err)
	}

	header, err := readHeader(file)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error reading region header: %w", err)
	}

	return &Region{
		file:   file,
		Header: header,
	}, nil
}

// Close closes the region file
func (r *Region) Close() error {
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}

func readHeader(file *os.File) (*Header, error) {
	header := &Header{}

	// Read location table (first 4KB)
	if err := binary.Read(file, binary.BigEndian, &header.Locations); err != nil {
		return nil, fmt.Errorf("error reading location table: %w", err)
	}

	// Read timestamp table (second 4KB)
	if err := binary.Read(file, binary.BigEndian, &header.Timestamps); err != nil {
		return nil, fmt.Errorf("error reading timestamp table: %w", err)
	}

	return header, nil
}

// ReadChunkData reads and returns the raw byte data of the specified chunk
func (r *Region) ReadChunkData(chunkX, chunkZ int) ([]byte, error) {
	if chunkX < 0 || chunkX >= 32 || chunkZ < 0 || chunkZ >= 32 {
		return nil, errors.New("chunk coordinates out of range")
	}

	index := (chunkX & 31) + (chunkZ&31)*32
	locationEntry := r.Header.Locations[index]

	//log.Printf("Reading chunk data at: X=%d, Z=%d, Location entry: %d", chunkX, chunkZ, locationEntry)

	offset := locationEntry >> 8       // Sector offset (3 bytes)
	sectorCount := byte(locationEntry) // Sector count (1 byte)
	if offset == 0 && sectorCount == 0 {
		//log.Printf("Chunk at (%d,%d) not present in region file", chunkX, chunkZ)
		return nil, errors.New("chunk not present in region file")
	}

	// Convert offset from sectors to byte offset
	byteOffset := int64(offset) * sectorSize

	// Seek to the chunk data
	_, err := r.file.Seek(byteOffset, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("error seeking to chunk offset: %w", err)
	}

	// Read chunk data length (4 bytes)
	var length int32
	if err := binary.Read(r.file, binary.BigEndian, &length); err != nil {
		return nil, fmt.Errorf("error reading chunk data length: %w", err)
	}

	if length == 0 {
		log.Printf("Chunk data length is zero at: X=%d, Z=%d, returning empty data", chunkX, chunkZ)
		return []byte{}, nil
	}

	if length <= 0 {
		return nil, errors.New("invalid chunk data length")
	}
	// Read compression type (1 byte)
	var compressionType byte
	if err := binary.Read(r.file, binary.BigEndian, &compressionType); err != nil {
		return nil, fmt.Errorf("error reading compression type: %w", err)
	}

	// Read compressed chunk data (length -1 bytes)
	compressedData := make([]byte, length-1)
	_, err = io.ReadFull(r.file, compressedData)
	if err != nil {
		return nil, fmt.Errorf("error reading compressed chunk data: %w", err)
	}
	/*compressedDataHex := fmt.Sprintf("%x", compressedData)
	debugger.PrintForDebug("Compressed data (hex): %s", compressedDataHex)*/
	switch compressionType {
	case 2:
		decompressedData, err := DecompressZlib(compressedData)
		if err != nil {
			return nil, fmt.Errorf("error decompressing zlib data: %w", err)
		}
		return decompressedData, nil
	default:
		return nil, fmt.Errorf("unsupported compression type: %d", compressionType)
	}
}

// ReadChunk reads and returns chunk data from the region file.
func (r *Region) ReadChunk(chunkX, chunkZ int) (*Chunk, error) {
	//log.Printf("Reading chunk at X=%d, Z=%d", chunkX, chunkZ)
	chunkData, err := r.ReadChunkData(chunkX, chunkZ)

	if err != nil {
		return nil, fmt.Errorf("error reading chunk data: %w", err)
	}

	nbtData, err := ReadChunkNBT(chunkData)
	if err != nil {
		return nil, fmt.Errorf("error parsing NBT: %w", err)
	}

	chunk := parseChunkData(nbtData)
	if chunk == nil {
		return nil, errors.New("error parsing chunk data")
	}

	return chunk, nil
}

// GetChunkLocation returns the location of the chunk within the region
func (r *Region) GetChunkLocation(chunkX, chunkZ int) (*ChunkLocation, error) {
	if chunkX < 0 || chunkX >= 32 || chunkZ < 0 || chunkZ >= 32 {
		return nil, errors.New("chunk coordinates out of range")
	}

	index := (chunkX & 31) + (chunkZ&31)*32
	locationEntry := r.Header.Locations[index]

	offset := locationEntry >> 8       // Sector offset (3 bytes)
	sectorCount := byte(locationEntry) // Sector count (1 byte)

	return &ChunkLocation{
		Offset:      int32(offset),
		SectorCount: sectorCount,
	}, nil
}

// GetChunkTimestamp returns the timestamp of when the chunk was last updated
func (r *Region) GetChunkTimestamp(chunkX, chunkZ int) (int32, error) {
	if chunkX < 0 || chunkX >= 32 || chunkZ < 0 || chunkZ >= 32 {
		return 0, errors.New("chunk coordinates out of range")
	}

	index := (chunkX & 31) + (chunkZ&31)*32
	return r.Header.Timestamps[index], nil
}
