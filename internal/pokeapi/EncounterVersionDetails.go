package pokeapi

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}
