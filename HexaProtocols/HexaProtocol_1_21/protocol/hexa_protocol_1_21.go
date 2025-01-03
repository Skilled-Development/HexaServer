package protocol

import (
	"HexaUtils/blocks"
	"HexaUtils/utils"
	_ "embed"
	"fmt"
)

var HexaProtocol_1_21_Instance *HexaProtocol_1_21

//go:embed blocks_1_21.json
var blockDataJSON []byte

type HexaProtocol_1_21 struct {
	BlocksMap blocks.BlockDataMap
}

func NewHexaProtocol_1_21() *HexaProtocol_1_21 {
	utils.PrintLog("Reading blocks_1_21.json File")
	// Convierte la data de json a la estructura
	blockDataMap, err := blocks.UnmarshalBlockDataMap(blockDataJSON)

	if err != nil {
		fmt.Println("Error deserializing the json", err)
	}

	HexaProtocol_1_21_Instance = &HexaProtocol_1_21{blockDataMap}
	return HexaProtocol_1_21_Instance
}

func (hp *HexaProtocol_1_21) GetBlockDataMap() blocks.BlockDataMap {
	return hp.BlocksMap
}

func GetHexaProtocol_1_21() *HexaProtocol_1_21 {
	return HexaProtocol_1_21_Instance
}
