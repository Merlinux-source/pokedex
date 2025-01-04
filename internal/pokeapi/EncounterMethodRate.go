package pokeapi

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}
