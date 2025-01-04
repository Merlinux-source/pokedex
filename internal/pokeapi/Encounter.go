package pokeapi

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          []NamedAPIResource `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}
