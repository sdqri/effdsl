package objects

import "encoding/json"

type SearchAfterS []any

func (sa SearchAfterS) SearchAfterInfo() string {
	return "Search After"
}

func (sa SearchAfterS) MarshalJSON() ([]byte, error) {
	type SearchAfterAlias SearchAfterS
	return json.Marshal((SearchAfterAlias)(sa))
}

func SearchAfter(sortValues ...any) SearchAfterResult {
	searchAfter := SearchAfterS(sortValues)
	return SearchAfterResult{
		Ok:  searchAfter,
		Err: nil,
	}
}
