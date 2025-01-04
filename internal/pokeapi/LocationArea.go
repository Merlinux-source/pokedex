package pokeapi

type LocationArea struct {
	Id                    int                   `json:"id"`
	Name                  string                `json:"name"`
	GameIndex             int                   `json:"game_index"`
	EncounterMethodeRates []EncounterMethodRate `json:"encounter_methode_rates"`
	Location              []NamedAPIResource    `json:"location"`
	Names                 []Name                `json:"names"`
	PokemonEncounters     []PokemonEncounter    `json:"pokemon_encounters"`
}
