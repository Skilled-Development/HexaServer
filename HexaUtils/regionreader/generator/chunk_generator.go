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
	continentalnesNoise := NewPerlinNoiseOctave(-8, 0.5, 2.0)

	// Populate the hashmap with x, z as keys and y as value
	stoneBlockState := &regionreader.BlockState{
		Name: "minecraft:stone",
	}

	airBlockState := &regionreader.BlockState{
		Name: "minecraft:air",
	}

	graniteBlockState := &regionreader.BlockState{
		Name: "minecraft:granite",
	}

	waterBlockState := &regionreader.BlockState{
		Name: "minecraft:water",
		Properties: map[string]string{
			"level": "0",
		},
	}

	grassBlockState := &regionreader.BlockState{
		Name: "minecraft:grass_block",
		Properties: map[string]string{
			"snowy": "false",
		},
	}

	// 5. Iterate through Vertical Sections (Y)
	for sectionY := int32(-4); sectionY < 20; sectionY++ {
		section := &regionreader.Section{
			Y: byte(sectionY),
			BlockStates: &regionreader.BlockStates{
				Palette: []*regionreader.BlockState{},
				Data:    make([]int64, 4096), // Changed to 4096 int64
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

		// Add granite to the palette if it's not there yet
		if !containsBlockState(section.BlockStates.Palette, graniteBlockState) {
			section.BlockStates.Palette = append(section.BlockStates.Palette, graniteBlockState)
		}
		granitePaletteIndex := indexOfBlockState(section.BlockStates.Palette, graniteBlockState)

		if !containsBlockState(section.BlockStates.Palette, waterBlockState) {
			section.BlockStates.Palette = append(section.BlockStates.Palette, waterBlockState)
		}
		waterPaletteIndex := indexOfBlockState(section.BlockStates.Palette, waterBlockState)

		if !containsBlockState(section.BlockStates.Palette, grassBlockState) {
			section.BlockStates.Palette = append(section.BlockStates.Palette, grassBlockState)
		}
		grassPaletteIndex := indexOfBlockState(section.BlockStates.Palette, grassBlockState)

		// 6. Iterate through Blocks within the Section (Y, Z, X)
		blockIndex := 0
		currentInt := int64(0)
		bitOffset := 0

		heightmap := make(map[[2]int]int)

		for y := int32(0); y < 16; y++ {
			for z := int32(0); z < 16; z++ {
				for x := int32(0); x < 16; x++ {

					//World block y position
					blockY := (sectionY * 16) + y

					// 7. Calculate Noise Value
					noiseValue := noise.Sample2D(float64(chunkX*16+x)/50, float64(chunkZ*16+z)/50)
					continentalnesNoiseValue := continentalnesNoise.Sample2D(float64(chunkX*16+x)/60, float64(chunkZ*16+z)/60)

					continentalnesHeight := 100.0
					if continentalnesNoiseValue <= 0.3 {
						m := (100.0 - 50.0) / (0.3 + 1.0)
						b := 50.0 - m*(-1.0)
						continentalnesHeight = m*continentalnesNoiseValue + b
					} else if continentalnesNoiseValue <= 0.4 {
						m := (150.0 - 100.0) / (0.4 - 0.3)
						b := 100.0 - m*(0.3)
						continentalnesHeight = m*continentalnesNoiseValue + b
					} else if continentalnesNoiseValue <= 1 {
						continentalnesHeight = 150.0
					}

					surfaceY := int(continentalnesHeight + (noiseValue * 70))

					// 8. Determine Block Type based on Noise
					var finalPaletteIndex int

					// If the noise is above this threshold, place stone
					if int(surfaceY) > int(blockY) {
						//randomNumber := rand.Intn(2)
						if blockIndex%2 == 0 {
							finalPaletteIndex = stonePaletteIndex
						} else {
							finalPaletteIndex = granitePaletteIndex
						}
						heightmap[[2]int{int(x), int(z)}] = int(surfaceY)
					} else {
						if blockY > 62 {
							finalPaletteIndex = airPaletteIndex
						} else {
							finalPaletteIndex = waterPaletteIndex
						}
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
		// Iterate through heightmap to change the heightmap at x, heightmap, z to grass, here i want to change the heightmap to grass
		for k, v := range heightmap {
			blockX := int32(k[0])
			blockZ := int32(k[1])
			surfaceY := int32(v)
			sectionBlockY := int32(sectionY * 16)
			if surfaceY >= sectionBlockY && surfaceY < sectionBlockY+16 {
				localY := surfaceY - sectionBlockY
				blockIndex := int64(localY*256 + blockZ*16 + blockX)
				index := blockIndex / 16
				offset := (blockIndex % 16) * 4
				section.BlockStates.Data[index] = (section.BlockStates.Data[index] &^ (0xF << offset)) | (int64(grassPaletteIndex) << offset)
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
