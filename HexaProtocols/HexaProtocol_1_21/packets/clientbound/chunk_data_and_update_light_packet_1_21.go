package clientbound

import (
	"HexaProtocol_1_21/protocol"
	"HexaUtils/entities/player"
	"HexaUtils/nbt"
	"HexaUtils/packets"
	"HexaUtils/regionreader"
	"HexaUtils/registries"
	debugger "HexaUtils/utils"
	"bytes"
	"encoding/binary"
	"math"
	"strings"
)

type ChunkDataAndUpdateLightPacket_1_21 struct {
	PacketID              int
	ServerBoundPacket     bool
	ProtocolVersion       int
	State                 player.ClientState
	ChunkX                int32
	ChunkZ                int32
	Heightmaps            nbt.Nbt
	Data                  []byte
	BlockEntitiesCount    int32
	BlockEntities         []byte // NBT
	SkyLightMask          []uint64
	BlockLightMask        []uint64
	EmptySkyLightMask     []uint64
	EmptyBlockLightMask   []uint64
	SkyLightArraysCount   int32
	SkyLightArrays        []byte // Modified to be a single byte array
	BlockLightArraysCount int32
	BlockLightArrays      []byte // Modified to be a single byte array
}

func NewChunkDataAndUpdateLightPacket_1_21_FromChunkStruct(
	chunk *regionreader.Chunk,
	playerToSend player.Player,
) *ChunkDataAndUpdateLightPacket_1_21 {
	debugger.PrintForDebug("\n\n")
	heightmapsNBT := createHeightmapsNBT(chunk.Heightmaps)

	debugger.PrintForDebug("\n\n")
	debugger.PrintForDebug("Now block entities")
	var blockEntitiesBytes []byte
	var blockEntitiesCount int32
	if len(chunk.BlockEntities) > 0 {
		blockEntitiesCount = int32(len(chunk.BlockEntities))
		debugger.PrintForDebug("Block Entities Count: %d", blockEntitiesCount)
		for _, blockEntity := range chunk.BlockEntities {
			blockEntityPacketWriter := packets.NewPacketWriter()
			packet_xz := ((blockEntity.X & 15) << 4) | (blockEntity.Z & 15)
			blockEntityPacketWriter.WriteUnsignedByte(uint8(packet_xz))
			blockEntityPacketWriter.WriteShort(int16(blockEntity.Y))
			blockEntityPacketWriter.WriteVarInt(0) //BlockEntity Type
			blockEntityNBT := createBlockEntityNBT(blockEntity)
			blockEntityPacketWriter.WriteNBT(*blockEntityNBT)
			blockEntitiesBytes = append(blockEntitiesBytes, blockEntityPacketWriter.GetPacketBuffer()...)

		}

	} else {
		blockEntitiesCount = 0
		blockEntitiesBytes = []byte{}
	}
	chunkData := serializeChunkData(chunk)

	return &ChunkDataAndUpdateLightPacket_1_21{
		PacketID:           0x27,
		ServerBoundPacket:  false,
		ProtocolVersion:    757,
		State:              player.Play,
		ChunkX:             chunk.XPos,
		ChunkZ:             chunk.ZPos,
		Heightmaps:         *heightmapsNBT,
		Data:               chunkData,
		BlockEntitiesCount: blockEntitiesCount,
		BlockEntities:      blockEntitiesBytes,
	}
}

func countSetBits(mask []uint64) int {
	count := 0
	for _, v := range mask {
		for i := 0; i < 64; i++ {
			if (v>>i)&1 == 1 {
				count++
			}
		}
	}
	return count
}

func createBlockEntityNBT(blockEntity *regionreader.BlockEntity) *nbt.Nbt {
	debugger.PrintForDebug("\n\n")
	debugger.PrintForDebug("Creating Block Entity NBT")
	debugger.PrintForDebug("Block Entity ID: %s", blockEntity.Id)
	debugger.PrintForDebug("Block Entity X: %d", blockEntity.X)
	debugger.PrintForDebug("Block Entity Y: %d", blockEntity.Y)
	debugger.PrintForDebug("Block Entity Z: %d", blockEntity.Z)
	debugger.PrintForDebug("Block Entity NBT Data: %v", blockEntity.NbtData)
	nbtData := map[string]interface{}{
		"id": blockEntity.Id,
		"x":  blockEntity.X,
		"y":  blockEntity.Y,
		"z":  blockEntity.Z,
	}

	if blockEntity.NbtData != nil {
		for key, value := range blockEntity.NbtData.ChildTags {
			nbtData[key] = value
		}
	}

	compound := nbt.NbtCompoundFromInterfaceMap(nbtData)
	return nbt.NewNbt("", compound)
}

func createHeightmapsNBT(heightmaps map[string][]int64) *nbt.Nbt {
	nbtData := map[string]interface{}{}
	debugger.PrintForDebug("Creating Heightmaps NBT")
	for key, value := range heightmaps {
		if key == "MOTION_BLOCKING" || key == "WORLD_SURFACE" {
			debugger.PrintForDebug("Heightmap Key: %s", key)
			debugger.PrintForDebug("Heightmap Value: %v", value)
			nbtData[key] = value
		}
	}

	compound := nbt.NbtCompoundFromInterfaceMap(nbtData)
	return nbt.NewNbt("", compound)
}

func serializeChunkData(chunk *regionreader.Chunk) []byte {
	var buffer bytes.Buffer

	for _, section := range chunk.Sections {
		debugger.PrintForDebug("\n\n")
		debugger.PrintForDebug("Serialize Section: %d ----------", section.Y)

		serializeSection(&buffer, section)
	}

	return buffer.Bytes()
}

func serializeSection(buffer *bytes.Buffer, section *regionreader.Section) {
	blockCount := calculateNonAirBlockCount(section.BlockStates)
	debugger.PrintForDebug("\n\n")
	debugger.PrintForDebug("Serialize Section: %d", section.Y)
	debugger.PrintForDebug("Block Count Of Air Blocks: %d", blockCount)
	binary.Write(buffer, binary.BigEndian, int16(blockCount))

	serializePalettedContainer(buffer, section.BlockStates)
	serializePalettedContainer(buffer, section.Biomes)
}

func calculateNonAirBlockCount(blockStates *regionreader.BlockStates) int {
	if blockStates == nil {
		return 0
	}
	count := 0
	for _, id := range blockStates.Data {
		//count += bits.OnesCount64(uint64(id))
		if id != 0 {
			count++
		}
	}
	return count
}

func serializePalettedContainer(buffer *bytes.Buffer, container interface{}) {
	switch v := container.(type) {
	case *regionreader.BlockStates:
		serializeBlockStates(buffer, v)
	case *regionreader.Biomes:
		serializeBiomes(buffer, v)
	}
}
func serializeBlockStates(buffer *bytes.Buffer, blockStates *regionreader.BlockStates) {
	if blockStates == nil {
		buffer.WriteByte(0)
		binary.Write(buffer, binary.BigEndian, int32(0))
		return
	}
	debugger.PrintForDebug("\n\n")
	paletteLength := len(blockStates.Palette)
	debugger.PrintForDebug("Palette Length: %d, For BlockStates", paletteLength)
	var bitsPerEntry byte

	if paletteLength <= 1 {
		bitsPerEntry = 0
	} else {
		bitsPerEntry = calculateBitsPerEntryBlock(paletteLength)
	}

	debugger.PrintForDebug("Bits Per Entry: %d", bitsPerEntry)

	buffer.WriteByte(bitsPerEntry)
	if bitsPerEntry == 0 {
		debugger.PrintForDebug("Bits Per Entry is 0")
		if len(blockStates.Palette) > 0 {
			debugger.PrintForDebug("Block Name: %s", blockStates.Palette[0].Name)
			blockStateId := getBlockStateID(blockStates.Palette[0].Name, blockStates.Palette[0].Properties)
			nbt.WriteVarInt(buffer, int32(blockStateId))
		} else {
			debugger.PrintForDebug("Block Name: minecraft:air")
			nbt.WriteVarInt(buffer, int32(0)) //Air
		}
		debugger.PrintForDebug("Data Array Length: 0")
		nbt.WriteVarInt(buffer, int32(0)) //Data array len
		return
	}

	if bitsPerEntry < 15 {
		debugger.PrintForDebug("Palette Length: %d", paletteLength)
		nbt.WriteVarInt(buffer, int32(paletteLength))
		debugger.PrintForDebug("Writing Palette")
		for _, blockState := range blockStates.Palette {
			blockStateId := getBlockStateID(blockState.Name, blockState.Properties)
			debugger.PrintForDebug("Writing BlockState ID: %d", blockStateId)
			nbt.WriteVarInt(buffer, int32(blockStateId))

		}
	}

	dataArray := createDataArray(blockStates.Data, int(bitsPerEntry), paletteLength)
	debugger.PrintForDebug("Data Array Length: %d", len(dataArray))
	nbt.WriteVarInt(buffer, int32(len(dataArray))) // Data Array Length
	for _, longVal := range dataArray {
		binary.Write(buffer, binary.BigEndian, longVal)
	}
}

func serializeBiomes(buffer *bytes.Buffer, biomes *regionreader.Biomes) {
	if biomes == nil {
		buffer.WriteByte(0)
		binary.Write(buffer, binary.BigEndian, int32(0))
		return
	}
	debugger.PrintForDebug("\n\n")
	debugger.PrintForDebug("Biomes Serialize")
	debugger.PrintForDebug("Biomes Palette: %v", biomes.Palette)
	debugger.PrintForDebug("Biomes Data: %v", biomes.Data)
	paletteLength := len(biomes.Palette)
	debugger.PrintForDebug("Palette Length: %d, For Biomes", paletteLength)

	var bitsPerEntry byte
	var maxPaletteIndex int = 0

	if paletteLength > 0 {
		for _, biome := range biomes.Palette {
			biomeID := getBiomeID(biome)
			if biomeID > maxPaletteIndex {
				maxPaletteIndex = biomeID
			}
		}
	}
	if paletteLength <= 1 {
		bitsPerEntry = 0
		debugger.PrintForDebug("Bits Per Entry: 0")
	} else {
		bitsPerEntry = calculateBitsPerEntryBiome(paletteLength)
		debugger.PrintForDebug("Bits Per Entry: %d", bitsPerEntry)
	}

	buffer.WriteByte(bitsPerEntry)

	if bitsPerEntry == 0 {
		debugger.PrintForDebug("Bits Per Entry is 0")
		if len(biomes.Palette) > 0 {
			biomeID := getBiomeID(biomes.Palette[0])
			nbt.WriteVarInt(buffer, int32(biomeID))
			debugger.PrintForDebug("Biome ID: %d", biomeID)
		} else {
			debugger.PrintForDebug("Default Biome: 0")
			nbt.WriteVarInt(buffer, int32(0)) //default biome
		}
		nbt.WriteVarInt(buffer, int32(0)) // Data Array Length
		debugger.PrintForDebug("Data Array Length: 0")
		return
	}

	if bitsPerEntry < 6 {
		debugger.PrintForDebug("Palette Length: %d", paletteLength)
		nbt.WriteVarInt(buffer, int32(paletteLength))
		for _, biome := range biomes.Palette {
			biomeID := getBiomeID(biome)
			debugger.PrintForDebug("Biome ID: %d", biomeID)
			nbt.WriteVarInt(buffer, int32(biomeID))
		}
	}

	dataArray := createDataArray(biomes.Data, int(bitsPerEntry), paletteLength)
	nbt.WriteVarInt(buffer, int32(len(dataArray))) // Data Array Length
	for _, longVal := range dataArray {
		binary.Write(buffer, binary.BigEndian, longVal)
	}
}

func calculateBitsPerEntryBlock(paletteLength int) byte {
	if paletteLength <= 1 {
		return 0
	}
	bits := byte(math.Ceil(math.Log2(float64(paletteLength))))
	if bits < 4 {
		return 4
	} else if bits < 15 {
		return bits
	} else {
		return 15 // Direct palette
	}

}
func calculateBitsPerEntryBiome(paletteLength int) byte {
	if paletteLength <= 1 {
		return 0
	}
	bits := byte(math.Ceil(math.Log2(float64(paletteLength))))
	if bits < 1 {
		return 1
	} else if bits < 4 {
		return bits
	} else {
		return 6 // Direct Palette
	}

}
func createDataArray(data []int64, bitsPerEntry int, paletteLength int) []int64 {
	if len(data) == 0 {
		return []int64{}
	}
	if bitsPerEntry == 0 {
		return []int64{}
	}
	if bitsPerEntry >= 15 {
		return data
	}

	longArrayLength := (4096*bitsPerEntry + 63) / 64
	result := make([]int64, longArrayLength)

	individualValueMask := int64((1 << bitsPerEntry) - 1)

	//We must iterate over the block numbers
	blockIndex := 0
	for index := 0; index < len(data); index++ {
		val := data[index]
		for i := 0; blockIndex < 4096 && i < 64; i += bitsPerEntry {

			startLong := (blockIndex * bitsPerEntry) / 64
			startOffset := (blockIndex * bitsPerEntry) % 64
			endLong := ((blockIndex+1)*bitsPerEntry - 1) / 64

			value := (val >> i) & individualValueMask

			result[startLong] |= (value << startOffset)
			if startLong != endLong {
				result[endLong] = (value >> (64 - startOffset))
			}

			blockIndex++
		}
	}
	return result
}
func getBlockStateID(name string, properties map[string]string) int {
	return protocol.GetHexaProtocol_1_21().GetBlockDataMap().GetBlockID(name, properties)
}

func getBiomeID(biome string) int {
	for index, b := range registries.BiomeRegistryInstance.BiomeRegistryEntries {
		if b.GetName() == strings.Split(biome, ":")[1] {
			return index
		}
	}
	return 0
}

func setBits(bitSet []uint64, startIndex int, numBits int) []uint64 {
	if len(bitSet) == 0 {
		bitSet = append(bitSet, 0)
	}

	for i := 0; i < numBits; i++ {
		bitIndex := startIndex + i
		longIndex := bitIndex / 64

		if longIndex >= len(bitSet) {
			newBitSet := make([]uint64, longIndex+1)
			copy(newBitSet, bitSet)
			bitSet = newBitSet
		}

		bitSet[longIndex] |= 1 << (bitIndex % 64)
	}
	return bitSet
}

func createSkyLightArray() []byte {
	arrayLength := 2048
	skyLightArray := make([]byte, arrayLength)
	var lightValue uint8 = 0b00001111 // Maximum light: 15

	for y := 0; y < 16; y++ {
		for z := 0; z < 16; z++ {
			for x := 0; x < 16; x += 2 {
				// Calculate the array index using the formula
				index := ((y << 8) | (z << 4) | x) / 2

				// If x is less than 16, it is within bounds.
				if x < 16 {
					// Pack the 4-bit light value into the lower half of the byte.
					skyLightArray[index] |= lightValue // Use a bitwise OR to avoid overwriting
				}

				if x+1 < 16 {
					// Pack the next 4-bit light value into the upper half of the byte.
					skyLightArray[index] |= (lightValue << 4)
				}
			}
		}
	}
	return skyLightArray
}
func createBlockLightArray() []byte {
	// Since there is no block light, lets keep the array empty
	return []byte{}
}

func (p *ChunkDataAndUpdateLightPacket_1_21) GetPacket() *packets.Packet {
	dataSize := len(p.Data)
	packet := packets.NewPacketWriter()
	packet.WriteVarInt(int32(p.PacketID))
	packet.WriteInt(p.ChunkX)
	packet.WriteInt(p.ChunkZ)
	packet.WriteNBT(p.Heightmaps)
	packet.WriteVarInt(int32(dataSize))
	packet.AppendByteArray(p.Data)
	packet.WriteVarInt(p.BlockEntitiesCount)
	packet.AppendByteArray(p.BlockEntities)

	numSections := 24

	// Sky Light Mask - All sections have sky light data.
	skyLightMask := []uint64{}
	skyLightMask = setBits(skyLightMask, 0, numSections)

	// Empty Sky Light Mask - No sections are empty.
	emptySkyLightMask := []uint64{} // remains empty

	// Block Light Mask - No sections have block light data, so all bits will be zero
	blockLightMask := []uint64{}

	// Empty Block Light Mask - All sections have empty block light, so, all bits will be 1
	emptyBlockLightMask := []uint64{}
	emptyBlockLightMask = setBits(emptyBlockLightMask, 0, numSections)

	// Create Sky Light Arrays - One for each section
	skyLightArrays := make([][]byte, numSections)
	for i := 0; i < numSections; i++ {
		skyLightArrays[i] = createSkyLightArray()
	}

	// Create Block Light Arrays, empty as there is no block light
	blockLightArrays := [][]byte{}

	packet.WriteBitSet(skyLightMask)
	packet.WriteBitSet(emptySkyLightMask)

	packet.WriteBitSet(blockLightMask)

	packet.WriteBitSet(emptyBlockLightMask)
	packet.WriteVarInt(int32(len(skyLightArrays)))
	for _, skyLightArray := range skyLightArrays {
		packet.WriteByteArray(skyLightArray)
	}

	packet.WriteVarInt(int32(len(blockLightArrays)))
	for _, blockLightArray := range blockLightArrays {
		packet.WriteByteArray(blockLightArray)
	}

	real_packet := packets.NewPacket(p.PacketID,
		p.ProtocolVersion,
		"ChunkDataAndUpdateLightPacket",
		packet.GetPacketBuffer(),
		p.ServerBoundPacket,
		p.State)
	return real_packet
}
