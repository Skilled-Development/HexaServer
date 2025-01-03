package regionreader

// Header represents the 8KB header of a region file
type Header struct {
	Locations  [1024]int32
	Timestamps [1024]int32
}

// ChunkLocation represents a location of a chunk in the region file
type ChunkLocation struct {
	Offset      int32
	SectorCount byte
}
