package pokeapi

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}
