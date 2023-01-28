package effdsl

import (
	objs "github.com/sdqri/effdsl/objects"
)

//--------------------------------------------------------------------------------------//
//                                      functions                                       //
//--------------------------------------------------------------------------------------//

type BodyOption func(*objs.SearchBody) error

func Define(opts ...BodyOption) (body *objs.SearchBody, err error) {
	body = new(objs.SearchBody)
	for _, opt := range opts {
		err = opt(body)
		if err != nil {
			return nil, err
		}
	}
	return body, nil
}

var D objs.DefineType = objs.Define

//--------------------------------------------------------------------------------------//
//                                      constants                                       //
//--------------------------------------------------------------------------------------//

const (
	SORT_DEFAULT objs.SortOrder = objs.SORT_DEFAULT
	SORT_ASC     objs.SortOrder = objs.SORT_ASC
	SORT_DESC    objs.SortOrder = objs.SORT_DESC
)
