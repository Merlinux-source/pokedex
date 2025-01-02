package pokeapi

const MapBaseUrl string = "https://pokeapi.co/api/v2/location-area/"

type MapLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type MapApiResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []MapLocation `json:"results"`
}
