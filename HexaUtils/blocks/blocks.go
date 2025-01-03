package blocks

import (
	"encoding/json"
	"runtime"
)

// PropertyValue representa un valor posible para una propiedad.
type PropertyValue string

// PropertyValues representa los valores permitidos para una propiedad.
type PropertyValues []PropertyValue

// BlockProperty representa una propiedad de un bloque.
type BlockProperty struct {
	Name   string         `json:"name"`
	Values PropertyValues `json:"values"`
}

// BlockProperties representa las propiedades con sus valores posibles.
type BlockProperties map[string]PropertyValues

// BlockState representa un estado específico de un bloque.
type BlockState struct {
	ID         int               `json:"id"`
	Properties map[string]string `json:"properties"`
	Default    bool              `json:"default,omitempty"` // opcional
}

// BlockData representa la información completa de un bloque.
type BlockData struct {
	Name               string          `json:"-"`
	DefaultStateOffset int             `json:"default_state_offset"`
	Properties         BlockProperties `json:"properties"`
	States             []BlockState    `json:"states"`

	// Mapa para búsqueda rápida por ID de estado.
	stateMap   map[int]*BlockState
	Definition BlockDefinition `json:"definition"`
}

// BlockDefinition representa la definición general del bloque.
type BlockDefinition struct {
	Type               string            `json:"type"`
	BlockSetType       string            `json:"block_set_type"`
	Properties         map[string]string `json:"properties"`                      // Puede ser un mapa vacío
	TicksToStayPressed int               `json:"ticks_to_stay_pressed,omitempty"` // opcional
}

// BlockDataMap representa el mapa principal de todos los bloques
type BlockDataMap map[string]*BlockData

// Metodo para crear el Mapa de estados.
func (bd *BlockData) IndexStates() {
	bd.stateMap = make(map[int]*BlockState)
	for i := range bd.States {
		state := &bd.States[i]
		bd.stateMap[state.ID] = state
	}
}

func (bdm BlockDataMap) GetBlockID(name string, properties map[string]string) int {
	blockData := bdm[name]
	if blockData == nil {
		return 0
	}

	if len(properties) == 0 {
		return blockData.States[0].ID
	}

	states := blockData.States
	for _, state := range states {
		match := true
		for key, value := range properties {
			if stateValue, ok := state.Properties[key]; !ok || stateValue != value {
				match = false
				break
			}
		}
		if match {
			return state.ID
		}
	}
	println("No se encontro el bloque")
	println("Nombre: ", name)
	println("Propiedades: ", properties)
	return 0
}

// Metodo para buscar por id dentro del bloque
func (bd *BlockData) GetState(id int) *BlockState {
	return bd.stateMap[id]
}

func UnmarshalBlockDataMap(data []byte) (BlockDataMap, error) {
	var blockDataMap map[string]BlockData
	if err := json.Unmarshal(data, &blockDataMap); err != nil {
		return nil, err
	}

	res := make(BlockDataMap, len(blockDataMap))
	for name, blockData := range blockDataMap {
		if len(blockData.States) == 0 {
			states, err := generateStates(&blockData)
			if err != nil {
				return nil, err
			}
			blockData.States = states
		}
		blockData.Name = name // Se guarda el nombre del bloque como clave
		blockData.IndexStates()
		res[name] = &blockData
	}

	// Liberar memoria del mapa temporal
	blockDataMap = nil

	// Forzar la recolección de basura
	runtime.GC()

	return res, nil
}

func generateStates(blockData *BlockData) ([]BlockState, error) {
	if len(blockData.Properties) == 0 {
		return []BlockState{}, nil
	}

	totalStates := 1
	for _, propValues := range blockData.Properties {
		totalStates *= len(propValues)
	}

	states := make([]BlockState, totalStates)
	propKeys := make([]string, 0, len(blockData.Properties))
	for k := range blockData.Properties {
		propKeys = append(propKeys, k)
	}

	for i := 0; i < totalStates; i++ {
		state := BlockState{
			ID:         i + blockData.DefaultStateOffset,
			Properties: make(map[string]string, len(propKeys)),
		}

		temp := i
		for _, propKey := range propKeys {
			valueIndex := temp % len(blockData.Properties[propKey])
			state.Properties[propKey] = string(blockData.Properties[propKey][valueIndex])
			temp = temp / len(blockData.Properties[propKey])
		}
		states[i] = state
	}

	// Liberar memoria de propKeys
	propKeys = nil

	return states, nil
}
