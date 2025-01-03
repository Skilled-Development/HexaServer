package registries

import "HexaUtils/nbt"

type ArmorTrimRegistry struct {
	Name                 string
	ArmorTrimeRegistries []ArmorTrimRegistryEntry
}

func (registry *ArmorTrimRegistry) GetName() string {
	return registry.Name
}

func NewArmorTrimRegistry() *ArmorTrimRegistry {
	return &ArmorTrimRegistry{
		Name: "trim_material",
		ArmorTrimeRegistries: []ArmorTrimRegistryEntry{
			Amethyst,
			Copper,
			Diamond,
			Emerald,
			Gold,
			Iron,
			Lapis,
			Netherite,
			Quartz,
			Redstone,
		},
	}
}

func (d ArmorTrimRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range d.ArmorTrimeRegistries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

type ArmorTrimRegistryEntry struct {
	BeautifiedName         string
	Name                   string
	AssetName              string
	Ingredient             string
	ItemModelIndex         float32
	DescriptionColor       string
	DescriptionTranslate   string
	OverrideArmorMaterials map[string]interface{}
}

func (entry *ArmorTrimRegistryEntry) GetBeautifiedName() string {
	return entry.BeautifiedName
}

func (entry *ArmorTrimRegistryEntry) SetBeautifiedName(beautifiedName string) {
	entry.BeautifiedName = beautifiedName
}

func (entry *ArmorTrimRegistryEntry) GetName() string {
	return entry.Name
}

func (entry *ArmorTrimRegistryEntry) SetName(name string) {
	entry.Name = name
}

func (entry *ArmorTrimRegistryEntry) GetAssetName() string {
	return entry.AssetName
}

func (entry *ArmorTrimRegistryEntry) SetAssetName(assetName string) {
	entry.AssetName = assetName
}

func (entry *ArmorTrimRegistryEntry) GetIngredient() string {
	return entry.Ingredient
}

func (entry *ArmorTrimRegistryEntry) SetIngredient(ingredient string) {
	entry.Ingredient = ingredient
}

func (entry *ArmorTrimRegistryEntry) GetItemModelIndex() float32 {
	return entry.ItemModelIndex
}

func (entry *ArmorTrimRegistryEntry) SetItemModelIndex(itemModelIndex float32) {
	entry.ItemModelIndex = itemModelIndex
}

func (entry *ArmorTrimRegistryEntry) GetDescriptionColor() string {
	return entry.DescriptionColor
}

func (entry *ArmorTrimRegistryEntry) SetDescriptionColor(descriptionColor string) {
	entry.DescriptionColor = descriptionColor
}

func (entry *ArmorTrimRegistryEntry) GetDescriptionTranslate() string {
	return entry.DescriptionTranslate
}

func (entry *ArmorTrimRegistryEntry) SetDescriptionTranslate(descriptionTranslate string) {
	entry.DescriptionTranslate = descriptionTranslate
}

func (entry *ArmorTrimRegistryEntry) GetOverrideArmorMaterials() map[string]interface{} {
	return entry.OverrideArmorMaterials
}

func (entry *ArmorTrimRegistryEntry) SetOverrideArmorMaterials(overrideArmorMaterials map[string]interface{}) {
	entry.OverrideArmorMaterials = overrideArmorMaterials
}

func (ae *ArmorTrimRegistryEntry) GetAsNBT() nbt.Nbt {

	nbtData := nbt.NewNbt(
		ae.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{
				"asset_name":       ae.GetAssetName(),
				"ingredient":       ae.GetIngredient(),
				"item_model_index": ae.GetItemModelIndex(),
			}
			if len(ae.GetDescriptionColor()) > 0 {
				data["description"] = map[string]interface{}{
					"translate": ae.GetDescriptionTranslate(),
					"color":     ae.GetDescriptionColor(),
				}
			}
			if len(ae.GetOverrideArmorMaterials()) > 0 {
				data["override_armor_materials"] = ae.GetOverrideArmorMaterials()
			}
			return data
		}()),
	)
	return *nbtData
}

/**
func (d *DamageTypeRegistryEntry) GetAsNBT() nbt.Nbt {
	nbtData := nbt.NewNbt(
		d.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{
				"message_id": d.GetMessageID(),
				"scaling":    d.GetScaling(),
				"exhaustion": d.GetExhaustion(),
			}
			if d.GetDeathMessageType() != "" {
				data["death_message_type"] = d.GetDeathMessageType()
			}
			if d.GetEffects() != "" {
				data["effects"] = d.GetEffects()
			}
			return data
		}()),
	)
	return *nbtData

}

**/

var (
	Amethyst = ArmorTrimRegistryEntry{
		BeautifiedName:         "Amethyst",
		Name:                   "amethyst",
		AssetName:              "amethyst",
		Ingredient:             "minecraft:amethyst_shard",
		ItemModelIndex:         1.0,
		DescriptionColor:       "#9A5CC6",
		DescriptionTranslate:   "trim_material.minecraft.amethyst",
		OverrideArmorMaterials: map[string]interface{}{},
	}

	Copper = ArmorTrimRegistryEntry{
		BeautifiedName:         "Copper",
		Name:                   "copper",
		AssetName:              "copper",
		Ingredient:             "minecraft:copper_ingot",
		ItemModelIndex:         0.5,
		DescriptionColor:       "#B4684D",
		DescriptionTranslate:   "trim_material.minecraft.copper",
		OverrideArmorMaterials: map[string]interface{}{},
	}

	Diamond = ArmorTrimRegistryEntry{
		BeautifiedName:         "Diamond",
		Name:                   "diamond",
		AssetName:              "diamond",
		Ingredient:             "minecraft:diamond",
		ItemModelIndex:         0.8,
		DescriptionColor:       "#6EECD2",
		DescriptionTranslate:   "trim_material.minecraft.diamond",
		OverrideArmorMaterials: map[string]interface{}{"minecraft:diamond": "diamond_darker"},
	}

	Emerald = ArmorTrimRegistryEntry{
		BeautifiedName:         "Emerald",
		Name:                   "emerald",
		AssetName:              "emerald",
		Ingredient:             "minecraft:emerald",
		ItemModelIndex:         0.7,
		DescriptionColor:       "#11A036",
		DescriptionTranslate:   "trim_material.minecraft.emerald",
		OverrideArmorMaterials: map[string]interface{}{},
	}

	Gold = ArmorTrimRegistryEntry{
		BeautifiedName:         "Gold",
		Name:                   "gold",
		AssetName:              "gold",
		Ingredient:             "minecraft:gold_ingot",
		ItemModelIndex:         0.6,
		DescriptionColor:       "#DEB12D",
		DescriptionTranslate:   "trim_material.minecraft.gold",
		OverrideArmorMaterials: map[string]interface{}{"minecraft:gold": "gold_darker"},
	}

	Iron = ArmorTrimRegistryEntry{
		BeautifiedName:         "Iron",
		Name:                   "iron",
		AssetName:              "iron",
		Ingredient:             "minecraft:iron_ingot",
		ItemModelIndex:         0.2,
		DescriptionColor:       "#ECECEC",
		DescriptionTranslate:   "trim_material.minecraft.iron",
		OverrideArmorMaterials: map[string]interface{}{"minecraft:iron": "iron_darker"},
	}

	Lapis = ArmorTrimRegistryEntry{
		BeautifiedName:         "Lapis",
		Name:                   "lapis",
		AssetName:              "lapis",
		Ingredient:             "minecraft:lapis_lazuli",
		ItemModelIndex:         0.9,
		DescriptionColor:       "#416E97",
		DescriptionTranslate:   "trim_material.minecraft.lapis",
		OverrideArmorMaterials: map[string]interface{}{},
	}

	Netherite = ArmorTrimRegistryEntry{
		BeautifiedName:         "Netherite",
		Name:                   "netherite",
		AssetName:              "netherite",
		Ingredient:             "minecraft:netherite_ingot",
		ItemModelIndex:         0.3,
		DescriptionColor:       "#625859",
		DescriptionTranslate:   "trim_material.minecraft.netherite",
		OverrideArmorMaterials: map[string]interface{}{"minecraft:netherite": "netherite_darker"},
	}

	Quartz = ArmorTrimRegistryEntry{
		BeautifiedName:         "Quartz",
		Name:                   "quartz",
		AssetName:              "quartz",
		Ingredient:             "minecraft:quartz",
		ItemModelIndex:         0.1,
		DescriptionColor:       "#E3D4C4",
		DescriptionTranslate:   "trim_material.minecraft.quartz",
		OverrideArmorMaterials: map[string]interface{}{},
	}

	Redstone = ArmorTrimRegistryEntry{
		BeautifiedName:         "Redstone",
		Name:                   "redstone",
		AssetName:              "redstone",
		Ingredient:             "minecraft:redstone",
		ItemModelIndex:         0.4,
		DescriptionColor:       "#971607",
		DescriptionTranslate:   "trim_material.minecraft.redstone",
		OverrideArmorMaterials: map[string]interface{}{},
	}
)
