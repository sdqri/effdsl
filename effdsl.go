package effdsl

import (
	objs "github.com/sdqri/effdsl/objects"
)

//--------------------------------------------------------------------------------------//
//                                    Type aliasing                                     //
//--------------------------------------------------------------------------------------//

type (
	M                  = objs.M
	SearchBody         = objs.SearchBody
	BodyOption         = objs.BodyOption
	QueryResult        = objs.QueryResult
	SortClauseResult   = objs.SortClauseResult
	SourceFitlerOption = objs.SourceFitlerOption
	D                  = objs.DefineType
)

//--------------------------------------------------------------------------------------//
//                                        Define                                        //
//--------------------------------------------------------------------------------------//

var (
	Define = objs.Define
)

//--------------------------------------------------------------------------------------//
//                                      constants                                       //
//--------------------------------------------------------------------------------------//

const (
	SORT_DEFAULT objs.SortOrder = objs.SORT_DEFAULT
	SORT_ASC     objs.SortOrder = objs.SORT_ASC
	SORT_DESC    objs.SortOrder = objs.SORT_DESC
)
