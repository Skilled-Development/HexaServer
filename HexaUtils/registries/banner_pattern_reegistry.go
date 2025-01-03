package registries

import "HexaUtils/nbt"

type BannerPatternRegistry struct {
	Name                         string
	BannerPatternRegistryEntries []BannerPatternRegistryEntry
}

func NewBannerPatternRegistry() *BannerPatternRegistry {
	return &BannerPatternRegistry{
		Name: "banner_pattern",
		BannerPatternRegistryEntries: []BannerPatternRegistryEntry{
			Base,
			Border,
			Bricks,
			Circle,
			Creeper,
			Cross,
			CurlyBorder,
			DiagonalLeft,
			DiagonalRight,
			DiagonalUpLeft,
			DiagonalUpRight,
			Flow,
			Flower,
			Globe,
			Gradient,
			GradientUp,
			Guster,
			HalfHorizontal,
			HalfHorizontalBottom,
			HalfVertical,
			HalfVerticalRight,
			Mojang,
			Piglin,
			Rhombus,
			Skull,
			SmallStripes,
			SquareBottomLeft,
			SquareBottomRight,
			SquareTopLeft,
			SquareTopRight,
			StraightCross,
			StripeBottom,
			StripeCenter,
			StripeDownleft,
			StripeDownright,
			StripeLeft,
			StripeMiddle,
			StripeRight,
			StripeTop,
			TrianglesBottom,
			TrianglesTop,
			TriangleBottom,
			TriangleTop,
		},
	}
}

func (b BannerPatternRegistry) GetEntries() []BannerPatternRegistryEntry {
	return b.BannerPatternRegistryEntries
}

func (b BannerPatternRegistry) GetEntryByName(name string) *BannerPatternRegistryEntry {
	for _, entry := range b.BannerPatternRegistryEntries {
		if entry.GetName() == name {
			return &entry
		}
	}
	return nil
}

func (b BannerPatternRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range b.BannerPatternRegistryEntries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

func (b BannerPatternRegistry) GetName() string {
	return b.Name
}

type BannerPatternRegistryEntry struct {
	BeautifiedName string
	Name           string
	AssetId        string
	TranslationKey string
}

func (b *BannerPatternRegistryEntry) GetAssetId() string {
	return b.AssetId
}

func (b *BannerPatternRegistryEntry) SetAssetId(AssetId string) {
	b.AssetId = AssetId
}

func (b *BannerPatternRegistryEntry) GetTranslationKey() string {
	return b.TranslationKey
}

func (b *BannerPatternRegistryEntry) SetTranslationKey(TranslationKey string) {
	b.TranslationKey = TranslationKey
}

func (b *BannerPatternRegistryEntry) GetName() string {
	return b.Name
}

func (b *BannerPatternRegistryEntry) SetName(name string) {
	b.Name = name
}

func (b *BannerPatternRegistryEntry) GetBeautifiedName() string {
	return b.BeautifiedName
}

func (b *BannerPatternRegistryEntry) SetBeautifiedName(beautifiedName string) {
	b.BeautifiedName = beautifiedName
}

func (b *BannerPatternRegistryEntry) GetAsNBT() nbt.Nbt {
	nbtData := nbt.NewNbt(
		b.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{
				"asset_id":        b.GetAssetId(),
				"translation_key": b.GetTranslationKey(),
			}
			return data
		}()),
	)
	return *nbtData
}

var (
	Base = BannerPatternRegistryEntry{
		BeautifiedName: "Base",
		Name:           "base",
		AssetId:        "minecraft:base",
		TranslationKey: "block.minecraft.banner.base",
	}

	Border = BannerPatternRegistryEntry{
		BeautifiedName: "Border",
		Name:           "border",
		AssetId:        "minecraft:border",
		TranslationKey: "block.minecraft.banner.border",
	}

	Bricks = BannerPatternRegistryEntry{
		BeautifiedName: "Bricks",
		Name:           "bricks",
		AssetId:        "minecraft:bricks",
		TranslationKey: "block.minecraft.banner.bricks",
	}

	Circle = BannerPatternRegistryEntry{
		BeautifiedName: "Circle",
		Name:           "circle",
		AssetId:        "minecraft:circle",
		TranslationKey: "block.minecraft.banner.circle",
	}

	Creeper = BannerPatternRegistryEntry{
		BeautifiedName: "Creeper",
		Name:           "creeper",
		AssetId:        "minecraft:creeper",
		TranslationKey: "block.minecraft.banner.creeper",
	}

	Cross = BannerPatternRegistryEntry{
		BeautifiedName: "Cross",
		Name:           "cross",
		AssetId:        "minecraft:cross",
		TranslationKey: "block.minecraft.banner.cross",
	}

	CurlyBorder = BannerPatternRegistryEntry{
		BeautifiedName: "CurlyBorder",
		Name:           "curly_border",
		AssetId:        "minecraft:curly_border",
		TranslationKey: "block.minecraft.banner.curly_border",
	}

	DiagonalLeft = BannerPatternRegistryEntry{
		BeautifiedName: "DiagonalLeft",
		Name:           "diagonal_left",
		AssetId:        "minecraft:diagonal_left",
		TranslationKey: "block.minecraft.banner.diagonal_left",
	}

	DiagonalRight = BannerPatternRegistryEntry{
		BeautifiedName: "DiagonalRight",
		Name:           "diagonal_right",
		AssetId:        "minecraft:diagonal_right",
		TranslationKey: "block.minecraft.banner.diagonal_right",
	}

	DiagonalUpLeft = BannerPatternRegistryEntry{
		BeautifiedName: "DiagonalUpLeft",
		Name:           "diagonal_up_left",
		AssetId:        "minecraft:diagonal_up_left",
		TranslationKey: "block.minecraft.banner.diagonal_up_left",
	}

	DiagonalUpRight = BannerPatternRegistryEntry{
		BeautifiedName: "DiagonalUpRight",
		Name:           "diagonal_up_right",
		AssetId:        "minecraft:diagonal_up_right",
		TranslationKey: "block.minecraft.banner.diagonal_up_right",
	}

	Flow = BannerPatternRegistryEntry{
		BeautifiedName: "Flow",
		Name:           "flow",
		AssetId:        "minecraft:flow",
		TranslationKey: "block.minecraft.banner.flow",
	}

	Flower = BannerPatternRegistryEntry{
		BeautifiedName: "Flower",
		Name:           "flower",
		AssetId:        "minecraft:flower",
		TranslationKey: "block.minecraft.banner.flower",
	}

	Globe = BannerPatternRegistryEntry{
		BeautifiedName: "Globe",
		Name:           "globe",
		AssetId:        "minecraft:globe",
		TranslationKey: "block.minecraft.banner.globe",
	}

	Gradient = BannerPatternRegistryEntry{
		BeautifiedName: "Gradient",
		Name:           "gradient",
		AssetId:        "minecraft:gradient",
		TranslationKey: "block.minecraft.banner.gradient",
	}

	GradientUp = BannerPatternRegistryEntry{
		BeautifiedName: "GradientUp",
		Name:           "gradient_up",
		AssetId:        "minecraft:gradient_up",
		TranslationKey: "block.minecraft.banner.gradient_up",
	}

	Guster = BannerPatternRegistryEntry{
		BeautifiedName: "Guster",
		Name:           "guster",
		AssetId:        "minecraft:guster",
		TranslationKey: "block.minecraft.banner.guster",
	}

	HalfHorizontal = BannerPatternRegistryEntry{
		BeautifiedName: "HalfHorizontal",
		Name:           "half_horizontal",
		AssetId:        "minecraft:half_horizontal",
		TranslationKey: "block.minecraft.banner.half_horizontal",
	}

	HalfHorizontalBottom = BannerPatternRegistryEntry{
		BeautifiedName: "HalfHorizontalBottom",
		Name:           "half_horizontal_bottom",
		AssetId:        "minecraft:half_horizontal_bottom",
		TranslationKey: "block.minecraft.banner.half_horizontal_bottom",
	}

	HalfVertical = BannerPatternRegistryEntry{
		BeautifiedName: "HalfVertical",
		Name:           "half_vertical",
		AssetId:        "minecraft:half_vertical",
		TranslationKey: "block.minecraft.banner.half_vertical",
	}

	HalfVerticalRight = BannerPatternRegistryEntry{
		BeautifiedName: "HalfVerticalRight",
		Name:           "half_vertical_right",
		AssetId:        "minecraft:half_vertical_right",
		TranslationKey: "block.minecraft.banner.half_vertical_right",
	}

	Mojang = BannerPatternRegistryEntry{
		BeautifiedName: "Mojang",
		Name:           "mojang",
		AssetId:        "minecraft:mojang",
		TranslationKey: "block.minecraft.banner.mojang",
	}

	Piglin = BannerPatternRegistryEntry{
		BeautifiedName: "Piglin",
		Name:           "piglin",
		AssetId:        "minecraft:piglin",
		TranslationKey: "block.minecraft.banner.piglin",
	}

	Rhombus = BannerPatternRegistryEntry{
		BeautifiedName: "Rhombus",
		Name:           "rhombus",
		AssetId:        "minecraft:rhombus",
		TranslationKey: "block.minecraft.banner.rhombus",
	}

	Skull = BannerPatternRegistryEntry{
		BeautifiedName: "Skull",
		Name:           "skull",
		AssetId:        "minecraft:skull",
		TranslationKey: "block.minecraft.banner.skull",
	}

	SmallStripes = BannerPatternRegistryEntry{
		BeautifiedName: "SmallStripes",
		Name:           "small_stripes",
		AssetId:        "minecraft:small_stripes",
		TranslationKey: "block.minecraft.banner.small_stripes",
	}

	SquareBottomLeft = BannerPatternRegistryEntry{
		BeautifiedName: "SquareBottomLeft",
		Name:           "square_bottom_left",
		AssetId:        "minecraft:square_bottom_left",
		TranslationKey: "block.minecraft.banner.square_bottom_left",
	}

	SquareBottomRight = BannerPatternRegistryEntry{
		BeautifiedName: "SquareBottomRight",
		Name:           "square_bottom_right",
		AssetId:        "minecraft:square_bottom_right",
		TranslationKey: "block.minecraft.banner.square_bottom_right",
	}

	SquareTopLeft = BannerPatternRegistryEntry{
		BeautifiedName: "SquareTopLeft",
		Name:           "square_top_left",
		AssetId:        "minecraft:square_top_left",
		TranslationKey: "block.minecraft.banner.square_top_left",
	}

	SquareTopRight = BannerPatternRegistryEntry{
		BeautifiedName: "SquareTopRight",
		Name:           "square_top_right",
		AssetId:        "minecraft:square_top_right",
		TranslationKey: "block.minecraft.banner.square_top_right",
	}

	StraightCross = BannerPatternRegistryEntry{
		BeautifiedName: "StraightCross",
		Name:           "straight_cross",
		AssetId:        "minecraft:straight_cross",
		TranslationKey: "block.minecraft.banner.straight_cross",
	}

	StripeBottom = BannerPatternRegistryEntry{
		BeautifiedName: "StripeBottom",
		Name:           "stripe_bottom",
		AssetId:        "minecraft:stripe_bottom",
		TranslationKey: "block.minecraft.banner.stripe_bottom",
	}

	StripeCenter = BannerPatternRegistryEntry{
		BeautifiedName: "StripeCenter",
		Name:           "stripe_center",
		AssetId:        "minecraft:stripe_center",
		TranslationKey: "block.minecraft.banner.stripe_center",
	}

	StripeDownleft = BannerPatternRegistryEntry{
		BeautifiedName: "StripeDownleft",
		Name:           "stripe_downleft",
		AssetId:        "minecraft:stripe_downleft",
		TranslationKey: "block.minecraft.banner.stripe_downleft",
	}

	StripeDownright = BannerPatternRegistryEntry{
		BeautifiedName: "StripeDownright",
		Name:           "stripe_downright",
		AssetId:        "minecraft:stripe_downright",
		TranslationKey: "block.minecraft.banner.stripe_downright",
	}

	StripeLeft = BannerPatternRegistryEntry{
		BeautifiedName: "StripeLeft",
		Name:           "stripe_left",
		AssetId:        "minecraft:stripe_left",
		TranslationKey: "block.minecraft.banner.stripe_left",
	}

	StripeMiddle = BannerPatternRegistryEntry{
		BeautifiedName: "StripeMiddle",
		Name:           "stripe_middle",
		AssetId:        "minecraft:stripe_middle",
		TranslationKey: "block.minecraft.banner.stripe_middle",
	}

	StripeRight = BannerPatternRegistryEntry{
		BeautifiedName: "StripeRight",
		Name:           "stripe_right",
		AssetId:        "minecraft:stripe_right",
		TranslationKey: "block.minecraft.banner.stripe_right",
	}

	StripeTop = BannerPatternRegistryEntry{
		BeautifiedName: "StripeTop",
		Name:           "stripe_top",
		AssetId:        "minecraft:stripe_top",
		TranslationKey: "block.minecraft.banner.stripe_top",
	}

	TrianglesBottom = BannerPatternRegistryEntry{
		BeautifiedName: "TrianglesBottom",
		Name:           "triangles_bottom",
		AssetId:        "minecraft:triangles_bottom",
		TranslationKey: "block.minecraft.banner.triangles_bottom",
	}

	TrianglesTop = BannerPatternRegistryEntry{
		BeautifiedName: "TrianglesTop",
		Name:           "triangles_top",
		AssetId:        "minecraft:triangles_top",
		TranslationKey: "block.minecraft.banner.triangles_top",
	}

	TriangleBottom = BannerPatternRegistryEntry{
		BeautifiedName: "TriangleBottom",
		Name:           "triangle_bottom",
		AssetId:        "minecraft:triangle_bottom",
		TranslationKey: "block.minecraft.banner.triangle_bottom",
	}

	TriangleTop = BannerPatternRegistryEntry{
		BeautifiedName: "TriangleTop",
		Name:           "triangle_top",
		AssetId:        "minecraft:triangle_top",
		TranslationKey: "block.minecraft.banner.triangle_top",
	}
)
