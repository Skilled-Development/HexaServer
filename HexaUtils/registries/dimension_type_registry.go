package registries

import "HexaUtils/nbt"

var DimensionTypeRegistryInstance *DimensionTypeRegistry

type DimensionTypeRegistry struct {
	Name                         string
	DimensionTypeRegistryEntries []DimensionTypeRegistryEntry
}

func (d DimensionTypeRegistry) GetName() string {
	return d.Name
}

func (d DimensionTypeRegistry) GetEntries() []DimensionTypeRegistryEntry {
	return d.DimensionTypeRegistryEntries
}

func (d DimensionTypeRegistry) GetEntryByName(name string) *DimensionTypeRegistryEntry {
	for _, entry := range d.DimensionTypeRegistryEntries {
		if entry.GetName() == name {
			return &entry
		}
	}
	return nil
}

func (d DimensionTypeRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range d.DimensionTypeRegistryEntries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

func NewDimensionTypeRegistry() *DimensionTypeRegistry {
	dimensionRegistry := &DimensionTypeRegistry{
		Name: "dimension_type",
		DimensionTypeRegistryEntries: []DimensionTypeRegistryEntry{
			Overworld,
			OverworldCaves,
			TheEndDimension,
			TheNether,
		},
	}
	DimensionTypeRegistryInstance = dimensionRegistry
	return dimensionRegistry
}

type DimensionTypeRegistryEntry struct {
	BeautifiedName              string
	Name                        string
	FixedTime                   float64 //optional
	HasSkylight                 bool
	HasCeiling                  bool //optional
	Ultrawarm                   bool //optional
	Natural                     bool
	CoordinateScale             float32 //optional
	BedWorks                    bool    //optional
	RespawnAnchorWorks          bool    //optional
	MinY                        int32
	Height                      int32
	LogicalHeight               int32  //optional
	InfiniBurn                  string //optional
	Effects                     string
	AmbientLight                float32
	PiglinSafe                  bool
	HasRaids                    bool        //optional
	MonsterSpawnLightLevel      interface{} //Could be int or compound
	MonsterSpawnBlockLightLimit int32
}

func (d *DimensionTypeRegistryEntry) GetName() string {
	return d.Name
}

func (d *DimensionTypeRegistryEntry) SetName(value string) {
	d.Name = value
}

func (d *DimensionTypeRegistryEntry) GetBeautifiedName() string {
	return d.BeautifiedName
}

func (d *DimensionTypeRegistryEntry) SetBeautifiedName(value string) {
	d.BeautifiedName = value
}

func (d *DimensionTypeRegistryEntry) GetFixedTime() float64 {
	return d.FixedTime
}

func (d *DimensionTypeRegistryEntry) SetFixedTime(value float64) {
	d.FixedTime = value
}

func (d *DimensionTypeRegistryEntry) GetHasSkylight() bool {
	return d.HasSkylight
}

func (d *DimensionTypeRegistryEntry) SetHasSkylight(value bool) {
	d.HasSkylight = value
}

func (d *DimensionTypeRegistryEntry) GetHasCeiling() bool {
	return d.HasCeiling
}

func (d *DimensionTypeRegistryEntry) SetHasCeiling(value bool) {
	d.HasCeiling = value
}

func (d *DimensionTypeRegistryEntry) GetUltrawarm() bool {
	return d.Ultrawarm
}

func (d *DimensionTypeRegistryEntry) SetUltrawarm(value bool) {
	d.Ultrawarm = value
}

func (d *DimensionTypeRegistryEntry) GetNatural() bool {
	return d.Natural
}

func (d *DimensionTypeRegistryEntry) SetNatural(value bool) {
	d.Natural = value
}

func (d *DimensionTypeRegistryEntry) GetCoordinateScale() float32 {
	return d.CoordinateScale
}

func (d *DimensionTypeRegistryEntry) SetCoordinateScale(value float32) {
	d.CoordinateScale = value
}

func (d *DimensionTypeRegistryEntry) GetBedWorks() bool {
	return d.BedWorks
}

func (d *DimensionTypeRegistryEntry) SetBedWorks(value bool) {
	d.BedWorks = value
}

func (d *DimensionTypeRegistryEntry) GetRespawnAnchorWorks() bool {
	return d.RespawnAnchorWorks
}

func (d *DimensionTypeRegistryEntry) SetRespawnAnchorWorks(value bool) {
	d.RespawnAnchorWorks = value
}

func (d *DimensionTypeRegistryEntry) GetMinY() int32 {
	return d.MinY
}

func (d *DimensionTypeRegistryEntry) SetMinY(value int32) {
	d.MinY = value
}

func (d *DimensionTypeRegistryEntry) GetHeight() int32 {
	return d.Height
}

func (d *DimensionTypeRegistryEntry) SetHeight(value int32) {
	d.Height = value
}

func (d *DimensionTypeRegistryEntry) GetLogicalHeight() int32 {
	return d.LogicalHeight
}

func (d *DimensionTypeRegistryEntry) SetLogicalHeight(value int32) {
	d.LogicalHeight = value
}

func (d *DimensionTypeRegistryEntry) GetInfiniBurn() string {
	return d.InfiniBurn
}

func (d *DimensionTypeRegistryEntry) SetInfiniBurn(value string) {
	d.InfiniBurn = value
}

func (d *DimensionTypeRegistryEntry) GetEffects() string {
	return d.Effects
}

func (d *DimensionTypeRegistryEntry) SetEffects(value string) {
	d.Effects = value
}

func (d *DimensionTypeRegistryEntry) GetAmbientLight() float32 {
	return d.AmbientLight
}

func (d *DimensionTypeRegistryEntry) SetAmbientLight(value float32) {
	d.AmbientLight = value
}

func (d *DimensionTypeRegistryEntry) GetPiglinSafe() bool {
	return d.PiglinSafe
}

func (d *DimensionTypeRegistryEntry) SetPiglinSafe(value bool) {
	d.PiglinSafe = value
}

func (d *DimensionTypeRegistryEntry) GetHasRaids() bool {
	return d.HasRaids
}

func (d *DimensionTypeRegistryEntry) SetHasRaids(value bool) {
	d.HasRaids = value
}

func (d *DimensionTypeRegistryEntry) GetMonsterSpawnLightLevel() interface{} {
	return d.MonsterSpawnLightLevel
}

func (d *DimensionTypeRegistryEntry) SetMonsterSpawnLightLevel(value interface{}) {
	d.MonsterSpawnLightLevel = value
}

func (d *DimensionTypeRegistryEntry) GetMonsterSpawnBlockLightLimit() int32 {
	return d.MonsterSpawnBlockLightLimit
}

func (d *DimensionTypeRegistryEntry) SetMonsterSpawnBlockLightLimit(value int32) {
	d.MonsterSpawnBlockLightLimit = value
}

func (d DimensionTypeRegistryEntry) GetAsNBT() nbt.Nbt {

	nbtData := nbt.NewNbt(
		d.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{}
			data["ambient_light"] = d.GetAmbientLight()
			data["bed_works"] = d.GetBedWorks()
			data["coordinate_scale"] = d.GetCoordinateScale()
			data["effects"] = d.GetEffects()
			data["has_ceiling"] = d.GetHasCeiling()
			data["has_raids"] = d.GetHasRaids()
			data["has_skylight"] = d.GetHasSkylight()
			data["height"] = d.GetHeight()
			data["infiniburn"] = d.GetInfiniBurn()
			data["logical_height"] = d.GetLogicalHeight()
			data["min_y"] = d.GetMinY()
			data["monster_spawn_block_light_limit"] = d.GetMonsterSpawnBlockLightLimit()
			data["monster_spawn_light_level"] = d.GetMonsterSpawnLightLevel()
			data["natural"] = d.GetNatural()
			data["piglin_safe"] = d.GetPiglinSafe()
			data["respawn_anchor_works"] = d.GetRespawnAnchorWorks()
			data["ultrawarm"] = d.GetUltrawarm()
			data["fixed_time"] = d.GetFixedTime()
			return data
		}()),
	)

	return *nbtData
}

var (
	Overworld = DimensionTypeRegistryEntry{
		BeautifiedName:              "Overworld",
		Name:                        "overworld",
		AmbientLight:                0.0,
		BedWorks:                    true,
		CoordinateScale:             1.0,
		Effects:                     "minecraft:overworld",
		HasCeiling:                  false,
		HasRaids:                    true,
		HasSkylight:                 true,
		Height:                      384,
		InfiniBurn:                  "#minecraft:infiniburn_overworld",
		LogicalHeight:               384,
		MinY:                        -64,
		MonsterSpawnBlockLightLimit: 0,
		MonsterSpawnLightLevel:      map[string]interface{}{"type": "minecraft:uniform", "max_inclusive": 7, "min_inclusive": 0},
		Natural:                     true,
		PiglinSafe:                  false,
		RespawnAnchorWorks:          false,
		Ultrawarm:                   false,
	}

	OverworldCaves = DimensionTypeRegistryEntry{
		BeautifiedName:              "OverworldCaves",
		Name:                        "overworld_caves",
		AmbientLight:                0.0,
		BedWorks:                    true,
		CoordinateScale:             1.0,
		Effects:                     "minecraft:overworld",
		HasCeiling:                  true,
		HasRaids:                    true,
		HasSkylight:                 true,
		Height:                      384,
		InfiniBurn:                  "#minecraft:infiniburn_overworld",
		LogicalHeight:               384,
		MinY:                        -64,
		MonsterSpawnBlockLightLimit: 0,
		MonsterSpawnLightLevel:      map[string]interface{}{"type": "minecraft:uniform", "max_inclusive": 7, "min_inclusive": 0},
		Natural:                     true,
		PiglinSafe:                  false,
		RespawnAnchorWorks:          false,
		Ultrawarm:                   false,
	}

	TheEndDimension = DimensionTypeRegistryEntry{
		BeautifiedName:              "TheEnd",
		Name:                        "the_end",
		AmbientLight:                0.0,
		BedWorks:                    false,
		CoordinateScale:             1.0,
		Effects:                     "minecraft:the_end",
		FixedTime:                   6000,
		HasCeiling:                  false,
		HasRaids:                    true,
		HasSkylight:                 false,
		Height:                      256,
		InfiniBurn:                  "#minecraft:infiniburn_end",
		LogicalHeight:               256,
		MinY:                        0,
		MonsterSpawnBlockLightLimit: 0,
		MonsterSpawnLightLevel:      map[string]interface{}{"type": "minecraft:uniform", "max_inclusive": 7, "min_inclusive": 0},
		Natural:                     false,
		PiglinSafe:                  false,
		RespawnAnchorWorks:          false,
		Ultrawarm:                   false,
	}

	TheNether = DimensionTypeRegistryEntry{
		BeautifiedName:              "TheNether",
		Name:                        "the_nether",
		AmbientLight:                0.1,
		BedWorks:                    false,
		CoordinateScale:             8.0,
		Effects:                     "minecraft:the_nether",
		FixedTime:                   18000,
		HasCeiling:                  true,
		HasRaids:                    false,
		HasSkylight:                 false,
		Height:                      256,
		InfiniBurn:                  "#minecraft:infiniburn_nether",
		LogicalHeight:               128,
		MinY:                        0,
		MonsterSpawnBlockLightLimit: 15,
		MonsterSpawnLightLevel:      map[string]interface{}{"type": "minecraft:uniform", "max_inclusive": 7, "min_inclusive": 7},
		Natural:                     false,
		PiglinSafe:                  true,
		RespawnAnchorWorks:          true,
		Ultrawarm:                   true,
	}
)
