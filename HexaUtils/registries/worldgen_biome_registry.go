package registries

import "HexaUtils/nbt"

var BiomeRegistryInstance *BiomeRegistry

// BiomeRegistry
type BiomeRegistry struct {
	Name                 string
	BiomeRegistryEntries []BiomeRegistryEntry
}

func NewBiomeRegistry() *BiomeRegistry {
	biomeRegistry := &BiomeRegistry{
		Name: "worldgen/biome",
		BiomeRegistryEntries: []BiomeRegistryEntry{
			Badlands,
			BambooJungle,
			BasaltDeltas,
			Beach,
			BirchForest,
			CherryGrove,
			ColdOcean,
			CrimsonForest,
			DarkForest,
			DeepColdOcean,
			DeepDark,
			DeepFrozenOcean,
			DeepLukewarmOcean,
			DeepOcean,
			Desert,
			DripstoneCaves,
			EndBarrens,
			EndHighlands,
			EndMidlands,
			ErodedBadlands,
			FlowerForest,
			Forest,
			FrozenOcean,
			FrozenPeaks,
			FrozenRiver,
			Grove,
			IceSpikes,
			JaggedPeaks,
			Jungle,
			LukewarmOcean,
			LushCaves,
			MangroveSwamp,
			Meadow,
			MushroomFields,
			NetherWastes,
			Ocean,
			OldGrowthBirchForest,
			OldGrowthPineTaiga,
			OldGrowthSpruceTaiga,
			Plains,
			River,
			Savanna,
			SavannaPlateau,
			SmallEndIslands,
			SnowyBeach,
			SnowyPlains,
			SnowySlopes,
			SnowyTaiga,
			SoulSandValley,
			SparseJungle,
			StonyPeaks,
			StonyShore,
			SunflowerPlains,
			Swamp,
			Taiga,
			TheEnd,
			TheVoid,
			WarmOcean,
			WarpedForest,
			WindsweptForest,
			WindsweptGravellyHills,
			WindsweptHills,
			WindsweptSavanna,
			WoodedBadlands,
		},
	}
	BiomeRegistryInstance = biomeRegistry
	return BiomeRegistryInstance
}

func (b BiomeRegistry) GetEntries() []BiomeRegistryEntry {
	return b.BiomeRegistryEntries
}

func (b BiomeRegistry) GetEntryByName(name string) *BiomeRegistryEntry {
	for _, entry := range b.BiomeRegistryEntries {
		if entry.GetName() == name {
			return &entry
		}
	}
	return nil
}

func (b BiomeRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range b.BiomeRegistryEntries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}
func (b BiomeRegistry) GetName() string {
	return b.Name
}

// BiomeRegistryEntry
type BiomeRegistryEntry struct {
	BeautifiedName      string
	Name                string
	HasPrecipitation    bool
	Temperature         float32
	TemperatureModifier string
	Downfall            float32
	Effects             BiomeEffects
}

// Getters and Setters for BiomeRegistryEntry
func (b *BiomeRegistryEntry) GetBeautifiedName() string {
	return b.BeautifiedName
}
func (b *BiomeRegistryEntry) SetBeautifiedName(name string) {
	b.BeautifiedName = name
}
func (b *BiomeRegistryEntry) GetName() string {
	return b.Name
}
func (b *BiomeRegistryEntry) SetName(name string) {
	b.Name = name
}

func (b *BiomeRegistryEntry) GetHasPrecipitation() bool {
	return b.HasPrecipitation
}
func (b *BiomeRegistryEntry) SetHasPrecipitation(hasPrecipitation bool) {
	b.HasPrecipitation = hasPrecipitation
}

func (b *BiomeRegistryEntry) GetTemperature() float32 {
	return b.Temperature
}
func (b *BiomeRegistryEntry) SetTemperature(temp float32) {
	b.Temperature = temp
}

func (b *BiomeRegistryEntry) GetTemperatureModifier() string {
	return b.TemperatureModifier
}
func (b *BiomeRegistryEntry) SetTemperatureModifier(modifier string) {
	b.TemperatureModifier = modifier
}

func (b *BiomeRegistryEntry) GetDownfall() float32 {
	return b.Downfall
}

func (b *BiomeRegistryEntry) SetDownfall(downfall float32) {
	b.Downfall = downfall
}
func (b *BiomeRegistryEntry) GetEffects() BiomeEffects {
	return b.Effects
}
func (b *BiomeRegistryEntry) SetEffects(effects BiomeEffects) {
	b.Effects = effects
}

func (b *BiomeRegistryEntry) GetAsNBT() nbt.Nbt {
	nbtData := nbt.NewNbt(
		b.GetName(),
		nbt.NbtCompoundFromInterfaceMap(func() map[string]interface{} {
			data := map[string]interface{}{}
			if b.GetHasPrecipitation() {
				data["has_precipitation"] = 1
			} else {
				data["has_precipitation"] = 0
			}
			data["temperature"] = b.GetTemperature()
			if b.GetTemperatureModifier() != "" {
				data["temperature_modifier"] = b.GetTemperatureModifier()
			}
			data["downfall"] = b.GetDownfall()
			data["effects"] = b.GetEffects().GetAsMap()

			return data
		}()),
	)
	return *nbtData

}

// BiomeEffects
type BiomeEffects struct {
	FogColor           int
	WaterColor         int
	WaterFogColor      int
	SkyColor           int
	FoliageColor       *int
	GrassColor         *int
	GrassColorModifier string
	Particle           *BiomeParticle
	AmbientSound       interface{} //Can be String or Compound
	MoodSound          *BiomeMoodSound
	AdditionsSound     *BiomeAdditionsSound
	Music              *BiomeMusic
}

// Getters and Setters for BiomeEffects
func (be *BiomeEffects) GetFogColor() int {
	return be.FogColor
}
func (be *BiomeEffects) SetFogColor(fogColor int) {
	be.FogColor = fogColor
}

func (be *BiomeEffects) GetWaterColor() int {
	return be.WaterColor
}

func (be *BiomeEffects) SetWaterColor(waterColor int) {
	be.WaterColor = waterColor
}

func (be *BiomeEffects) GetWaterFogColor() int {
	return be.WaterFogColor
}

func (be *BiomeEffects) SetWaterFogColor(waterFogColor int) {
	be.WaterFogColor = waterFogColor
}

func (be *BiomeEffects) GetSkyColor() int {
	return be.SkyColor
}

func (be *BiomeEffects) SetSkyColor(skyColor int) {
	be.SkyColor = skyColor
}
func (be *BiomeEffects) GetFoliageColor() *int {
	return be.FoliageColor
}

func (be *BiomeEffects) SetFoliageColor(foliageColor *int) {
	be.FoliageColor = foliageColor
}

func (be *BiomeEffects) GetGrassColor() *int {
	return be.GrassColor
}
func (be *BiomeEffects) SetGrassColor(grassColor *int) {
	be.GrassColor = grassColor
}

func (be *BiomeEffects) GetGrassColorModifier() string {
	return be.GrassColorModifier
}

func (be *BiomeEffects) SetGrassColorModifier(modifier string) {
	be.GrassColorModifier = modifier
}
func (be *BiomeEffects) GetParticle() *BiomeParticle {
	return be.Particle
}
func (be *BiomeEffects) SetParticle(particle *BiomeParticle) {
	be.Particle = particle
}
func (be *BiomeEffects) GetAmbientSound() interface{} {
	return be.AmbientSound
}

func (be *BiomeEffects) SetAmbientSound(sound interface{}) {
	be.AmbientSound = sound
}

func (be *BiomeEffects) GetMoodSound() *BiomeMoodSound {
	return be.MoodSound
}

func (be *BiomeEffects) SetMoodSound(moodSound *BiomeMoodSound) {
	be.MoodSound = moodSound
}

func (be *BiomeEffects) GetAdditionsSound() *BiomeAdditionsSound {
	return be.AdditionsSound
}

func (be *BiomeEffects) SetAdditionsSound(additionsSound *BiomeAdditionsSound) {
	be.AdditionsSound = additionsSound
}

func (be *BiomeEffects) GetMusic() *BiomeMusic {
	return be.Music
}

func (be *BiomeEffects) SetMusic(music *BiomeMusic) {
	be.Music = music
}

func (be BiomeEffects) GetAsMap() map[string]interface{} {
	data := map[string]interface{}{
		"fog_color":       be.GetFogColor(),
		"water_color":     be.GetWaterColor(),
		"water_fog_color": be.GetWaterFogColor(),
		"sky_color":       be.GetSkyColor(),
	}
	if be.GetFoliageColor() != nil {
		data["foliage_color"] = *be.GetFoliageColor()
	}
	if be.GetGrassColor() != nil {
		data["grass_color"] = *be.GetGrassColor()
	}
	if be.GetGrassColorModifier() != "" {
		data["grass_color_modifier"] = be.GetGrassColorModifier()
	}
	if be.GetParticle() != nil {
		data["particle"] = be.GetParticle().GetAsMap()
	}
	if sound, ok := be.GetAmbientSound().(string); ok {
		data["ambient_sound"] = sound
	} else if sound, ok := be.GetAmbientSound().(map[string]interface{}); ok {
		data["ambient_sound"] = sound
	}
	if be.GetMoodSound() != nil {
		data["mood_sound"] = be.GetMoodSound().GetAsMap()
	}
	if be.GetAdditionsSound() != nil {
		data["additions_sound"] = be.GetAdditionsSound().GetAsMap()
	}
	if be.GetMusic() != nil {
		data["music"] = be.GetMusic().GetAsMap()
	}
	return data
}

// BiomeParticle
type BiomeParticle struct {
	Options     map[string]interface{}
	Probability float32
}

// Getters and Setters for BiomeParticle
func (bp *BiomeParticle) GetOptions() map[string]interface{} {
	return bp.Options
}
func (bp *BiomeParticle) SetOptions(options map[string]interface{}) {
	bp.Options = options
}
func (bp *BiomeParticle) GetProbability() float32 {
	return bp.Probability
}
func (bp *BiomeParticle) SetProbability(probability float32) {
	bp.Probability = probability
}
func (bp *BiomeParticle) GetAsMap() map[string]interface{} {
	return map[string]interface{}{
		"options":     bp.GetOptions(),
		"probability": bp.GetProbability(),
	}

}

// BiomeMoodSound
type BiomeMoodSound struct {
	Sound             string
	TickDelay         int
	BlockSearchExtent int
	Offset            float64
}

// Getters and Setters for BiomeMoodSound
func (bms *BiomeMoodSound) GetSound() string {
	return bms.Sound
}

func (bms *BiomeMoodSound) SetSound(sound string) {
	bms.Sound = sound
}

func (bms *BiomeMoodSound) GetTickDelay() int {
	return bms.TickDelay
}

func (bms *BiomeMoodSound) SetTickDelay(tickDelay int) {
	bms.TickDelay = tickDelay
}
func (bms *BiomeMoodSound) GetBlockSearchExtent() int {
	return bms.BlockSearchExtent
}

func (bms *BiomeMoodSound) SetBlockSearchExtent(extent int) {
	bms.BlockSearchExtent = extent
}

func (bms *BiomeMoodSound) GetOffset() float64 {
	return bms.Offset
}

func (bms *BiomeMoodSound) SetOffset(offset float64) {
	bms.Offset = offset
}
func (bms *BiomeMoodSound) GetAsMap() map[string]interface{} {
	data := map[string]interface{}{
		"sound":               bms.GetSound(),
		"tick_delay":          bms.GetTickDelay(),
		"block_search_extent": bms.GetBlockSearchExtent(),
		"offset":              bms.GetOffset(),
	}
	return data
}

// BiomeAdditionsSound
type BiomeAdditionsSound struct {
	Sound      string
	TickChance float64
}

// Getters and Setters for BiomeAdditionsSound
func (bas *BiomeAdditionsSound) GetSound() string {
	return bas.Sound
}
func (bas *BiomeAdditionsSound) SetSound(sound string) {
	bas.Sound = sound
}

func (bas *BiomeAdditionsSound) GetTickChance() float64 {
	return bas.TickChance
}
func (bas *BiomeAdditionsSound) SetTickChance(tickChance float64) {
	bas.TickChance = tickChance
}
func (bas *BiomeAdditionsSound) GetAsMap() map[string]interface{} {
	data := map[string]interface{}{
		"sound":       bas.GetSound(),
		"tick_chance": bas.GetTickChance(),
	}
	return data
}

// BiomeMusic
type BiomeMusic struct {
	Sound               string
	MinDelay            int
	MaxDelay            int
	ReplaceCurrentMusic bool
}

// Getters and Setters for BiomeMusic
func (bm *BiomeMusic) GetSound() string {
	return bm.Sound
}

func (bm *BiomeMusic) SetSound(sound string) {
	bm.Sound = sound
}

func (bm *BiomeMusic) GetMinDelay() int {
	return bm.MinDelay
}

func (bm *BiomeMusic) SetMinDelay(minDelay int) {
	bm.MinDelay = minDelay
}

func (bm *BiomeMusic) GetMaxDelay() int {
	return bm.MaxDelay
}

func (bm *BiomeMusic) SetMaxDelay(maxDelay int) {
	bm.MaxDelay = maxDelay
}

func (bm *BiomeMusic) GetReplaceCurrentMusic() bool {
	return bm.ReplaceCurrentMusic
}
func (bm *BiomeMusic) SetReplaceCurrentMusic(replaceCurrentMusic bool) {
	bm.ReplaceCurrentMusic = replaceCurrentMusic
}

func (bm *BiomeMusic) GetAsMap() map[string]interface{} {
	data := map[string]interface{}{
		"sound":     bm.GetSound(),
		"min_delay": bm.GetMinDelay(),
		"max_delay": bm.GetMaxDelay(),
	}
	if bm.GetReplaceCurrentMusic() {
		data["replace_current_music"] = 1
	} else {
		data["replace_current_music"] = 0
	}
	return data
}

var (
	Badlands = BiomeRegistryEntry{
		BeautifiedName:      "Badlands",
		Name:                "badlands",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.badlands",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	BambooJungle = BiomeRegistryEntry{
		BeautifiedName:      "BambooJungle",
		Name:                "bamboo_jungle",
		HasPrecipitation:    true,
		Temperature:         0.95,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7842047,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.bamboo_jungle",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	BasaltDeltas = BiomeRegistryEntry{
		BeautifiedName:      "BasaltDeltas",
		Name:                "basalt_deltas",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      6840176,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			Particle: &BiomeParticle{
				Options: map[string]interface{}{
					"type": "minecraft:white_ash",
				},
				Probability: 0.118093334,
			},
			AmbientSound: "minecraft:ambient.basalt_deltas.loop",
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.basalt_deltas.mood",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			AdditionsSound: &BiomeAdditionsSound{
				Sound:      "minecraft:ambient.basalt_deltas.additions",
				TickChance: 0.0111,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.nether.basalt_deltas",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Beach = BiomeRegistryEntry{
		BeautifiedName:      "Beach",
		Name:                "beach",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.4,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	BirchForest = BiomeRegistryEntry{
		BeautifiedName:      "BirchForest",
		Name:                "birch_forest",
		HasPrecipitation:    true,
		Temperature:         0.6,
		TemperatureModifier: "none",
		Downfall:            0.6,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8037887,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	CherryGrove = BiomeRegistryEntry{
		BeautifiedName:      "CherryGrove",
		Name:                "cherry_grove",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    6141935,
			WaterFogColor: 6141935,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.cherry_grove",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	ColdOcean = BiomeRegistryEntry{
		BeautifiedName:      "ColdOcean",
		Name:                "cold_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4020182,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	CrimsonForest = BiomeRegistryEntry{
		BeautifiedName:      "CrimsonForest",
		Name:                "crimson_forest",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      3343107,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			Particle: &BiomeParticle{
				Options: map[string]interface{}{
					"type": "minecraft:crimson_spore",
				},
				Probability: 0.025,
			},
			AmbientSound: "minecraft:ambient.crimson_forest.loop",
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.crimson_forest.mood",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			AdditionsSound: &BiomeAdditionsSound{
				Sound:      "minecraft:ambient.crimson_forest.additions",
				TickChance: 0.0111,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.nether.crimson_forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	DarkForest = BiomeRegistryEntry{
		BeautifiedName:      "DarkForest",
		Name:                "dark_forest",
		HasPrecipitation:    true,
		Temperature:         0.7,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7972607,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	DeepColdOcean = BiomeRegistryEntry{
		BeautifiedName:      "DeepColdOcean",
		Name:                "deep_cold_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4020182,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	DeepDark = BiomeRegistryEntry{
		BeautifiedName:      "DeepDark",
		Name:                "deep_dark",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.4,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.deep_dark",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	DeepFrozenOcean = BiomeRegistryEntry{
		BeautifiedName:      "DeepFrozenOcean",
		Name:                "deep_frozen_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "frozen",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    3750089,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	DeepLukewarmOcean = BiomeRegistryEntry{
		BeautifiedName:      "DeepLukewarmOcean",
		Name:                "deep_lukewarm_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4566514,
			WaterFogColor: 267827,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	DeepOcean = BiomeRegistryEntry{
		BeautifiedName:      "DeepOcean",
		Name:                "deep_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	Desert = BiomeRegistryEntry{
		BeautifiedName:      "Desert",
		Name:                "desert",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.desert",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	DripstoneCaves = BiomeRegistryEntry{
		BeautifiedName:      "DripstoneCaves",
		Name:                "dripstone_caves",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.4,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.dripstone_caves",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	EndBarrens = BiomeRegistryEntry{
		BeautifiedName:      "EndBarrens",
		Name:                "end_barrens",
		HasPrecipitation:    false,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      10518688,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      0,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	EndHighlands = BiomeRegistryEntry{
		BeautifiedName:      "EndHighlands",
		Name:                "end_highlands",
		HasPrecipitation:    false,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      10518688,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      0,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	EndMidlands = BiomeRegistryEntry{
		BeautifiedName:      "EndMidlands",
		Name:                "end_midlands",
		HasPrecipitation:    false,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      10518688,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      0,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	ErodedBadlands = BiomeRegistryEntry{
		BeautifiedName:      "ErodedBadlands",
		Name:                "eroded_badlands",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.badlands",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	FlowerForest = BiomeRegistryEntry{
		BeautifiedName:      "FlowerForest",
		Name:                "flower_forest",
		HasPrecipitation:    true,
		Temperature:         0.7,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7972607,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.flower_forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Forest = BiomeRegistryEntry{
		BeautifiedName:      "Forest",
		Name:                "forest",
		HasPrecipitation:    true,
		Temperature:         0.7,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7972607,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	FrozenOcean = BiomeRegistryEntry{
		BeautifiedName:      "FrozenOcean",
		Name:                "frozen_ocean",
		HasPrecipitation:    true,
		Temperature:         0.0,
		TemperatureModifier: "frozen",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    3750089,
			WaterFogColor: 329011,
			SkyColor:      8364543,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	FrozenPeaks = BiomeRegistryEntry{
		BeautifiedName:      "FrozenPeaks",
		Name:                "frozen_peaks",
		HasPrecipitation:    true,
		Temperature:         -0.7,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8756735,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.frozen_peaks",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	FrozenRiver = BiomeRegistryEntry{
		BeautifiedName:      "FrozenRiver",
		Name:                "frozen_river",
		HasPrecipitation:    true,
		Temperature:         0.0,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    3750089,
			WaterFogColor: 329011,
			SkyColor:      8364543,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	Grove = BiomeRegistryEntry{
		BeautifiedName:      "Grove",
		Name:                "grove",
		HasPrecipitation:    true,
		Temperature:         -0.2,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8495359,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.grove",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	IceSpikes = BiomeRegistryEntry{
		BeautifiedName:      "IceSpikes",
		Name:                "ice_spikes",
		HasPrecipitation:    true,
		Temperature:         0.0,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8364543,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	JaggedPeaks = BiomeRegistryEntry{
		BeautifiedName:      "JaggedPeaks",
		Name:                "jagged_peaks",
		HasPrecipitation:    true,
		Temperature:         -0.7,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8756735,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.jagged_peaks",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Jungle = BiomeRegistryEntry{
		BeautifiedName:      "Jungle",
		Name:                "jungle",
		HasPrecipitation:    true,
		Temperature:         0.95,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7842047,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.jungle",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	LukewarmOcean = BiomeRegistryEntry{
		BeautifiedName:      "LukewarmOcean",
		Name:                "lukewarm_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4566514,
			WaterFogColor: 267827,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	LushCaves = BiomeRegistryEntry{
		BeautifiedName:      "LushCaves",
		Name:                "lush_caves",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.lush_caves",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	MangroveSwamp = BiomeRegistryEntry{
		BeautifiedName:      "MangroveSwamp",
		Name:                "mangrove_swamp",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    3832426,
			WaterFogColor: 5077600,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.swamp",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Meadow = BiomeRegistryEntry{
		BeautifiedName:      "Meadow",
		Name:                "meadow",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    937679,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.meadow",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	MushroomFields = BiomeRegistryEntry{
		BeautifiedName:      "MushroomFields",
		Name:                "mushroom_fields",
		HasPrecipitation:    true,
		Temperature:         0.9,
		TemperatureModifier: "none",
		Downfall:            1.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7842047,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	NetherWastes = BiomeRegistryEntry{
		BeautifiedName:      "NetherWastes",
		Name:                "nether_wastes",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      3344392,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			AmbientSound:  "minecraft:ambient.nether_wastes.loop",
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.nether_wastes.mood",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			AdditionsSound: &BiomeAdditionsSound{
				Sound:      "minecraft:ambient.nether_wastes.additions",
				TickChance: 0.0111,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.nether.nether_wastes",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Ocean = BiomeRegistryEntry{
		BeautifiedName:      "Ocean",
		Name:                "ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	OldGrowthBirchForest = BiomeRegistryEntry{
		BeautifiedName:      "OldGrowthBirchForest",
		Name:                "old_growth_birch_forest",
		HasPrecipitation:    true,
		Temperature:         0.6,
		TemperatureModifier: "none",
		Downfall:            0.6,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8037887,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	OldGrowthPineTaiga = BiomeRegistryEntry{
		BeautifiedName:      "OldGrowthPineTaiga",
		Name:                "old_growth_pine_taiga",
		HasPrecipitation:    true,
		Temperature:         0.3,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8168447,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.old_growth_taiga",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	OldGrowthSpruceTaiga = BiomeRegistryEntry{
		BeautifiedName:      "OldGrowthSpruceTaiga",
		Name:                "old_growth_spruce_taiga",
		HasPrecipitation:    true,
		Temperature:         0.25,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8233983,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.old_growth_taiga",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Plains = BiomeRegistryEntry{
		BeautifiedName:      "Plains",
		Name:                "plains",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.4,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	River = BiomeRegistryEntry{
		BeautifiedName:      "River",
		Name:                "river",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	Savanna = BiomeRegistryEntry{
		BeautifiedName:      "Savanna",
		Name:                "savanna",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SavannaPlateau = BiomeRegistryEntry{
		BeautifiedName:      "SavannaPlateau",
		Name:                "savanna_plateau",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SmallEndIslands = BiomeRegistryEntry{
		BeautifiedName:      "SmallEndIslands",
		Name:                "small_end_islands",
		HasPrecipitation:    false,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      10518688,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      0,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SnowyBeach = BiomeRegistryEntry{
		BeautifiedName:      "SnowyBeach",
		Name:                "snowy_beach",
		HasPrecipitation:    true,
		Temperature:         0.05,
		TemperatureModifier: "none",
		Downfall:            0.3,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4020182,
			WaterFogColor: 329011,
			SkyColor:      8364543,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SnowyPlains = BiomeRegistryEntry{
		BeautifiedName:      "SnowyPlains",
		Name:                "snowy_plains",
		HasPrecipitation:    true,
		Temperature:         0.0,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8364543,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SnowySlopes = BiomeRegistryEntry{
		BeautifiedName:      "SnowySlopes",
		Name:                "snowy_slopes",
		HasPrecipitation:    true,
		Temperature:         -0.3,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8560639,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.snowy_slopes",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	SnowyTaiga = BiomeRegistryEntry{
		BeautifiedName:      "SnowyTaiga",
		Name:                "snowy_taiga",
		HasPrecipitation:    true,
		Temperature:         -0.5,
		TemperatureModifier: "none",
		Downfall:            0.4,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4020182,
			WaterFogColor: 329011,
			SkyColor:      8625919,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SoulSandValley = BiomeRegistryEntry{
		BeautifiedName:      "SoulSandValley",
		Name:                "soul_sand_valley",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      1787717,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			Particle: &BiomeParticle{
				Options: map[string]interface{}{
					"type": "minecraft:ash",
				},
				Probability: 0.00625,
			},
			AmbientSound: "minecraft:ambient.soul_sand_valley.loop",
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.soul_sand_valley.mood",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			AdditionsSound: &BiomeAdditionsSound{
				Sound:      "minecraft:ambient.soul_sand_valley.additions",
				TickChance: 0.0111,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.nether.soul_sand_valley",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	SparseJungle = BiomeRegistryEntry{
		BeautifiedName:      "SparseJungle",
		Name:                "sparse_jungle",
		HasPrecipitation:    true,
		Temperature:         0.95,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7842047,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.sparse_jungle",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	StonyPeaks = BiomeRegistryEntry{
		BeautifiedName:      "StonyPeaks",
		Name:                "stony_peaks",
		HasPrecipitation:    true,
		Temperature:         1.0,
		TemperatureModifier: "none",
		Downfall:            0.3,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7776511,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.stony_peaks",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	StonyShore = BiomeRegistryEntry{
		BeautifiedName:      "StonyShore",
		Name:                "stony_shore",
		HasPrecipitation:    true,
		Temperature:         0.2,
		TemperatureModifier: "none",
		Downfall:            0.3,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8233727,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	SunflowerPlains = BiomeRegistryEntry{
		BeautifiedName:      "SunflowerPlains",
		Name:                "sunflower_plains",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.4,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	Swamp = BiomeRegistryEntry{
		BeautifiedName:      "Swamp",
		Name:                "swamp",
		HasPrecipitation:    true,
		Temperature:         0.8,
		TemperatureModifier: "none",
		Downfall:            0.9,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    6388580,
			WaterFogColor: 2302743,
			SkyColor:      7907327,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.swamp",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	Taiga = BiomeRegistryEntry{
		BeautifiedName:      "Taiga",
		Name:                "taiga",
		HasPrecipitation:    true,
		Temperature:         0.25,
		TemperatureModifier: "none",
		Downfall:            0.8,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8233983,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	TheEnd = BiomeRegistryEntry{
		BeautifiedName:      "TheEnd",
		Name:                "the_end",
		HasPrecipitation:    false,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      10518688,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      0,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	TheVoid = BiomeRegistryEntry{
		BeautifiedName:      "TheVoid",
		Name:                "the_void",
		HasPrecipitation:    false,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	WarmOcean = BiomeRegistryEntry{
		BeautifiedName:      "WarmOcean",
		Name:                "warm_ocean",
		HasPrecipitation:    true,
		Temperature:         0.5,
		TemperatureModifier: "none",
		Downfall:            0.5,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4445678,
			WaterFogColor: 270131,
			SkyColor:      8103167,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	WarpedForest = BiomeRegistryEntry{
		BeautifiedName:      "WarpedForest",
		Name:                "warped_forest",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      1705242,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			Particle: &BiomeParticle{
				Options: map[string]interface{}{
					"type": "minecraft:warped_spore",
				},
				Probability: 0.01428,
			},
			AmbientSound: "minecraft:ambient.warped_forest.loop",
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.warped_forest.mood",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			AdditionsSound: &BiomeAdditionsSound{
				Sound:      "minecraft:ambient.warped_forest.additions",
				TickChance: 0.0111,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.nether.warped_forest",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}

	WindsweptForest = BiomeRegistryEntry{
		BeautifiedName:      "WindsweptForest",
		Name:                "windswept_forest",
		HasPrecipitation:    true,
		Temperature:         0.2,
		TemperatureModifier: "none",
		Downfall:            0.3,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8233727,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	WindsweptGravellyHills = BiomeRegistryEntry{
		BeautifiedName:      "WindsweptGravellyHills",
		Name:                "windswept_gravelly_hills",
		HasPrecipitation:    true,
		Temperature:         0.2,
		TemperatureModifier: "none",
		Downfall:            0.3,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8233727,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	WindsweptHills = BiomeRegistryEntry{
		BeautifiedName:      "WindsweptHills",
		Name:                "windswept_hills",
		HasPrecipitation:    true,
		Temperature:         0.2,
		TemperatureModifier: "none",
		Downfall:            0.3,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      8233727,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	WindsweptSavanna = BiomeRegistryEntry{
		BeautifiedName:      "WindsweptSavanna",
		Name:                "windswept_savanna",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
		},
	}

	WoodedBadlands = BiomeRegistryEntry{
		BeautifiedName:      "WoodedBadlands",
		Name:                "wooded_badlands",
		HasPrecipitation:    false,
		Temperature:         2.0,
		TemperatureModifier: "none",
		Downfall:            0.0,
		Effects: BiomeEffects{
			FogColor:      12638463,
			WaterColor:    4159204,
			WaterFogColor: 329011,
			SkyColor:      7254527,
			MoodSound: &BiomeMoodSound{
				Sound:             "minecraft:ambient.cave",
				TickDelay:         6000,
				BlockSearchExtent: 8,
				Offset:            2.0,
			},
			Music: &BiomeMusic{
				Sound:               "minecraft:music.overworld.badlands",
				MinDelay:            12000,
				MaxDelay:            24000,
				ReplaceCurrentMusic: false,
			},
		},
	}
)
