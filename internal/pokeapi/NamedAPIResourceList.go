package pokeapi

type NamedAPIResourceList struct {
	Count    int
	Next     string
	Previous string
	Results  []NamedAPIResource
}
