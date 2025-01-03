package registries

import "HexaUtils/nbt"

type WolfVariantRegistry struct {
	Name                  string
	WolfVariantRegistries []WolfVariantRegistryEntry
}

func (registry *WolfVariantRegistry) GetName() string {
	return registry.Name
}

func NewWolfVariantRegistry() *WolfVariantRegistry {
	return &WolfVariantRegistry{
		Name: "wolf_variant",
		WolfVariantRegistries: []WolfVariantRegistryEntry{
			Ashen,
			Black,
			Chestnut,
			Pale,
			Rusty,
			Snowy,
			Spotted,
			Striped,
			Woods,
		},
	}
}

func (d WolfVariantRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range d.WolfVariantRegistries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

type WolfVariantRegistryEntry struct {
	BeautifiedName string
	Name           string
	WildTexture    string
	TameTexture    string
	AngryTexture   string
	Biomes         []string
}

func (entry *WolfVariantRegistryEntry) GetBeautifiedName() string {
	return entry.BeautifiedName
}

func (entry *WolfVariantRegistryEntry) SetBeautifiedName(beautifiedName string) {
	entry.BeautifiedName = beautifiedName
}

func (entry *WolfVariantRegistryEntry) GetName() string {
	return entry.Name
}

func (entry *WolfVariantRegistryEntry) SetName(name string) {
	entry.Name = name
}

func (entry *WolfVariantRegistryEntry) GetWildTexture() string {
	return entry.WildTexture
}

func (entry *WolfVariantRegistryEntry) SetWildTexture(wildTexture string) {
	entry.WildTexture = wildTexture
}

func (entry *WolfVariantRegistryEntry) GetTameTexture() string {
	return entry.TameTexture
}

func (entry *WolfVariantRegistryEntry) SetTameTexture(tameTexture string) {
	entry.TameTexture = tameTexture
}

func (entry *WolfVariantRegistryEntry) GetAngryTexture() string {
	return entry.AngryTexture
}

func (entry *WolfVariantRegistryEntry) SetAngryTexture(angryTexture string) {
	entry.AngryTexture = angryTexture
}

func (entry *WolfVariantRegistryEntry) GetBiomes() []string {
	return entry.Biomes
}

func (entry *WolfVariantRegistryEntry) SetBiomes(biomes []string) {
	entry.Biomes = biomes
}

func (entry *WolfVariantRegistryEntry) GetAsNBT() nbt.Nbt {
	nbtData := nbt.NewNbt(
		entry.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{
				"wild_texture":  entry.GetWildTexture(),
				"tame_texture":  entry.GetTameTexture(),
				"angry_texture": entry.GetAngryTexture(),
				"biomes":        entry.GetBiomes(),
			}
			return data
		}()),
	)
	return *nbtData
}

var (
	Ashen = WolfVariantRegistryEntry{
		BeautifiedName: "Ashen",
		Name:           "ashen",
		AngryTexture:   "minecraft:entity/wolf/wolf_ashen_angry",
		Biomes:         []string{"minecraft:snowy_taiga"},
		TameTexture:    "minecraft:entity/wolf/wolf_ashen_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_ashen",
	}

	Black = WolfVariantRegistryEntry{
		BeautifiedName: "Black",
		Name:           "black",
		AngryTexture:   "minecraft:entity/wolf/wolf_black_angry",
		Biomes:         []string{"minecraft:old_growth_pine_taiga"},
		TameTexture:    "minecraft:entity/wolf/wolf_black_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_black",
	}

	Chestnut = WolfVariantRegistryEntry{
		BeautifiedName: "Chestnut",
		Name:           "chestnut",
		AngryTexture:   "minecraft:entity/wolf/wolf_chestnut_angry",
		Biomes:         []string{"minecraft:old_growth_spruce_taiga"},
		TameTexture:    "minecraft:entity/wolf/wolf_chestnut_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_chestnut",
	}

	Pale = WolfVariantRegistryEntry{
		BeautifiedName: "Pale",
		Name:           "pale",
		AngryTexture:   "minecraft:entity/wolf/wolf_angry",
		Biomes:         []string{"minecraft:taiga"},
		TameTexture:    "minecraft:entity/wolf/wolf_tame",
		WildTexture:    "minecraft:entity/wolf/wolf",
	}

	Rusty = WolfVariantRegistryEntry{
		BeautifiedName: "Rusty",
		Name:           "rusty",
		AngryTexture:   "minecraft:entity/wolf/wolf_rusty_angry",
		Biomes:         []string{"minecraft:jungle"}, // Corrected from #minecraft:is_jungle
		TameTexture:    "minecraft:entity/wolf/wolf_rusty_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_rusty",
	}

	Snowy = WolfVariantRegistryEntry{
		BeautifiedName: "Snowy",
		Name:           "snowy",
		AngryTexture:   "minecraft:entity/wolf/wolf_snowy_angry",
		Biomes:         []string{"minecraft:grove"},
		TameTexture:    "minecraft:entity/wolf/wolf_snowy_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_snowy",
	}

	Spotted = WolfVariantRegistryEntry{
		BeautifiedName: "Spotted",
		Name:           "spotted",
		AngryTexture:   "minecraft:entity/wolf/wolf_spotted_angry",
		Biomes:         []string{"minecraft:savanna"},
		TameTexture:    "minecraft:entity/wolf/wolf_spotted_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_spotted",
	}

	Striped = WolfVariantRegistryEntry{
		BeautifiedName: "Striped",
		Name:           "striped",
		AngryTexture:   "minecraft:entity/wolf/wolf_striped_angry",
		Biomes:         []string{"minecraft:badlands"},
		TameTexture:    "minecraft:entity/wolf/wolf_striped_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_striped",
	}

	Woods = WolfVariantRegistryEntry{
		BeautifiedName: "Woods",
		Name:           "woods",
		AngryTexture:   "minecraft:entity/wolf/wolf_woods_angry",
		Biomes:         []string{"minecraft:forest"},
		TameTexture:    "minecraft:entity/wolf/wolf_woods_tame",
		WildTexture:    "minecraft:entity/wolf/wolf_woods",
	}
)
