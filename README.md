# effdsl

[![GoDoc](https://pkg.go.dev/badge/github.com/sdqri/effdsl?status.svg)](https://pkg.go.dev/github.com/sdqri/effdsl?tab=doc)

Functional Elasticsearch DSL for Golang!

What effdsl does:

- Allows for constructing Elasticsearch queries using function calls instead of JSON objects.
- Supports a variety of query types and options.
- Simplifies the process of constructing complex queries.
- Provides type safety and compile-time checking of queries.
- Is easy to use and integrate with existing Go codebases.

Note: For the complete list of supported query types and options for each one, please refer to the [here](FEATURES.md) file in the effdsl GitHub repository.

## Getting started

### Getting effdsl

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/sdqri/effdsl"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `effdsl` package:

```sh
$ go get -u github.com/sdqri/effdsl
```

### How to use

Start with `effdsl.Define()`, and use `effdsl.?` to find suitable options.

### Examples:

An example query :

```go
import "github.com/sdqri/effdsl"

body, err := effdsl.Define(
	effdsl.WithSourceFilter(
		effdsl.WithIncludes("field1", "field2"),
		effdsl.WithExcludes("field3", "field4"),
	),
	effdsl.WithPaginate(2, 100),
	effdsl.WithQuery(
		effdsl.BoolQuery(
			effdsl.Must(
				effdsl.QueryString("value1", effdsl.WithFields("title", "content")),
			),
			effdsl.Filter(
				effdsl.RangeQuery("published_at", effdsl.WithGT("now-24h")),
				effdsl.TermQuery("field5.keyword", "value2"),
				effdsl.ExistsQuery("field6"),
			),
			effdsl.MustNot(
				effdsl.QueryString("value3", effdsl.WithFields("title", "content")),
			),
			effdsl.Should(
				effdsl.QueryString("value4", effdsl.WithFields("title", "content")),
			),
		),
	),
	effdsl.WithSort(
		effdsl.SortClause("field1", effdsl.SORT_DESC),
		effdsl.SortClause("_score", effdsl.SORT_DEFAULT),
	),
	effdsl.WithCollpse("field7"),
)
```

The code above constructs the query below :

```
{
   "_source":{
      "includes":[
         "field1",
         "field2"
      ],
      "excludes":[
         "field3",
         "field4"
      ]
   },
   "from":2,
   "size":100,
   "query":{
      "bool":{
         "must":[
            {
               "query_string":{
                  "query":"value1",
                  "fields":[
                     "title",
                     "content"
                  ]
               }
            }
         ],
         "filter":[
            {
               "range":{
                  "published_at":{
                     "gt":"now-24h"
                  }
               }
            },
            {
               "term":{
                  "field5.keyword":{
                     "value":"value2"
                  }
               }
            },
            {
               "exists":{
                  "field":"field6"
               }
            }
         ],
         "must_not":[
            {
               "query_string":{
                  "query":"value3",
                  "fields":[
                     "title",
                     "content"
                  ]
               }
            }
         ],
         "should":[
            {
               "query_string":{
                  "query":"value4",
                  "fields":[
                     "title",
                     "content"
                  ]
               }
            }
         ]
      }
   },
   "sort":[
      {
         "field1":"desc"
      },
      "_score"
   ],
   "collapse":{
      "field":"field7"
   }
}
```

## Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](issues).
