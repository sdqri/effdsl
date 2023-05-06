package utils

import "encoding/json"

func CastStruct[S any, D any](source S) (destination D, err error) {
	// Convert struct to JSON string
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return
	}

	// Unmarshal JSON into new struct
	if err = json.Unmarshal(sourceJSON, &destination); err != nil {
		return
	}

	return
}
