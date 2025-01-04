package pokeapi

const MapBaseUrl string = "https://pokeapi.co/api/v2/location-area/"

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

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}

type LocationArea struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Region      NamedAPIResource `json:"region"`
	Names       []string         `json:"names"`
	GameIndices []int            `json:"game_indices"`
}
