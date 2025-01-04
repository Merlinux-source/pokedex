package pokeapi

type Locations struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Region      NamedAPIResource   `json:"region"`
	Names       []string           `json:"names"`
	GameIndices []int              `json:"game_indices"`
	Areas       []NamedAPIResource `json:"areas"`
}
