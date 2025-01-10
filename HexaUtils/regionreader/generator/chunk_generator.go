package generator

import (
	"HexaUtils/regionreader"
	"math/rand"
)

func GenerateChunk(chunkX, chunkZ int32, seed int64) *regionreader.Chunk {
	// 1. Initialize Chunk Data
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

	// 2. Seed Random Generator
	rand.Seed(seed + int64(chunkX) + int64(chunkZ)*1000)

	// 3. Generate Heightmaps
	for x := int32(0); x < 16; x++ {
		for z := int32(0); z < 16; z++ {
			height := int64(60 + rand.Intn(10))
			chunk.Heightmaps["MOTION_BLOCKING"][x+z*16] = height
			chunk.Heightmaps["WORLD_SURFACE"][x+z*16] = height
		}
	}
	// 4. Initialize Simplex Noise
	noise := NewPerlinNoise()

	// Populate the hashmap with x, z as keys and y as value
	stoneBlockState := &regionreader.BlockState{
		Name: "minecraft:stone",
	}

	airBlockState := &regionreader.BlockState{
		Name: "minecraft:air",
	}

	// 5. Iterate through Vertical Sections (Y)
	for sectionY := int32(-4); sectionY < 20; sectionY++ {
		section := &regionreader.Section{
			Y: byte(sectionY),
			BlockStates: &regionreader.BlockStates{
				Palette: []*regionreader.BlockState{},
				Data:    make([]int64, 256), // Changed to 256 int64
			},
			Biomes: &regionreader.Biomes{
				Palette: []string{"minecraft:plains"},
				Data:    make([]int64, 64),
			},
			BlockLight: []byte{},
			SkyLight:   []byte{},
		}

		// Add stone to the palette if it's not there yet
		if !containsBlockState(section.BlockStates.Palette, stoneBlockState) {
			section.BlockStates.Palette = append(section.BlockStates.Palette, stoneBlockState)
		}
		stonePaletteIndex := indexOfBlockState(section.BlockStates.Palette, stoneBlockState)

		// Add air to the palette if it's not there yet
		if !containsBlockState(section.BlockStates.Palette, airBlockState) {
			section.BlockStates.Palette = append(section.BlockStates.Palette, airBlockState)
		}
		airPaletteIndex := indexOfBlockState(section.BlockStates.Palette, airBlockState)

		// 6. Iterate through Blocks within the Section (Y, Z, X)
		blockIndex := 0
		currentInt := int64(0)
		bitOffset := 0

		for y := int32(0); y < 16; y++ {
			for z := int32(0); z < 16; z++ {
				for x := int32(0); x < 16; x++ {

					//World block y position
					blockY := (sectionY * 16) + y

					// 7. Calculate Noise Value
					noiseValue := noise.Sample2D(float64(chunkX*16+x)/20, float64(chunkZ*16+z)/20)
					surfaceY := int(float64(100) + (noiseValue * 30))

					// 8. Determine Block Type based on Noise
					var finalPaletteIndex int

					// If the noise is above this threshold, place stone
					if int(surfaceY) > int(blockY) {
						finalPaletteIndex = stonePaletteIndex
					} else {
						finalPaletteIndex = airPaletteIndex
					}

					// Pack the palette index into the current int64
					currentInt |= int64(finalPaletteIndex) << bitOffset
					bitOffset += 4

					// If the current int64 is full, store it and reset
					if bitOffset == 64 {
						section.BlockStates.Data[blockIndex] = currentInt
						currentInt = 0
						bitOffset = 0
						blockIndex++
					}

				}
			}
		}

		//Store remaining bits if any
		if bitOffset > 0 {
			section.BlockStates.Data[blockIndex] = currentInt
		}

		for i := 0; i < 64; i++ {
			section.Biomes.Data[i] = int64(0)
		}

		chunk.Sections = append(chunk.Sections, section)
	}

	return chunk
}

func containsBlockState(blockStatePalette []*regionreader.BlockState, blockState *regionreader.BlockState) bool {
	for _, state := range blockStatePalette {
		if state.Name == blockState.Name {
			match := true
			if len(state.Properties) != len(blockState.Properties) {
				match = false
			} else {
				for k, v := range blockState.Properties {
					if state.Properties[k] != v {
						match = false
						break
					}
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func indexOfBlockState(blockStatePalette []*regionreader.BlockState, blockState *regionreader.BlockState) int {
	for i, state := range blockStatePalette {
		if state.Name == blockState.Name {
			match := true
			if len(state.Properties) != len(blockState.Properties) {
				match = false
			} else {
				for k, v := range blockState.Properties {
					if state.Properties[k] != v {
						match = false
						break
					}
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}
