# Introduction

![GitHub Release](https://img.shields.io/github/v/release/sdqri/effdsl)
[![GoDoc](https://pkg.go.dev/badge/github.com/sdqri/effdsl?status.svg)](https://pkg.go.dev/github.com/sdqri/effdsl?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdqri/effdsl)](https://goreportcard.com/report/github.com/sdqri/effdsl)
![GitHub License](https://img.shields.io/github/license/sdqri/effdsl)
<a href="https://github.com/sdqri/effdsl/pulls" style="text-decoration: none;">
    <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
</a>

**effdsl** provides a simple and functional way to build Elasticsearch queries in Go.Instead of relying on maps or, even worse, raw string literals to describe queries—which can be error-prone and lack features like easy parameterization, type safety, auto-completion, and compile-time validation— **effdsl** allows you to construct queries using intuitive function calls. This reduces the risk of subtle bugs caused by misspellings, makes parameterization easier and safer, and simplifies the process of constructing complex queries.

Moreover, its design makes the procedural creation of queries both straightforward and refined, making it particularly useful for cases where queries need to be generated programmatically. (This started because I needed to implement an interpreter of an internal DSL into Elasticsearch queries.) The module focuses solely on building the query body, without direct integration with the database, allowing seamless integration into an existing Go codebase.

**effdsl** supports most compound queries, full-text queries, and term-level queries. For a complete list of supported query types, please refer to the [API coverage](https://sdqri.github.com/effdsl/api_coverage) file in the **effdsl** GitHub repository. If there's a query type that isn't yet supported, feel free to open an issue or, even better, submit a pull request. 🙌

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

Start with `effdsl.Define()`, and use types and documentations to find suitable options.

### 🔍 Examples:

**Traditional Way:**

Here’s a simple match query in the traditional way using raw strings in Go:

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

**Using effdsl:**

And here’s the same query using effdsl:

```go
import (
    es "github.com/elastic/go-elasticsearch/v8"
	mq "github.com/sdqri/effdsl/queries/matchquery"
)

query, err := effdsl.Define(mq.MatchQuery("message", "Hello World"))

res, err := es.Search(
  es.Search.WithBody(strings.NewReader(query)),
)
```

For more examples and details on query parameters, visit the [documentation](https://sdqri.github.com/effdsl).

## 🤝 Contribution
Contributions are welcome! Whether it's fixing a bug 🐛, adding a new feature 🌟, or improving the documentation 📚, your help is appreciated. Please check out the CONTRIBUTING.md guide to get started.

## 📜 License
This project is licensed under the **MIT License**. For more details, see the [License](LICENSE.md) file. 📄 ( **In short:** You can use, modify, and distribute this software freely as long as you include the original copyright notice and license. The software is provided "as-is" without warranties or guarantees.)
