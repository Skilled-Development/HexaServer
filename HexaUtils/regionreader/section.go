package regionreader

// Section represents a section of a Minecraft chunk.
type Section struct {
	Y           byte         `json:"y"`
	BlockStates *BlockStates `json:"blockStates"`
	Biomes      *Biomes      `json:"biomes"`
	BlockLight  []byte       `json:"blockLight"`
	SkyLight    []byte       `json:"skyLight"`
}
type Biomes struct {
	Palette []string `json:"palette"`
	Data    []int64  `json:"data"`
}

type BlockStates struct {
	Palette []*BlockState `json:"palette"`
	Data    []int64       `json:"data"`
}

type BlockState struct {
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}
