package pokeapi

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}
