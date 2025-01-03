package registries

import "HexaUtils/nbt"

type ArmorTrimPatternRegistry struct {
	Name                 string
	ArmorTrimeRegistries []ArmorTrimPatternRegistryEntry
}

func (registry *ArmorTrimPatternRegistry) GetName() string {
	return registry.Name
}

func NewArmorTrimPatternRegistry() *ArmorTrimPatternRegistry {
	return &ArmorTrimPatternRegistry{
		Name: "trim_pattern",
		ArmorTrimeRegistries: []ArmorTrimPatternRegistryEntry{
			Bolt,
			Coast,
			Dune,
			Eye,
			FlowTrim,
			Host,
			Raiser,
			Rib,
			Sentry,
			Shaper,
			Silence,
			Snout,
			Spire,
			Tide,
			Vex,
			Ward,
			Wayfinder,
			Wild,
		},
	}
}

func (d ArmorTrimPatternRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range d.ArmorTrimeRegistries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

type ArmorTrimPatternRegistryEntry struct {
	BeautifiedName       string
	Name                 string
	AssetId              string
	TemplateItem         string
	DescriptionTranslate string
	Decal                bool
}

func (entry *ArmorTrimPatternRegistryEntry) GetBeautifiedName() string {
	return entry.BeautifiedName
}

func (entry *ArmorTrimPatternRegistryEntry) SetBeautifiedName(beautifiedName string) {
	entry.BeautifiedName = beautifiedName
}

func (entry *ArmorTrimPatternRegistryEntry) GetName() string {
	return entry.Name
}

func (entry *ArmorTrimPatternRegistryEntry) SetName(name string) {
	entry.Name = name
}

func (entry *ArmorTrimPatternRegistryEntry) GetAssetId() string {
	return entry.AssetId
}

func (entry *ArmorTrimPatternRegistryEntry) SetAssetId(assetId string) {
	entry.AssetId = assetId
}

func (entry *ArmorTrimPatternRegistryEntry) GetTemplateItem() string {
	return entry.TemplateItem
}

func (entry *ArmorTrimPatternRegistryEntry) SetTemplateItem(templateItem string) {
	entry.TemplateItem = templateItem
}

func (entry *ArmorTrimPatternRegistryEntry) GetDescriptionTranslate() string {
	return entry.DescriptionTranslate
}

func (entry *ArmorTrimPatternRegistryEntry) SetDescriptionTranslate(descriptionTranslate string) {
	entry.DescriptionTranslate = descriptionTranslate
}

func (entry *ArmorTrimPatternRegistryEntry) IsDecal() bool {
	return entry.Decal
}

func (entry *ArmorTrimPatternRegistryEntry) SetDecal(decal bool) {
	entry.Decal = decal
}

func (entry *ArmorTrimPatternRegistryEntry) GetAsNBT() nbt.Nbt {
	nbtData := nbt.NewNbt(
		entry.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{
				"asset_id":      entry.GetAssetId(),
				"template_item": entry.GetTemplateItem(),
				"decal":         entry.IsDecal(),
			}
			if len(entry.GetDescriptionTranslate()) > 0 {
				data["description"] = map[string]interface{}{
					"translate": entry.GetDescriptionTranslate(),
				}
			}
			return data
		}()),
	)
	return *nbtData
}

var (
	Bolt = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Bolt",
		Name:                 "bolt",
		AssetId:              "minecraft:bolt",
		Decal:                false,
		TemplateItem:         "minecraft:bolt_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.bolt",
	}

	Coast = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Coast",
		Name:                 "coast",
		AssetId:              "minecraft:coast",
		Decal:                false,
		TemplateItem:         "minecraft:coast_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.coast",
	}

	Dune = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Dune",
		Name:                 "dune",
		AssetId:              "minecraft:dune",
		Decal:                false,
		TemplateItem:         "minecraft:dune_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.dune",
	}

	Eye = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Eye",
		Name:                 "eye",
		AssetId:              "minecraft:eye",
		Decal:                false,
		TemplateItem:         "minecraft:eye_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.eye",
	}

	FlowTrim = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Flow",
		Name:                 "flow",
		AssetId:              "minecraft:flow",
		Decal:                false,
		TemplateItem:         "minecraft:flow_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.flow",
	}

	Host = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Host",
		Name:                 "host",
		AssetId:              "minecraft:host",
		Decal:                false,
		TemplateItem:         "minecraft:host_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.host",
	}

	Raiser = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Raiser",
		Name:                 "raiser",
		AssetId:              "minecraft:raiser",
		Decal:                false,
		TemplateItem:         "minecraft:raiser_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.raiser",
	}

	Rib = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Rib",
		Name:                 "rib",
		AssetId:              "minecraft:rib",
		Decal:                false,
		TemplateItem:         "minecraft:rib_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.rib",
	}

	Sentry = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Sentry",
		Name:                 "sentry",
		AssetId:              "minecraft:sentry",
		Decal:                false,
		TemplateItem:         "minecraft:sentry_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.sentry",
	}

	Shaper = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Shaper",
		Name:                 "shaper",
		AssetId:              "minecraft:shaper",
		Decal:                false,
		TemplateItem:         "minecraft:shaper_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.shaper",
	}

	Silence = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Silence",
		Name:                 "silence",
		AssetId:              "minecraft:silence",
		Decal:                false,
		TemplateItem:         "minecraft:silence_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.silence",
	}

	Snout = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Snout",
		Name:                 "snout",
		AssetId:              "minecraft:snout",
		Decal:                false,
		TemplateItem:         "minecraft:snout_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.snout",
	}

	Spire = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Spire",
		Name:                 "spire",
		AssetId:              "minecraft:spire",
		Decal:                false,
		TemplateItem:         "minecraft:spire_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.spire",
	}

	Tide = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Tide",
		Name:                 "tide",
		AssetId:              "minecraft:tide",
		Decal:                false,
		TemplateItem:         "minecraft:tide_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.tide",
	}

	Vex = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Vex",
		Name:                 "vex",
		AssetId:              "minecraft:vex",
		Decal:                false,
		TemplateItem:         "minecraft:vex_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.vex",
	}

	Ward = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Ward",
		Name:                 "ward",
		AssetId:              "minecraft:ward",
		Decal:                false,
		TemplateItem:         "minecraft:ward_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.ward",
	}

	Wayfinder = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Wayfinder",
		Name:                 "wayfinder",
		AssetId:              "minecraft:wayfinder",
		Decal:                false,
		TemplateItem:         "minecraft:wayfinder_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.wayfinder",
	}

	Wild = ArmorTrimPatternRegistryEntry{
		BeautifiedName:       "Wild",
		Name:                 "wild",
		AssetId:              "minecraft:wild",
		Decal:                false,
		TemplateItem:         "minecraft:wild_armor_trim_smithing_template",
		DescriptionTranslate: "trim_pattern.minecraft.wild",
	}
)
