package pokeapi

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}
