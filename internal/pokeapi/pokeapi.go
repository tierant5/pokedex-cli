package pokeapi

const (
	BaseUrl = "https://pokeapi.co/api/v2/"
)

type PokeAPIType interface {
	getApiResource() string
}
