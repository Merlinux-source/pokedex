package pokeapi

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}
