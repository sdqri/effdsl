# effdsl

Functional Elasticsearch DSL for Golang!

Check supported features [here](FEATURES.md).

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

Start with `effdsl.Define()`, and use `effdsl.D.?` to find suitable options.

### Examples:

An example query :

```go
import "github.com/sdqri/effdsl"

body, err := effdsl.Define(
	effdsl.D.WithSourceFilter(
		effdsl.D.WithIncludes("field1", "field2"),
		effdsl.D.WithExcludes("field3", "field4"),
	),
	effdsl.D.WithQuery(
		effdsl.D.BoolQuery(
			effdsl.D.Must(
				effdsl.D.QueryString("value1", effdsl.D.WithFields("title", "content")),
			),
			effdsl.D.Filter(
				effdsl.D.RangeQuery("published_at", effdsl.D.WithGT("now-24h")),
				effdsl.D.TermQuery("field5.keyword", "value2"),
				effdsl.D.ExistsQuery("field6"),
			),
			effdsl.D.MustNot(
				effdsl.D.QueryString("value3", effdsl.D.WithFields("title", "content")),
			),
			effdsl.D.Should(
				effdsl.D.QueryString("value4", effdsl.D.WithFields("title", "content")),
			),
		),
	),
	effdsl.D.WithSort(
		effdsl.D.SortClause("field1", effdsl.SORT_DESC),
		effdsl.D.SortClause("_score", effdsl.SORT_DEFAULT),
	),
	effdsl.D.WithCollpse("field7"),
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
