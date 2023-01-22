package effdsl

import (
	objs "github.com/sdqri/effdsl/objects"
)

//--------------------------------------------------------------------------------------//
//                                      functions                                       //
//--------------------------------------------------------------------------------------//

var Define objs.DefineType = objs.Define
var D objs.DefineType = objs.Define

//--------------------------------------------------------------------------------------//
//                                      constants                                       //
//--------------------------------------------------------------------------------------//

const (
	SORT_DEFAULT objs.SortOrder = objs.SORT_DEFAULT
	SORT_ASC     objs.SortOrder = objs.SORT_ASC
	SORT_DESC    objs.SortOrder = objs.SORT_DESC
)
