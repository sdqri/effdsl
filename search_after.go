package effdsl

import "encoding/json"

type SearchAfterS []any

func (sa SearchAfterS) SearchAfterInfo() string {
	return "Search After"
}

func (sa SearchAfterS) MarshalJSON() ([]byte, error) {
	type SearchAfterBase SearchAfterS
	return json.Marshal((SearchAfterBase)(sa))
}

func SearchAfter(sortValues ...any) SearchAfterResult {
	searchAfter := SearchAfterS(sortValues)
	return SearchAfterResult{
		Ok:  searchAfter,
		Err: nil,
	}
}
