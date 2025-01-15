package world

type WorldType int

const (
	Overworld WorldType = 0
	Nether    WorldType = 1
	End       WorldType = 2
)

func (w WorldType) String() string {
	switch w {
	case Overworld:
		return "Overworld"
	case Nether:
		return "Nether"
	case End:
		return "End"
	}
	return "Unknown"
}

func (w WorldType) Int() int {
	return int(w)
}

func (w WorldType) Uint() uint {
	return uint(w)
}

func (w WorldType) Int32() int32 {
	return int32(w)
}
