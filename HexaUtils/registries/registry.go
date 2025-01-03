package registries

import "HexaUtils/nbt"

type Registry interface {
	GetName() string
	GetEntriesAsNBTs() []nbt.Nbt
}
