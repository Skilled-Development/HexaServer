package regionreader

import (
	"HexaUtils/nbt"

	"github.com/google/uuid"
)

// Chunk represents a Minecraft chunk.
type Chunk struct {
	XPos           int32              `json:"xPos"`
	ZPos           int32              `json:"zPos"`
	YPos           int32              `json:"yPos"`
	DataVersion    int32              `json:"dataVersion"`
	Status         string             `json:"status"`
	LastUpdate     int64              `json:"lastUpdate"`
	Sections       []*Section         `json:"sections"`
	BlockEntities  []*BlockEntity     `json:"blockEntities"`
	CarvingMasks   map[string][]byte  `json:"carvingMasks"`
	Heightmaps     map[string][]int64 `json:"heightmaps"`
	Lights         [][]int16          `json:"lights"`
	Entities       []*Entity          `json:"entities"`
	FluidTicks     []*Tick            `json:"fluidTicks"`
	BlockTicks     []*Tick            `json:"blockTicks"`
	InhabitedTime  int64              `json:"inhabitedTime"`
	PostProcessing [][]int16          `json:"postProcessing"`
	Structures     *Structures        `json:"structures"`
}

type Structures struct {
	References map[string][]int64    `json:"references"`
	Starts     map[string]*Structure `json:"starts"`
}
type Structure struct {
	BB        []int32           `json:"bb"`
	Biome     string            `json:"biome"`
	Children  []*StructurePiece `json:"children"`
	ChunkX    int32             `json:"chunkX"`
	ChunkZ    int32             `json:"chunkZ"`
	Id        string            `json:"id"`
	Valid     bool              `json:"valid"`
	Processed []ChunkPosition   `json:"processed"`
}
type ChunkPosition struct {
	X int32 `json:"x"`
	Z int32 `json:"z"`
}

type StructurePiece struct {
	BB                []int32     `json:"bb"`
	BiomeType         string      `json:"biomeType"`
	C                 *BlockState `json:"c"`
	CA                *BlockState `json:"ca"`
	CB                *BlockState `json:"cb"`
	CC                *BlockState `json:"cc"`
	CD                *BlockState `json:"cd"`
	Chest             bool        `json:"chest"`
	D                 string      `json:"d"`
	Depth             int32       `json:"depth"`
	Entrances         []*BB       `json:"entrances"`
	EntryDoor         string      `json:"entryDoor"`
	GD                int32       `json:"gd"`
	HasPlacedChest0   bool        `json:"hasPlacedChest0"`
	HasPlacedChest1   bool        `json:"hasPlacedChest1"`
	HasPlacedChest2   bool        `json:"hasPlacedChest2"`
	HasPlacedChest3   bool        `json:"hasPlacedChest3"`
	Height            int32       `json:"height"`
	HPos              int32       `json:"hPos"`
	Hps               bool        `json:"hps"`
	Hr                bool        `json:"hr"`
	Id                string      `json:"id"`
	Integrity         float32     `json:"integrity"`
	IsLarge           bool        `json:"isLarge"`
	Junctions         []*Junction `json:"junctions"`
	Left              bool        `json:"left"`
	LeftHigh          bool        `json:"leftHigh"`
	LeftLow           bool        `json:"leftLow"`
	Length            int32       `json:"length"`
	Mob               bool        `json:"mob"`
	Num               int32       `json:"num"`
	O                 int32       `json:"o"`
	PlacedHiddenChest bool        `json:"placedHiddenChest"`
	PlacedMainChest   bool        `json:"placedMainChest"`
	PlacedTrap1       bool        `json:"placedTrap1"`
	PlacedTrap2       bool        `json:"placedTrap2"`
	PosX              int32       `json:"posX"`
	PosY              int32       `json:"posY"`
	PosZ              int32       `json:"posZ"`
	Right             bool        `json:"right"`
	RightHigh         bool        `json:"rightHigh"`
	RightLow          bool        `json:"rightLow"`
	Rot               string      `json:"rot"`
	Sc                bool        `json:"sc"`
	Seed              int64       `json:"seed"`
	Source            bool        `json:"source"`
	Steps             int32       `json:"steps"`
	T                 int32       `json:"t"`
	Tall              bool        `json:"tall"`
	Template          string      `json:"template"`
	Terrace           bool        `json:"terrace"`
	Tf                bool        `json:"tf"`
	TPX               int32       `json:"tpx"`
	TPY               int32       `json:"tpy"`
	TPZ               int32       `json:"tpz"`
	Type              int32       `json:"type"`
	VCount            int32       `json:"vCount"`
	Width             int32       `json:"width"`
	Witch             bool        `json:"witch"`
	Zombie            bool        `json:"zombie"`
}

type Junction struct {
	SourceX       int32  `json:"sourceX"`
	SourceGroundY int32  `json:"sourceGroundY"`
	SourceZ       int32  `json:"sourceZ"`
	DeltaY        int32  `json:"deltaY"`
	DestProj      string `json:"destProj"`
}
type BB struct {
	MinX int32 `json:"minX"`
	MinY int32 `json:"minY"`
	MinZ int32 `json:"minZ"`
	MaxX int32 `json:"maxX"`
	MaxY int32 `json:"maxY"`
	MaxZ int32 `json:"maxZ"`
}

// BlockEntity represents a block entity in a chunk.
type BlockEntity struct {
	Id      string           `json:"id"`
	X       int32            `json:"x"`
	Y       int32            `json:"y"`
	Z       int32            `json:"z"`
	NbtData *nbt.NbtCompound `json:"nbtData"`
}

type Entity struct {
	UUID    uuid.UUID        `json:"uuid"`
	Id      string           `json:"id"`
	Pos     []float64        `json:"pos"`
	NbtData *nbt.NbtCompound `json:"nbtData"`
}

// Tick represents a scheduled block or fluid tick.
type Tick struct {
	I int32 `json:"i"`
	P int32 `json:"p"`
	T int32 `json:"t"`
	X int32 `json:"x"`
	Y int32 `json:"y"`
	Z int32 `json:"z"`
}
