package utils

type KnownPack struct {
	Namespace string
	ID        string
	Version   string
}

func NewKnownPack(namespace string, id string, version string) *KnownPack {
	return &KnownPack{
		Namespace: namespace,
		ID:        id,
		Version:   version,
	}
}

func (p *KnownPack) GetNamespace() string {
	return p.Namespace
}

func (p *KnownPack) GetID() string {
	return p.ID
}

func (p *KnownPack) GetVersion() string {
	return p.Version
}

func (p *KnownPack) SetNamespace(namespace string) {
	p.Namespace = namespace
}

func (p *KnownPack) SetID(id string) {
	p.ID = id
}

func (p *KnownPack) SetVersion(version string) {
	p.Version = version
}
