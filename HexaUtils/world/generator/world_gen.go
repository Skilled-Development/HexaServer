package generator

import (
	"HexaUtils/regionreader"
	"HexaUtils/world"
	"math/rand"
)

type WorldGenerator struct {
	Seed      int64
	WorldName string
	WorldType world.WorldType
}

func NewWorldGenerator(seed int64, worldName string, worldType world.WorldType) WorldGenerator {
	return WorldGenerator{
		Seed:      seed,
		WorldName: worldName,
		WorldType: worldType,
	}
}

func (w *WorldGenerator) GetSeed() int64 {
	return w.Seed
}

func (w *WorldGenerator) GetWorldName() string {
	return w.WorldName
}

func (w *WorldGenerator) GetWorldType() world.WorldType {
	return w.WorldType
}

func (w *WorldGenerator) SetSeed(seed int64) {
	w.Seed = seed
}

func (w *WorldGenerator) SetWorldName(worldName string) {
	w.WorldName = worldName
}

func (w *WorldGenerator) SetWorldType(worldType world.WorldType) {
	w.WorldType = worldType
}

func (w *WorldGenerator) GenerateChunk(chunkX int32, chunkZ int32) *regionreader.Chunk {
	chunk := &regionreader.Chunk{
		XPos:           chunkX,
		ZPos:           chunkZ,
		YPos:           -4,
		DataVersion:    3465,
		Status:         "full",
		LastUpdate:     0,
		InhabitedTime:  0,
		CarvingMasks:   make(map[string][]byte),
		Lights:         make([][]int16, 0),
		Entities:       make([]*regionreader.Entity, 0),
		FluidTicks:     make([]*regionreader.Tick, 0),
		BlockTicks:     make([]*regionreader.Tick, 0),
		PostProcessing: make([][]int16, 0),
		Structures: &regionreader.Structures{
			References: make(map[string][]int64),
			Starts:     make(map[string]*regionreader.Structure),
		},
		Heightmaps: map[string][]int64{
			"MOTION_BLOCKING": make([]int64, 256),
			"WORLD_SURFACE":   make([]int64, 256),
		},
		Sections:      make([]*regionreader.Section, 0),
		BlockEntities: make([]*regionreader.BlockEntity, 0),
	}

	rand.Seed(w.Seed + int64(chunkX) + int64(chunkZ)*1000)

	return chunk
}
