package pokeapi

type Pokemon struct {
	Id                     int                `json:"id"`
	Name                   string             `json:"name"`
	BaseExperience         int                `json:"base_experience"`
	Height                 int                `json:"height"`
	IsDefault              bool               `json:"is_default"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	Abilities              []PokemonAbility   `json:"abilities"`
	Forms                  []NamedAPIResource `json:"forms"`
	GameIndicies           []VersionGameIndex `json:"game_indicies"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove      `json:"moves"`
	PastTypes              []PokemonTypePast  `json:"past_types"`
	Sprites                PokemonSprites     `json:"sprites"`
	Cries                  PokemonCries       `json:"cries"`
	Species                NamedAPIResource   `json:"species"`
	Stats                  []PokemonStat      `json:"stats"`
	Types                  []PokemonType      `json:"types"`
}
