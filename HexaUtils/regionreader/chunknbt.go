package regionreader

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"HexaUtils/nbt"
	debugger "HexaUtils/utils"

	"github.com/google/uuid"
	google_uuid "github.com/google/uuid"
)

func ReadChunkNBT(data []byte) (*nbt.Nbt, error) {
	r := bytes.NewReader(data)
	nbtData, err := nbt.DeserializeUnnamedNbt(r)
	if err != nil {
		return nil, fmt.Errorf("error deserializing NBT: %w", err)
	}
	return nbtData, nil
}

func parseChunkData(nbtData *nbt.Nbt) *Chunk {
	if nbtData == nil || nbtData.RootTag == nil {
		return nil
	}

	chunk := &Chunk{
		CarvingMasks: make(map[string][]byte),
		Heightmaps:   make(map[string][]int64),
		Structures: &Structures{
			References: make(map[string][]int64),
			Starts:     make(map[string]*Structure),
		},
	}

	if xPos, ok := nbtData.GetInt("xPos"); ok {
		chunk.XPos = xPos
	}
	if zPos, ok := nbtData.GetInt("zPos"); ok {
		chunk.ZPos = zPos
	}
	if yPos, ok := nbtData.GetInt("yPos"); ok {
		chunk.YPos = yPos
	}
	if dataVersion, ok := nbtData.GetInt("DataVersion"); ok {
		chunk.DataVersion = dataVersion
	}

	if status, ok := nbtData.GetString("Status"); ok {
		chunk.Status = status
	}

	if lastUpdate, ok := nbtData.GetLong("LastUpdate"); ok {
		chunk.LastUpdate = lastUpdate
	}

	if sectionsList, ok := nbtData.GetList("sections"); ok {
		chunk.Sections = make([]*Section, 0)
		for _, sectionTag := range sectionsList {
			if compoundTag, ok := sectionTag.(nbt.CompoundTag); ok {
				section := parseSection(compoundTag.Value)
				if section != nil {
					chunk.Sections = append(chunk.Sections, section)
				}
			}
		}
	}

	if blockEntitiesList, ok := nbtData.GetList("block_entities"); ok {
		chunk.BlockEntities = make([]*BlockEntity, 0)
		for _, blockEntityTag := range blockEntitiesList {
			if compoundTag, ok := blockEntityTag.(nbt.CompoundTag); ok {
				blockEntity := parseBlockEntity(compoundTag.Value)
				if blockEntity != nil {
					chunk.BlockEntities = append(chunk.BlockEntities, blockEntity)
				}
			}
		}
	}

	if carvingMasksCompound, ok := nbtData.GetCompound("CarvingMasks"); ok {
		for maskName, maskTag := range carvingMasksCompound.ChildTags {
			if byteArrayTag, ok := maskTag.(nbt.ByteArrayTag); ok {
				chunk.CarvingMasks[maskName] = byteArrayTag.Value
			}
		}
	}

	if heightmapsCompound, ok := nbtData.GetCompound("Heightmaps"); ok {
		for heightmapName, heightmapTag := range heightmapsCompound.ChildTags {
			if intArrayTag, ok := heightmapTag.(nbt.LongArrayTag); ok {
				chunk.Heightmaps[heightmapName] = intArrayTag.Value
			}
		}
	}

	if lightsList, ok := nbtData.GetList("Lights"); ok {
		chunk.Lights = make([][]int16, len(lightsList))
		for i, lightListTag := range lightsList {
			if lightList, ok := lightListTag.(nbt.ListTag); ok {
				chunk.Lights[i] = make([]int16, len(lightList.Value))
				for j, pos := range lightList.Value {
					if shortTag, ok := pos.(nbt.ShortTag); ok {
						chunk.Lights[i][j] = shortTag.Value
					}
				}
			}
		}
	}

	if entitiesList, ok := nbtData.GetList("Entities"); ok {
		chunk.Entities = make([]*Entity, 0)
		for _, entityTag := range entitiesList {
			if compoundTag, ok := entityTag.(nbt.CompoundTag); ok {
				entity := parseEntity(compoundTag.Value)
				if entity != nil {
					chunk.Entities = append(chunk.Entities, entity)
				}
			}
		}
	}

	if fluidTicksList, ok := nbtData.GetList("fluid_ticks"); ok {
		chunk.FluidTicks = make([]*Tick, 0)
		for _, tickTag := range fluidTicksList {
			if compoundTag, ok := tickTag.(nbt.CompoundTag); ok {
				tick := parseTick(compoundTag.Value)
				if tick != nil {
					chunk.FluidTicks = append(chunk.FluidTicks, tick)
				}
			}
		}
	}

	if blockTicksList, ok := nbtData.GetList("block_ticks"); ok {
		chunk.BlockTicks = make([]*Tick, 0)
		for _, tickTag := range blockTicksList {
			if compoundTag, ok := tickTag.(nbt.CompoundTag); ok {
				tick := parseTick(compoundTag.Value)
				if tick != nil {
					chunk.BlockTicks = append(chunk.BlockTicks, tick)
				}
			}
		}
	}

	if inhabitedTime, ok := nbtData.GetLong("InhabitedTime"); ok {
		chunk.InhabitedTime = inhabitedTime
	}

	if postProcessingList, ok := nbtData.GetList("PostProcessing"); ok {
		chunk.PostProcessing = make([][]int16, len(postProcessingList))
		for i, postProcessingTag := range postProcessingList {
			if listTag, ok := postProcessingTag.(nbt.ListTag); ok {
				chunk.PostProcessing[i] = make([]int16, len(listTag.Value))
				for j, pos := range listTag.Value {
					if shortTag, ok := pos.(nbt.ShortTag); ok {
						chunk.PostProcessing[i][j] = shortTag.Value
					}
				}
			}
		}
	}

	if structuresCompound, ok := nbtData.GetCompound("structures"); ok {
		parseStructures(structuresCompound, chunk.Structures)
	}

	return chunk
}

func parseStructures(structuresCompound *nbt.NbtCompound, structures *Structures) {

	if referencesCompound, ok := structuresCompound.GetCompound("References"); ok {
		for refName, refTag := range referencesCompound.ChildTags {
			if longArrayTag, ok := refTag.(nbt.LongArrayTag); ok {
				structures.References[refName] = longArrayTag.Value
			}
		}
	}

	if startsCompound, ok := structuresCompound.GetCompound("starts"); ok {
		for startName, startTag := range startsCompound.ChildTags {
			if compoundTag, ok := startTag.(nbt.CompoundTag); ok {
				start := parseStructure(compoundTag.Value)
				structures.Starts[startName] = start
			}
		}
	}
}

func parseStructure(compound *nbt.NbtCompound) *Structure {
	structure := &Structure{}

	if bb, ok := compound.GetIntArray("BB"); ok {
		structure.BB = bb
	}
	if biome, ok := compound.GetString("biome"); ok {
		structure.Biome = biome
	}

	if childrenList, ok := compound.GetList("Children"); ok {
		structure.Children = make([]*StructurePiece, 0)
		for _, childTag := range childrenList {
			if compoundTag, ok := childTag.(nbt.CompoundTag); ok {
				piece := parseStructurePiece(compoundTag.Value)
				if piece != nil {
					structure.Children = append(structure.Children, piece)
				}
			}
		}
	}

	if chunkX, ok := compound.GetInt("ChunkX"); ok {
		structure.ChunkX = chunkX
	}
	if chunkZ, ok := compound.GetInt("ChunkZ"); ok {
		structure.ChunkZ = chunkZ
	}
	if id, ok := compound.GetString("id"); ok {
		structure.Id = id
	}

	if valid, ok := compound.GetBool("Valid"); ok {
		structure.Valid = valid
	}

	if processedList, ok := compound.GetList("Processed"); ok {
		structure.Processed = make([]ChunkPosition, 0)
		for _, processedTag := range processedList {
			if compoundTag, ok := processedTag.(nbt.CompoundTag); ok {
				pos := parseChunkPosition(compoundTag.Value)
				if pos != nil {
					structure.Processed = append(structure.Processed, *pos)
				}
			}
		}
	}

	return structure
}
func parseChunkPosition(compound *nbt.NbtCompound) *ChunkPosition {
	pos := &ChunkPosition{}
	if x, ok := compound.GetInt("X"); ok {
		pos.X = x
	}
	if z, ok := compound.GetInt("Z"); ok {
		pos.Z = z
	}
	return pos
}

func parseStructurePiece(compound *nbt.NbtCompound) *StructurePiece {
	piece := &StructurePiece{}
	if bb, ok := compound.GetIntArray("BB"); ok {
		piece.BB = bb
	}
	if biomeType, ok := compound.GetString("BiomeType"); ok {
		piece.BiomeType = biomeType
	}
	if cCompound, ok := compound.GetCompound("C"); ok {
		piece.C = parseBlockState(cCompound)
	}
	if caCompound, ok := compound.GetCompound("CA"); ok {
		piece.CA = parseBlockState(caCompound)
	}
	if cbCompound, ok := compound.GetCompound("CB"); ok {
		piece.CB = parseBlockState(cbCompound)
	}
	if ccCompound, ok := compound.GetCompound("CC"); ok {
		piece.CC = parseBlockState(ccCompound)
	}
	if cdCompound, ok := compound.GetCompound("CD"); ok {
		piece.CD = parseBlockState(cdCompound)
	}

	if chest, ok := compound.GetBool("Chest"); ok {
		piece.Chest = chest
	}

	if d, ok := compound.GetString("D"); ok {
		piece.D = d
	}
	if depth, ok := compound.GetInt("Depth"); ok {
		piece.Depth = depth
	}

	if entrancesList, ok := compound.GetList("Entrances"); ok {
		piece.Entrances = make([]*BB, 0)
		for _, entranceTag := range entrancesList {
			if compoundTag, ok := entranceTag.(nbt.CompoundTag); ok {
				bb := parseBB(compoundTag.Value)
				if bb != nil {
					piece.Entrances = append(piece.Entrances, bb)
				}
			}
		}
	}

	if entryDoor, ok := compound.GetString("EntryDoor"); ok {
		piece.EntryDoor = entryDoor
	}

	if gd, ok := compound.GetInt("GD"); ok {
		piece.GD = gd
	}

	if hasPlacedChest0, ok := compound.GetBool("hasPlacedChest0"); ok {
		piece.HasPlacedChest0 = hasPlacedChest0
	}
	if hasPlacedChest1, ok := compound.GetBool("hasPlacedChest1"); ok {
		piece.HasPlacedChest1 = hasPlacedChest1
	}
	if hasPlacedChest2, ok := compound.GetBool("hasPlacedChest2"); ok {
		piece.HasPlacedChest2 = hasPlacedChest2
	}
	if hasPlacedChest3, ok := compound.GetBool("hasPlacedChest3"); ok {
		piece.HasPlacedChest3 = hasPlacedChest3
	}

	if height, ok := compound.GetInt("Height"); ok {
		piece.Height = height
	}
	if hPos, ok := compound.GetInt("HPos"); ok {
		piece.HPos = hPos
	}

	if hps, ok := compound.GetBool("hps"); ok {
		piece.Hps = hps
	}

	if hr, ok := compound.GetBool("hr"); ok {
		piece.Hr = hr
	}

	if id, ok := compound.GetString("id"); ok {
		piece.Id = id
	}

	if integrity, ok := compound.GetFloat("integrity"); ok {
		piece.Integrity = integrity
	}
	if isLarge, ok := compound.GetBool("isLarge"); ok {
		piece.IsLarge = isLarge
	}

	if junctionsList, ok := compound.GetList("junctions"); ok {
		piece.Junctions = make([]*Junction, 0)
		for _, junctionTag := range junctionsList {
			if compoundTag, ok := junctionTag.(nbt.CompoundTag); ok {
				junction := parseJunction(compoundTag.Value)
				if junction != nil {
					piece.Junctions = append(piece.Junctions, junction)
				}
			}
		}
	}

	if left, ok := compound.GetBool("Left"); ok {
		piece.Left = left
	}
	if leftHigh, ok := compound.GetBool("leftHigh"); ok {
		piece.LeftHigh = leftHigh
	}
	if leftLow, ok := compound.GetBool("leftLow"); ok {
		piece.LeftLow = leftLow
	}
	if length, ok := compound.GetInt("Length"); ok {
		piece.Length = length
	}
	if mob, ok := compound.GetBool("Mob"); ok {
		piece.Mob = mob
	}
	if num, ok := compound.GetInt("Num"); ok {
		piece.Num = num
	}
	if o, ok := compound.GetInt("O"); ok {
		piece.O = o
	}

	if placedHiddenChest, ok := compound.GetBool("placedHiddenChest"); ok {
		piece.PlacedHiddenChest = placedHiddenChest
	}
	if placedMainChest, ok := compound.GetBool("placedMainChest"); ok {
		piece.PlacedMainChest = placedMainChest
	}
	if placedTrap1, ok := compound.GetBool("placedTrap1"); ok {
		piece.PlacedTrap1 = placedTrap1
	}
	if placedTrap2, ok := compound.GetBool("placedTrap2"); ok {
		piece.PlacedTrap2 = placedTrap2
	}

	if posX, ok := compound.GetInt("PosX"); ok {
		piece.PosX = posX
	}
	if posY, ok := compound.GetInt("PosY"); ok {
		piece.PosY = posY
	}
	if posZ, ok := compound.GetInt("PosZ"); ok {
		piece.PosZ = posZ
	}

	if right, ok := compound.GetBool("Right"); ok {
		piece.Right = right
	}

	if rightHigh, ok := compound.GetBool("rightHigh"); ok {
		piece.RightHigh = rightHigh
	}
	if rightLow, ok := compound.GetBool("rightLow"); ok {
		piece.RightLow = rightLow
	}

	if rot, ok := compound.GetString("Rot"); ok {
		piece.Rot = rot
	}
	if sc, ok := compound.GetBool("sc"); ok {
		piece.Sc = sc
	}

	if seed, ok := compound.GetLong("Seed"); ok {
		piece.Seed = seed
	}
	if source, ok := compound.GetBool("Source"); ok {
		piece.Source = source
	}

	if steps, ok := compound.GetInt("Steps"); ok {
		piece.Steps = steps
	}

	if t, ok := compound.GetInt("T"); ok {
		piece.T = t
	}

	if tall, ok := compound.GetBool("Tall"); ok {
		piece.Tall = tall
	}
	if template, ok := compound.GetString("Template"); ok {
		piece.Template = template
	}

	if terrace, ok := compound.GetBool("Terrace"); ok {
		piece.Terrace = terrace
	}

	if tf, ok := compound.GetBool("tf"); ok {
		piece.Tf = tf
	}
	if tpx, ok := compound.GetInt("TPX"); ok {
		piece.TPX = tpx
	}
	if tpy, ok := compound.GetInt("TPY"); ok {
		piece.TPY = tpy
	}
	if tpz, ok := compound.GetInt("TPZ"); ok {
		piece.TPZ = tpz
	}

	if tpz, ok := compound.GetInt("Type"); ok {
		piece.Type = tpz
	}
	if vCount, ok := compound.GetInt("VCount"); ok {
		piece.VCount = vCount
	}
	if width, ok := compound.GetInt("Width"); ok {
		piece.Width = width
	}
	if witch, ok := compound.GetBool("Witch"); ok {
		piece.Witch = witch
	}

	if zombie, ok := compound.GetBool("Zombie"); ok {
		piece.Zombie = zombie
	}

	return piece
}
func parseBB(compound *nbt.NbtCompound) *BB {
	bb := &BB{}
	if minX, ok := compound.GetInt("MinX"); ok {
		bb.MinX = minX
	}
	if minY, ok := compound.GetInt("MinY"); ok {
		bb.MinY = minY
	}
	if minZ, ok := compound.GetInt("MinZ"); ok {
		bb.MinZ = minZ
	}
	if maxX, ok := compound.GetInt("MaxX"); ok {
		bb.MaxX = maxX
	}
	if maxY, ok := compound.GetInt("MaxY"); ok {
		bb.MaxY = maxY
	}
	if maxZ, ok := compound.GetInt("MaxZ"); ok {
		bb.MaxZ = maxZ
	}
	return bb
}

func parseJunction(compound *nbt.NbtCompound) *Junction {
	junction := &Junction{}
	if sourceX, ok := compound.GetInt("source_x"); ok {
		junction.SourceX = sourceX
	}
	if sourceGroundY, ok := compound.GetInt("source_ground_y"); ok {
		junction.SourceGroundY = sourceGroundY
	}
	if sourceZ, ok := compound.GetInt("source_z"); ok {
		junction.SourceZ = sourceZ
	}
	if deltaY, ok := compound.GetInt("delta_y"); ok {
		junction.DeltaY = deltaY
	}
	if destProj, ok := compound.GetString("dest_proj"); ok {
		junction.DestProj = destProj
	}
	return junction
}

func parseBlockState(compound *nbt.NbtCompound) *BlockState {
	state := &BlockState{}

	if name, ok := compound.GetString("Name"); ok {
		state.Name = name
	}
	if propertiesCompound, ok := compound.GetCompound("Properties"); ok {
		state.Properties = make(map[string]string)
		for key, valueTag := range propertiesCompound.ChildTags {
			if stringTag, ok := valueTag.(nbt.StringTag); ok {
				state.Properties[key] = stringTag.Value
			}
		}
	}

	return state
}

func parseSection(compound *nbt.NbtCompound) *Section {
	section := &Section{
		BlockStates: &BlockStates{},
		Biomes:      &Biomes{},
	}

	debugger.PrintForDebug("Section deserializer")

	if y, ok := compound.GetByte("Y"); ok {
		section.Y = byte(y) // Corrected type conversion
	}
	if blockStatesCompound, ok := compound.GetCompound("block_states"); ok {
		if paletteList, ok := blockStatesCompound.GetList("palette"); ok {
			section.BlockStates.Palette = make([]*BlockState, 0)
			for _, paletteTag := range paletteList {
				if compoundTag, ok := paletteTag.(nbt.CompoundTag); ok {
					blockState := parseBlockState(compoundTag.Value)
					if blockState != nil {
						section.BlockStates.Palette = append(section.BlockStates.Palette, blockState)
					}
				}
			}
		}
		if dataArray, ok := blockStatesCompound.GetLongArray("data"); ok {
			section.BlockStates.Data = dataArray
			/*
				// Generate block position to state strings
				blockStrings := make([]string, 0, 4096)
				bitsPerIndex := calculateBitsPerIndex(len(section.BlockStates.Palette))
				for y := 0; y < 16; y++ {
					for z := 0; z < 16; z++ {
						for x := 0; x < 16; x++ {
							index := calculateIndex(x, y, z)
							paletteIndex := getPaletteIndexFromData(dataArray, index, bitsPerIndex)

							var blockName string
							if paletteIndex >= 0 && paletteIndex < len(section.BlockStates.Palette) {
								blockName = section.BlockStates.Palette[paletteIndex].Name
							} else {
								blockName = "unknown" // Or handle it as an error if necessary
							}
							blockString := fmt.Sprintf("[%d,%d,%d]-%s", x, y, z, blockName)
							blockStrings = append(blockStrings, blockString)
						}
					}
				}
				// Print the result
				debugger.PrintForDebug("Block states:")
				for _, s := range blockStrings {
					debugger.PrintForDebug(s)
				}*/
		}
	}

	if biomesCompound, ok := compound.GetCompound("biomes"); ok {
		debugger.PrintForDebug("---------- Found biomes")
		if paletteList, ok := biomesCompound.GetList("palette"); ok {
			debugger.PrintForDebug("---------- Found palette")
			section.Biomes.Palette = make([]string, 0)
			for _, paletteTag := range paletteList {
				debugger.PrintForDebug("---------- Found palette tag")
				debugger.PrintForDebug(fmt.Sprintf("%v", paletteTag))
				elementTag := paletteTag.(nbt.StringTag)
				elemntString := elementTag.Value
				section.Biomes.Palette = append(section.Biomes.Palette, elemntString)
				/*if compoundTag, ok := paletteTag.(nbt.CompoundTag); ok {
					if name, ok := compoundTag.Value.GetString("Name"); ok {
						debugger.PrintForDebug("---------- Found name")
						debugger.PrintForDebug(name)
						section.Biomes.Palette = append(section.Biomes.Palette, name)
					}
				}*/
			}
		}
		if dataArray, ok := biomesCompound.GetLongArray("data"); ok {
			debugger.PrintForDebug("---------- Found data")
			section.Biomes.Data = dataArray
		}
	}

	if blockLightTag, ok := compound.ChildTags["BlockLight"]; ok {
		if byteArrayTag, ok := blockLightTag.(nbt.ByteArrayTag); ok {
			section.BlockLight = byteArrayTag.Value
		}
	}

	if skyLightTag, ok := compound.ChildTags["SkyLight"]; ok {
		if byteArrayTag, ok := skyLightTag.(nbt.ByteArrayTag); ok {
			section.SkyLight = byteArrayTag.Value
		}
	}

	return section
}

func calculateBitsPerIndex(paletteSize int) int {
	if paletteSize <= 1 {
		return 0 // If pallete has no blocks, no bit are used for indexing
	}
	bits := 0
	for i := 1; i < paletteSize; i <<= 1 {
		bits++
	}
	if bits < 4 {
		return 4 //Minimum of 4 bits
	}
	return bits
}

func calculateIndex(x, y, z int) int {
	return y*256 + z*16 + x
}

func getPaletteIndexFromData(data []int64, index int, bitsPerIndex int) int {
	// Calculate which long in the data array contains the index we want
	longIndex := (index * bitsPerIndex) / 64

	// Calculate the bit offset within that long
	bitOffset := (index * bitsPerIndex) % 64

	var mask int64 = (1 << bitsPerIndex) - 1

	// Get the relevant long
	value := data[longIndex]

	// Apply the mask and shift to get the palette index
	paletteIndex := (value >> bitOffset) & mask

	return int(paletteIndex)
}

func parseBlockEntity(compound *nbt.NbtCompound) *BlockEntity {
	blockEntity := &BlockEntity{
		NbtData: compound,
	}
	if id, ok := compound.GetString("id"); ok {
		blockEntity.Id = id
	}
	if x, ok := compound.GetInt("x"); ok {
		blockEntity.X = x
	}
	if y, ok := compound.GetInt("y"); ok {
		blockEntity.Y = y
	}
	if z, ok := compound.GetInt("z"); ok {
		blockEntity.Z = z
	}

	return blockEntity
}

func parseEntity(compound *nbt.NbtCompound) *Entity {
	entity := &Entity{
		NbtData: compound,
	}
	if id, ok := compound.GetString("id"); ok {
		entity.Id = id
	}
	if pos, ok := compound.GetList("Pos"); ok {
		entity.Pos = make([]float64, 0)
		for _, posTag := range pos {
			if doubleTag, ok := posTag.(nbt.DoubleTag); ok {
				entity.Pos = append(entity.Pos, doubleTag.Value)
			}
		}
	}
	if uuidMost, ok := compound.GetLong("UUIDMost"); ok {
		if uuidLeast, ok := compound.GetLong("UUIDLeast"); ok {
			entity.UUID = uuidFromLongs(uuidMost, uuidLeast)
		}
	}

	return entity
}

func parseTick(compound *nbt.NbtCompound) *Tick {
	tick := &Tick{}

	if i, ok := compound.GetInt("i"); ok {
		tick.I = i
	}
	if p, ok := compound.GetInt("p"); ok {
		tick.P = p
	}
	if t, ok := compound.GetInt("t"); ok {
		tick.T = t
	}
	if x, ok := compound.GetInt("x"); ok {
		tick.X = x
	}
	if y, ok := compound.GetInt("y"); ok {
		tick.Y = y
	}
	if z, ok := compound.GetInt("z"); ok {
		tick.Z = z
	}
	return tick
}

func uuidFromLongs(most, least int64) uuid.UUID {
	// Convert the two int64 values to 16 bytes (UUID)
	var uuidBytes [16]byte
	binary.BigEndian.PutUint64(uuidBytes[0:8], uint64(most))
	binary.BigEndian.PutUint64(uuidBytes[8:16], uint64(least))

	// Parse the UUID
	uuid, err := uuid.FromBytes(uuidBytes[:])
	if err != nil {
		return google_uuid.New()
	}
	return uuid
}
