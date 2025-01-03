package registries

import "HexaUtils/nbt"

type PaitingVariantRegistry struct {
	Name                     string
	PaitingVariantRegistries []PaitingVariantRegistryEntry
}

func (registry *PaitingVariantRegistry) GetName() string {
	return registry.Name
}

func NewPaitingVariantRegistry() *PaitingVariantRegistry {
	return &PaitingVariantRegistry{
		Name: "painting_variant",
		PaitingVariantRegistries: []PaitingVariantRegistryEntry{
			Alban,
			Aztec,
			Aztec2,
			Backyard,
			Baroque,
			Bomb,
			Bouquet,
			BurningSkull,
			Bust,
			Cavebird,
			Changing,
			Cotan,
			Courbet,
			Creebet,
			DonkeyKong,
			Earth,
			Endboss,
			Fern,
			Fighters,
			Finding,
			Fire,
			Graham,
			Humble,
			Kebab,
			Lowmist,
			Match,
			Meditative,
			Orb,
			Owlemons,
			Passage,
			Pigscene,
			Plant,
			Pointer,
			Pond,
			Pool,
			PrairieRide,
			Sea,
			Skeleton,
			SkullAndRoses,
			Stage,
			Sunflowers,
			Sunset,
			Tides,
			Unpacked,
			Void,
			Wanderer,
			Wasteland,
			Water,
			Wind,
			WitherPaiting,
		},
	}
}

func (d PaitingVariantRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range d.PaitingVariantRegistries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

type PaitingVariantRegistryEntry struct {
	BeautifiedName string
	Name           string
	AssetId        string
	Height         int
	Width          int
}

func (entry *PaitingVariantRegistryEntry) GetBeautifiedName() string {
	return entry.BeautifiedName
}

func (entry *PaitingVariantRegistryEntry) GetName() string {
	return entry.Name
}

func (entry *PaitingVariantRegistryEntry) GetAssetId() string {
	return entry.AssetId
}

func (entry *PaitingVariantRegistryEntry) GetWidth() int {
	return entry.Width
}

func (entry *PaitingVariantRegistryEntry) GetHeight() int {
	return entry.Height
}

func (entry *PaitingVariantRegistryEntry) GetAsNBT() nbt.Nbt {
	nbtData := nbt.NewNbt(
		entry.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{
				"asset_id": entry.GetAssetId(),
				"height":   entry.GetHeight(),
				"width":    entry.GetWidth(),
			}
			return data
		}()),
	)
	return *nbtData
}

var (
	Alban = PaitingVariantRegistryEntry{
		BeautifiedName: "Alban",
		Name:           "alban",
		AssetId:        "minecraft:alban",
		Height:         1,
		Width:          1,
	}

	Aztec = PaitingVariantRegistryEntry{
		BeautifiedName: "Aztec",
		Name:           "aztec",
		AssetId:        "minecraft:aztec",
		Height:         1,
		Width:          1,
	}

	Aztec2 = PaitingVariantRegistryEntry{
		BeautifiedName: "Aztec2",
		Name:           "aztec2",
		AssetId:        "minecraft:aztec2",
		Height:         1,
		Width:          1,
	}

	Backyard = PaitingVariantRegistryEntry{
		BeautifiedName: "Backyard",
		Name:           "backyard",
		AssetId:        "minecraft:backyard",
		Height:         4,
		Width:          3,
	}

	Baroque = PaitingVariantRegistryEntry{
		BeautifiedName: "Baroque",
		Name:           "baroque",
		AssetId:        "minecraft:baroque",
		Height:         2,
		Width:          2,
	}

	Bomb = PaitingVariantRegistryEntry{
		BeautifiedName: "Bomb",
		Name:           "bomb",
		AssetId:        "minecraft:bomb",
		Height:         1,
		Width:          1,
	}

	Bouquet = PaitingVariantRegistryEntry{
		BeautifiedName: "Bouquet",
		Name:           "bouquet",
		AssetId:        "minecraft:bouquet",
		Height:         3,
		Width:          3,
	}

	BurningSkull = PaitingVariantRegistryEntry{
		BeautifiedName: "BurningSkull",
		Name:           "burning_skull",
		AssetId:        "minecraft:burning_skull",
		Height:         4,
		Width:          4,
	}

	Bust = PaitingVariantRegistryEntry{
		BeautifiedName: "Bust",
		Name:           "bust",
		AssetId:        "minecraft:bust",
		Height:         2,
		Width:          2,
	}

	Cavebird = PaitingVariantRegistryEntry{
		BeautifiedName: "Cavebird",
		Name:           "cavebird",
		AssetId:        "minecraft:cavebird",
		Height:         3,
		Width:          3,
	}

	Changing = PaitingVariantRegistryEntry{
		BeautifiedName: "Changing",
		Name:           "changing",
		AssetId:        "minecraft:changing",
		Height:         2,
		Width:          4,
	}

	Cotan = PaitingVariantRegistryEntry{
		BeautifiedName: "Cotan",
		Name:           "cotan",
		AssetId:        "minecraft:cotan",
		Height:         3,
		Width:          3,
	}

	Courbet = PaitingVariantRegistryEntry{
		BeautifiedName: "Courbet",
		Name:           "courbet",
		AssetId:        "minecraft:courbet",
		Height:         1,
		Width:          2,
	}

	Creebet = PaitingVariantRegistryEntry{
		BeautifiedName: "Creebet",
		Name:           "creebet",
		AssetId:        "minecraft:creebet",
		Height:         1,
		Width:          2,
	}

	DonkeyKong = PaitingVariantRegistryEntry{
		BeautifiedName: "DonkeyKong",
		Name:           "donkey_kong",
		AssetId:        "minecraft:donkey_kong",
		Height:         3,
		Width:          4,
	}

	Earth = PaitingVariantRegistryEntry{
		BeautifiedName: "Earth",
		Name:           "earth",
		AssetId:        "minecraft:earth",
		Height:         2,
		Width:          2,
	}

	Endboss = PaitingVariantRegistryEntry{
		BeautifiedName: "Endboss",
		Name:           "endboss",
		AssetId:        "minecraft:endboss",
		Height:         3,
		Width:          3,
	}

	Fern = PaitingVariantRegistryEntry{
		BeautifiedName: "Fern",
		Name:           "fern",
		AssetId:        "minecraft:fern",
		Height:         3,
		Width:          3,
	}

	Fighters = PaitingVariantRegistryEntry{
		BeautifiedName: "Fighters",
		Name:           "fighters",
		AssetId:        "minecraft:fighters",
		Height:         2,
		Width:          4,
	}

	Finding = PaitingVariantRegistryEntry{
		BeautifiedName: "Finding",
		Name:           "finding",
		AssetId:        "minecraft:finding",
		Height:         2,
		Width:          4,
	}

	Fire = PaitingVariantRegistryEntry{
		BeautifiedName: "Fire",
		Name:           "fire",
		AssetId:        "minecraft:fire",
		Height:         2,
		Width:          2,
	}

	Graham = PaitingVariantRegistryEntry{
		BeautifiedName: "Graham",
		Name:           "graham",
		AssetId:        "minecraft:graham",
		Height:         2,
		Width:          1,
	}

	Humble = PaitingVariantRegistryEntry{
		BeautifiedName: "Humble",
		Name:           "humble",
		AssetId:        "minecraft:humble",
		Height:         2,
		Width:          2,
	}

	Kebab = PaitingVariantRegistryEntry{
		BeautifiedName: "Kebab",
		Name:           "kebab",
		AssetId:        "minecraft:kebab",
		Height:         1,
		Width:          1,
	}

	Lowmist = PaitingVariantRegistryEntry{
		BeautifiedName: "Lowmist",
		Name:           "lowmist",
		AssetId:        "minecraft:lowmist",
		Height:         2,
		Width:          4,
	}

	Match = PaitingVariantRegistryEntry{
		BeautifiedName: "Match",
		Name:           "match",
		AssetId:        "minecraft:match",
		Height:         2,
		Width:          2,
	}

	Meditative = PaitingVariantRegistryEntry{
		BeautifiedName: "Meditative",
		Name:           "meditative",
		AssetId:        "minecraft:meditative",
		Height:         1,
		Width:          1,
	}

	Orb = PaitingVariantRegistryEntry{
		BeautifiedName: "Orb",
		Name:           "orb",
		AssetId:        "minecraft:orb",
		Height:         4,
		Width:          4,
	}

	Owlemons = PaitingVariantRegistryEntry{
		BeautifiedName: "Owlemons",
		Name:           "owlemons",
		AssetId:        "minecraft:owlemons",
		Height:         3,
		Width:          3,
	}

	Passage = PaitingVariantRegistryEntry{
		BeautifiedName: "Passage",
		Name:           "passage",
		AssetId:        "minecraft:passage",
		Height:         2,
		Width:          4,
	}

	Pigscene = PaitingVariantRegistryEntry{
		BeautifiedName: "Pigscene",
		Name:           "pigscene",
		AssetId:        "minecraft:pigscene",
		Height:         4,
		Width:          4,
	}

	Plant = PaitingVariantRegistryEntry{
		BeautifiedName: "Plant",
		Name:           "plant",
		AssetId:        "minecraft:plant",
		Height:         1,
		Width:          1,
	}

	Pointer = PaitingVariantRegistryEntry{
		BeautifiedName: "Pointer",
		Name:           "pointer",
		AssetId:        "minecraft:pointer",
		Height:         4,
		Width:          4,
	}

	Pond = PaitingVariantRegistryEntry{
		BeautifiedName: "Pond",
		Name:           "pond",
		AssetId:        "minecraft:pond",
		Height:         4,
		Width:          3,
	}

	Pool = PaitingVariantRegistryEntry{
		BeautifiedName: "Pool",
		Name:           "pool",
		AssetId:        "minecraft:pool",
		Height:         1,
		Width:          2,
	}

	PrairieRide = PaitingVariantRegistryEntry{
		BeautifiedName: "PrairieRide",
		Name:           "prairie_ride",
		AssetId:        "minecraft:prairie_ride",
		Height:         2,
		Width:          1,
	}

	Sea = PaitingVariantRegistryEntry{
		BeautifiedName: "Sea",
		Name:           "sea",
		AssetId:        "minecraft:sea",
		Height:         1,
		Width:          2,
	}

	Skeleton = PaitingVariantRegistryEntry{
		BeautifiedName: "Skeleton",
		Name:           "skeleton",
		AssetId:        "minecraft:skeleton",
		Height:         3,
		Width:          4,
	}

	SkullAndRoses = PaitingVariantRegistryEntry{
		BeautifiedName: "SkullAndRoses",
		Name:           "skull_and_roses",
		AssetId:        "minecraft:skull_and_roses",
		Height:         2,
		Width:          2,
	}

	Stage = PaitingVariantRegistryEntry{
		BeautifiedName: "Stage",
		Name:           "stage",
		AssetId:        "minecraft:stage",
		Height:         2,
		Width:          2,
	}

	Sunflowers = PaitingVariantRegistryEntry{
		BeautifiedName: "Sunflowers",
		Name:           "sunflowers",
		AssetId:        "minecraft:sunflowers",
		Height:         3,
		Width:          3,
	}

	Sunset = PaitingVariantRegistryEntry{
		BeautifiedName: "Sunset",
		Name:           "sunset",
		AssetId:        "minecraft:sunset",
		Height:         1,
		Width:          2,
	}

	Tides = PaitingVariantRegistryEntry{
		BeautifiedName: "Tides",
		Name:           "tides",
		AssetId:        "minecraft:tides",
		Height:         3,
		Width:          3,
	}

	Unpacked = PaitingVariantRegistryEntry{
		BeautifiedName: "Unpacked",
		Name:           "unpacked",
		AssetId:        "minecraft:unpacked",
		Height:         4,
		Width:          4,
	}

	Void = PaitingVariantRegistryEntry{
		BeautifiedName: "Void",
		Name:           "void",
		AssetId:        "minecraft:void",
		Height:         2,
		Width:          2,
	}

	Wanderer = PaitingVariantRegistryEntry{
		BeautifiedName: "Wanderer",
		Name:           "wanderer",
		AssetId:        "minecraft:wanderer",
		Height:         2,
		Width:          1,
	}

	Wasteland = PaitingVariantRegistryEntry{
		BeautifiedName: "Wasteland",
		Name:           "wasteland",
		AssetId:        "minecraft:wasteland",
		Height:         1,
		Width:          1,
	}

	Water = PaitingVariantRegistryEntry{
		BeautifiedName: "Water",
		Name:           "water",
		AssetId:        "minecraft:water",
		Height:         2,
		Width:          2,
	}

	Wind = PaitingVariantRegistryEntry{
		BeautifiedName: "Wind",
		Name:           "wind",
		AssetId:        "minecraft:wind",
		Height:         2,
		Width:          2,
	}

	WitherPaiting = PaitingVariantRegistryEntry{
		BeautifiedName: "Wither",
		Name:           "wither",
		AssetId:        "minecraft:wither",
		Height:         2,
		Width:          2,
	}
)
