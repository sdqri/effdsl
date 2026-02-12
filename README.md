# effdsl

![GitHub Release](https://img.shields.io/github/v/release/sdqri/effdsl)
[![GoDoc](https://pkg.go.dev/badge/github.com/sdqri/effdsl/v2?status.svg)](https://pkg.go.dev/github.com/sdqri/effdsl/v2?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdqri/effdsl)](https://goreportcard.com/report/github.com/sdqri/effdsl)
![GitHub License](https://img.shields.io/github/license/sdqri/effdsl)
<a href="https://github.com/sdqri/effdsl/pulls" style="text-decoration: none;"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome"></a>
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

effdsl is a composable Go DSL for building Elasticsearch queries and aggregations with type-safe options and predictable JSON output.

## Key Features

- Type-safe query construction without raw JSON maps
- Functional options API for readable, composable queries
- Broad aggregation coverage (metrics, bucket, pipeline)

For details and examples, see the [documentation](https://sdqri.github.io/effdsl).

## Getting started

### Install

With [Go module](https://github.com/golang/go/wiki/Modules) support, add the import and Go will fetch dependencies automatically:

```go
import "github.com/sdqri/effdsl/v2"
```

Or install directly:

```sh
go get -u github.com/sdqri/effdsl/v2
```

### How to use

Start with `effdsl.Define()` and compose queries with options.

### Examples

**Raw JSON example**

Example match query using a raw JSON string:

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
)

query := `{
  "query": {
    "match": {
      "message": {
        "query": "Hello World"
      }
    }
  }
}`

res, err := es.Search(
  es.Search.WithBody(strings.NewReader(query)),
)
```

**Using effdsl**

The same query using effdsl:

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
    
    "github.com/sdqri/effdsl/v2"
    mq "github.com/sdqri/effdsl/v2/queries/matchquery"
)

query, err := effdsl.Define(
    effdsl.WithQuery(
        mq.MatchQuery("message", "Hello World"),
    ),
)

res, err := es.Search(
  es.Search.WithBody(strings.NewReader(query)),
)
```

For more examples and details on query parameters, visit the [documentation](https://sdqri.github.io/effdsl).

## Contributing
Contributions are welcome. Thanks for helping improve effdsl. 🤝 Please see [CONTRIBUTING.md](CONTRIBUTING.md) to get started.

## License
This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md).
