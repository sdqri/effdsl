package effdsl

import "encoding/json"

type PITS struct {
	ID        string `json:"id,omitempty"`
	KeepAlive string `json:"keep_alive,omitempty"`
}

func (pit PITS) MarshalJSON() ([]byte, error) {
	type PITBase PITS
	return json.Marshal((PITBase)(pit))
}

func PIT(id string, keepAlive string) PITS {
	pit := PITS{
		ID:        id,
		KeepAlive: keepAlive,
	}
	return pit
}
