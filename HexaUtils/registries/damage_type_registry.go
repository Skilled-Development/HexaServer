package registries

import "HexaUtils/nbt"

type DamageTypeRegistry struct {
	Name                      string
	DamegeTypeRegistryEntries []DamageTypeRegistryEntry
}

func NewDamageTypeRegistry() *DamageTypeRegistry {
	return &DamageTypeRegistry{
		Name: "damage_type",
		DamegeTypeRegistryEntries: []DamageTypeRegistryEntry{
			Arrow,
			BadRespawnPoint,
			Cactus,
			Campfire,
			Cramming,
			DragonBreath,
			Drown,
			DryOut,
			Explosion,
			Fall,
			FallingAnvil,
			FallingBlock,
			FallingStalactite,
			Fireball,
			Fireworks,
			FlyIntoWall,
			Freeze,
			Generic,
			GenericKill,
			HotFloor,
			IndirectMagic,
			InFire,
			InWall,
			Lava,
			LightningBolt,
			Magic,
			MobAttack,
			MobAttackNoAggro,
			MobProjectile,
			OnFire,
			OutsideBorder,
			OutOfWorld,
			PlayerAttack,
			PlayerExplosion,
			SonicBoom,
			Spit,
			Stalagmite,
			Starve,
			Sting,
			SweetBerryBush,
			Thorns,
			Thrown,
			Trident,
			UnattributedFireball,
			WindCharge,
			Wither,
			WitherSkull,
		},
	}
}

func (d DamageTypeRegistry) GetEntries() []DamageTypeRegistryEntry {
	return d.DamegeTypeRegistryEntries
}

func (d DamageTypeRegistry) GetEntryByName(name string) *DamageTypeRegistryEntry {
	for _, entry := range d.DamegeTypeRegistryEntries {
		if entry.GetName() == name {
			return &entry
		}
	}
	return nil
}

func (d DamageTypeRegistry) GetEntriesAsNBTs() []nbt.Nbt {
	var nbts []nbt.Nbt
	for _, entry := range d.DamegeTypeRegistryEntries {
		nbts = append(nbts, entry.GetAsNBT())
	}
	return nbts
}

func (d DamageTypeRegistry) GetName() string {
	return d.Name
}

type DamageTypeRegistryEntry struct {
	BeautifiedName   string
	Name             string
	MessageID        string
	Scaling          string
	Exhaustion       float32
	DeathMessageType string
	Effects          string
}

func (d *DamageTypeRegistryEntry) GetMessageID() string {
	return d.MessageID
}

func (d *DamageTypeRegistryEntry) SetMessageID(MessageID string) {
	d.MessageID = MessageID
}

func (d *DamageTypeRegistryEntry) GetScaling() string {
	return d.Scaling
}

func (d *DamageTypeRegistryEntry) SetScaling(scaling string) {
	d.Scaling = scaling
}

func (d *DamageTypeRegistryEntry) GetExhaustion() float32 {
	return d.Exhaustion
}

func (d *DamageTypeRegistryEntry) SetExhaustion(exhaustion float32) {
	d.Exhaustion = exhaustion
}

func (d *DamageTypeRegistryEntry) GetDeathMessageType() string {
	return d.DeathMessageType
}

func (d *DamageTypeRegistryEntry) SetDeathMessageType(deathMessageType string) {
	d.DeathMessageType = deathMessageType
}

func (d *DamageTypeRegistryEntry) GetEffects() string {
	return d.Effects
}

func (d *DamageTypeRegistryEntry) SetEffects(effects string) {
	d.Effects = effects
}

func (d *DamageTypeRegistryEntry) GetName() string {
	return d.Name
}

func (d *DamageTypeRegistryEntry) SetName(name string) {
	d.Name = name
}

func (d *DamageTypeRegistryEntry) GetBeautifiedName() string {
	return d.BeautifiedName
}

func (d *DamageTypeRegistryEntry) SetBeautifiedName(beautifiedName string) {
	d.BeautifiedName = beautifiedName
}

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

var (
	Arrow = DamageTypeRegistryEntry{
		BeautifiedName: "Arrow",
		Name:           "arrow",
		Exhaustion:     0.1,
		MessageID:      "arrow",
		Scaling:        "when_caused_by_living_non_player",
	}

	BadRespawnPoint = DamageTypeRegistryEntry{
		BeautifiedName:   "BadRespawnPoint",
		Name:             "bad_respawn_point",
		DeathMessageType: "intentional_game_design",
		Exhaustion:       0.1,
		MessageID:        "badRespawnPoint",
		Scaling:          "always",
	}

	Cactus = DamageTypeRegistryEntry{
		BeautifiedName: "Cactus",
		Name:           "cactus",
		Exhaustion:     0.1,
		MessageID:      "cactus",
		Scaling:        "when_caused_by_living_non_player",
	}

	Campfire = DamageTypeRegistryEntry{
		BeautifiedName: "Campfire",
		Name:           "campfire",
		Effects:        "burning",
		Exhaustion:     0.1,
		MessageID:      "inFire",
		Scaling:        "when_caused_by_living_non_player",
	}

	Cramming = DamageTypeRegistryEntry{
		BeautifiedName: "Cramming",
		Name:           "cramming",
		Exhaustion:     0.0,
		MessageID:      "cramming",
		Scaling:        "when_caused_by_living_non_player",
	}

	DragonBreath = DamageTypeRegistryEntry{
		BeautifiedName: "DragonBreath",
		Name:           "dragon_breath",
		Exhaustion:     0.0,
		MessageID:      "dragonBreath",
		Scaling:        "when_caused_by_living_non_player",
	}

	Drown = DamageTypeRegistryEntry{
		BeautifiedName: "Drown",
		Name:           "drown",
		Effects:        "drowning",
		Exhaustion:     0.0,
		MessageID:      "drown",
		Scaling:        "when_caused_by_living_non_player",
	}

	DryOut = DamageTypeRegistryEntry{
		BeautifiedName: "DryOut",
		Name:           "dry_out",
		Exhaustion:     0.1,
		MessageID:      "dryout",
		Scaling:        "when_caused_by_living_non_player",
	}

	Explosion = DamageTypeRegistryEntry{
		BeautifiedName: "Explosion",
		Name:           "explosion",
		Exhaustion:     0.1,
		MessageID:      "explosion",
		Scaling:        "always",
	}

	Fall = DamageTypeRegistryEntry{
		BeautifiedName:   "Fall",
		Name:             "fall",
		DeathMessageType: "fall_variants",
		Exhaustion:       0.0,
		MessageID:        "fall",
		Scaling:          "when_caused_by_living_non_player",
	}

	FallingAnvil = DamageTypeRegistryEntry{
		BeautifiedName: "FallingAnvil",
		Name:           "falling_anvil",
		Exhaustion:     0.1,
		MessageID:      "anvil",
		Scaling:        "when_caused_by_living_non_player",
	}

	FallingBlock = DamageTypeRegistryEntry{
		BeautifiedName: "FallingBlock",
		Name:           "falling_block",
		Exhaustion:     0.1,
		MessageID:      "fallingBlock",
		Scaling:        "when_caused_by_living_non_player",
	}

	FallingStalactite = DamageTypeRegistryEntry{
		BeautifiedName: "FallingStalactite",
		Name:           "falling_stalactite",
		Exhaustion:     0.1,
		MessageID:      "fallingStalactite",
		Scaling:        "when_caused_by_living_non_player",
	}

	Fireball = DamageTypeRegistryEntry{
		BeautifiedName: "Fireball",
		Name:           "fireball",
		Effects:        "burning",
		Exhaustion:     0.1,
		MessageID:      "fireball",
		Scaling:        "when_caused_by_living_non_player",
	}

	Fireworks = DamageTypeRegistryEntry{
		BeautifiedName: "Fireworks",
		Name:           "fireworks",
		Exhaustion:     0.1,
		MessageID:      "fireworks",
		Scaling:        "when_caused_by_living_non_player",
	}

	FlyIntoWall = DamageTypeRegistryEntry{
		BeautifiedName: "FlyIntoWall",
		Name:           "fly_into_wall",
		Exhaustion:     0.0,
		MessageID:      "flyIntoWall",
		Scaling:        "when_caused_by_living_non_player",
	}

	Freeze = DamageTypeRegistryEntry{
		BeautifiedName: "Freeze",
		Name:           "freeze",
		Effects:        "freezing",
		Exhaustion:     0.0,
		MessageID:      "freeze",
		Scaling:        "when_caused_by_living_non_player",
	}

	Generic = DamageTypeRegistryEntry{
		BeautifiedName: "Generic",
		Name:           "generic",
		Exhaustion:     0.0,
		MessageID:      "generic",
		Scaling:        "when_caused_by_living_non_player",
	}

	GenericKill = DamageTypeRegistryEntry{
		BeautifiedName: "GenericKill",
		Name:           "generic_kill",
		Exhaustion:     0.0,
		MessageID:      "genericKill",
		Scaling:        "when_caused_by_living_non_player",
	}

	HotFloor = DamageTypeRegistryEntry{
		BeautifiedName: "HotFloor",
		Name:           "hot_floor",
		Effects:        "burning",
		Exhaustion:     0.1,
		MessageID:      "hotFloor",
		Scaling:        "when_caused_by_living_non_player",
	}

	IndirectMagic = DamageTypeRegistryEntry{
		BeautifiedName: "IndirectMagic",
		Name:           "indirect_magic",
		Exhaustion:     0.0,
		MessageID:      "indirectMagic",
		Scaling:        "when_caused_by_living_non_player",
	}

	InFire = DamageTypeRegistryEntry{
		BeautifiedName: "InFire",
		Name:           "in_fire",
		Effects:        "burning",
		Exhaustion:     0.1,
		MessageID:      "inFire",
		Scaling:        "when_caused_by_living_non_player",
	}

	InWall = DamageTypeRegistryEntry{
		BeautifiedName: "InWall",
		Name:           "in_wall",
		Exhaustion:     0.0,
		MessageID:      "inWall",
		Scaling:        "when_caused_by_living_non_player",
	}

	Lava = DamageTypeRegistryEntry{
		BeautifiedName: "Lava",
		Name:           "lava",
		Effects:        "burning",
		Exhaustion:     0.1,
		MessageID:      "lava",
		Scaling:        "when_caused_by_living_non_player",
	}

	LightningBolt = DamageTypeRegistryEntry{
		BeautifiedName: "LightningBolt",
		Name:           "lightning_bolt",
		Exhaustion:     0.1,
		MessageID:      "lightningBolt",
		Scaling:        "when_caused_by_living_non_player",
	}

	Magic = DamageTypeRegistryEntry{
		BeautifiedName: "Magic",
		Name:           "magic",
		Exhaustion:     0.0,
		MessageID:      "magic",
		Scaling:        "when_caused_by_living_non_player",
	}

	MobAttack = DamageTypeRegistryEntry{
		BeautifiedName: "MobAttack",
		Name:           "mob_attack",
		Exhaustion:     0.1,
		MessageID:      "mob",
		Scaling:        "when_caused_by_living_non_player",
	}

	MobAttackNoAggro = DamageTypeRegistryEntry{
		BeautifiedName: "MobAttackNoAggro",
		Name:           "mob_attack_no_aggro",
		Exhaustion:     0.1,
		MessageID:      "mob",
		Scaling:        "when_caused_by_living_non_player",
	}

	MobProjectile = DamageTypeRegistryEntry{
		BeautifiedName: "MobProjectile",
		Name:           "mob_projectile",
		Exhaustion:     0.1,
		MessageID:      "mob",
		Scaling:        "when_caused_by_living_non_player",
	}

	OnFire = DamageTypeRegistryEntry{
		BeautifiedName: "OnFire",
		Name:           "on_fire",
		Effects:        "burning",
		Exhaustion:     0.0,
		MessageID:      "onFire",
		Scaling:        "when_caused_by_living_non_player",
	}

	OutsideBorder = DamageTypeRegistryEntry{
		BeautifiedName: "OutsideBorder",
		Name:           "outside_border",
		Exhaustion:     0.0,
		MessageID:      "outsideBorder",
		Scaling:        "when_caused_by_living_non_player",
	}

	OutOfWorld = DamageTypeRegistryEntry{
		BeautifiedName: "OutOfWorld",
		Name:           "out_of_world",
		Exhaustion:     0.0,
		MessageID:      "outOfWorld",
		Scaling:        "when_caused_by_living_non_player",
	}

	PlayerAttack = DamageTypeRegistryEntry{
		BeautifiedName: "PlayerAttack",
		Name:           "player_attack",
		Exhaustion:     0.1,
		MessageID:      "player",
		Scaling:        "when_caused_by_living_non_player",
	}

	PlayerExplosion = DamageTypeRegistryEntry{
		BeautifiedName: "PlayerExplosion",
		Name:           "player_explosion",
		Exhaustion:     0.1,
		MessageID:      "explosion.player",
		Scaling:        "always",
	}

	SonicBoom = DamageTypeRegistryEntry{
		BeautifiedName: "SonicBoom",
		Name:           "sonic_boom",
		Exhaustion:     0.0,
		MessageID:      "sonic_boom",
		Scaling:        "always",
	}

	Spit = DamageTypeRegistryEntry{
		BeautifiedName: "Spit",
		Name:           "spit",
		Exhaustion:     0.1,
		MessageID:      "mob",
		Scaling:        "when_caused_by_living_non_player",
	}

	Stalagmite = DamageTypeRegistryEntry{
		BeautifiedName: "Stalagmite",
		Name:           "stalagmite",
		Exhaustion:     0.0,
		MessageID:      "stalagmite",
		Scaling:        "when_caused_by_living_non_player",
	}

	Starve = DamageTypeRegistryEntry{
		BeautifiedName: "Starve",
		Name:           "starve",
		Exhaustion:     0.0,
		MessageID:      "starve",
		Scaling:        "when_caused_by_living_non_player",
	}

	Sting = DamageTypeRegistryEntry{
		BeautifiedName: "Sting",
		Name:           "sting",
		Exhaustion:     0.1,
		MessageID:      "sting",
		Scaling:        "when_caused_by_living_non_player",
	}

	SweetBerryBush = DamageTypeRegistryEntry{
		BeautifiedName: "SweetBerryBush",
		Name:           "sweet_berry_bush",
		Effects:        "poking",
		Exhaustion:     0.1,
		MessageID:      "sweetBerryBush",
		Scaling:        "when_caused_by_living_non_player",
	}

	Thorns = DamageTypeRegistryEntry{
		BeautifiedName: "Thorns",
		Name:           "thorns",
		Effects:        "thorns",
		Exhaustion:     0.1,
		MessageID:      "thorns",
		Scaling:        "when_caused_by_living_non_player",
	}

	Thrown = DamageTypeRegistryEntry{
		BeautifiedName: "Thrown",
		Name:           "thrown",
		Exhaustion:     0.1,
		MessageID:      "thrown",
		Scaling:        "when_caused_by_living_non_player",
	}

	Trident = DamageTypeRegistryEntry{
		BeautifiedName: "Trident",
		Name:           "trident",
		Exhaustion:     0.1,
		MessageID:      "trident",
		Scaling:        "when_caused_by_living_non_player",
	}

	UnattributedFireball = DamageTypeRegistryEntry{
		BeautifiedName: "UnattributedFireball",
		Name:           "unattributed_fireball",
		Effects:        "burning",
		Exhaustion:     0.1,
		MessageID:      "onFire",
		Scaling:        "when_caused_by_living_non_player",
	}

	WindCharge = DamageTypeRegistryEntry{
		BeautifiedName: "WindCharge",
		Name:           "wind_charge",
		Exhaustion:     0.1,
		MessageID:      "mob",
		Scaling:        "when_caused_by_living_non_player",
	}

	Wither = DamageTypeRegistryEntry{
		BeautifiedName: "Wither",
		Name:           "wither",
		Exhaustion:     0.0,
		MessageID:      "wither",
		Scaling:        "when_caused_by_living_non_player",
	}

	WitherSkull = DamageTypeRegistryEntry{
		BeautifiedName: "WitherSkull",
		Name:           "wither_skull",
		Exhaustion:     0.1,
		MessageID:      "witherSkull",
		Scaling:        "when_caused_by_living_non_player",
	}
)
