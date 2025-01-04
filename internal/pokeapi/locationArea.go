package pokeapi

const APIV2_LocationAreaBaseURL string = "https://pokeapi.co/api/v2/location-area/"
const APIV2_PokemonBaseURL string = "https://pokeapi.co/api/v2/pokemon/"

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedAPIResourceList struct {
	Count    int
	Next     string
	Previous string
	Results  []NamedAPIResource
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}

type Locations struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Region      NamedAPIResource   `json:"region"`
	Names       []string           `json:"names"`
	GameIndices []int              `json:"game_indices"`
	Areas       []NamedAPIResource `json:"areas"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          []NamedAPIResource `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type LocationArea struct {
	Id                    int                   `json:"id"`
	Name                  string                `json:"name"`
	GameIndex             int                   `json:"game_index"`
	EncounterMethodeRates []EncounterMethodRate `json:"encounter_methode_rates"`
	Location              []NamedAPIResource    `json:"location"`
	Names                 []Name                `json:"names"`
	PokemonEncounters     []PokemonEncounter    `json:"pokemon_encounters"`
}
