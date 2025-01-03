package config

import "HexaUtils/registries"

var RegistriesManagerInstance *RegistriesManager

type RegistriesManager struct {
	Registries []registries.Registry
}

func NewRegistriesManager() *RegistriesManager {
	registryManager := &RegistriesManager{}
	registryManager.Registries = append(registryManager.Registries, registries.NewDamageTypeRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewArmorTrimRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewArmorTrimPatternRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewBannerPatternRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewBiomeRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewPaitingVariantRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewWolfVariantRegistry())
	registryManager.Registries = append(registryManager.Registries, registries.NewDimensionTypeRegistry())
	RegistriesManagerInstance = registryManager
	return registryManager
}
