package objects

import "encoding/json"

type PITS struct {
	ID        string `json:"id,omitempty"`
	KeepAlive string `json:"keep_alive,omitempty"`
}

func (pit PITS) MarshalJSON() ([]byte, error) {
	type PITAlias PITS
	return json.Marshal((PITAlias)(pit))
}

func PIT(id string, keepAlive string) PITS {
	pit := PITS{
		ID:        id,
		KeepAlive: keepAlive,
	}
	return pit
}
