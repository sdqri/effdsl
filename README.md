# effdsl

![GitHub Release](https://img.shields.io/github/v/release/sdqri/effdsl)
[![GoDoc](https://pkg.go.dev/badge/github.com/sdqri/effdsl?status.svg)](https://pkg.go.dev/github.com/sdqri/effdsl?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/sdqri/effdsl)](https://goreportcard.com/report/github.com/sdqri/effdsl)
![GitHub License](https://img.shields.io/github/license/sdqri/effdsl)
<a href="https://github.com/sdqri/effdsl/pulls" style="text-decoration: none;"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome"></a>
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

This module provides a simple and functional way to build Elasticsearch queries in Go.

### 🚀 Key Features

- **Type-safe query construction:** 🛡️ Avoids error-prone maps and raw string literals by using intuitive function calls, enhancing type safety, auto-completion, and compile-time validation.
- **Procedural query creation:** ✨ Designed for straightforward and refined query building, particularly useful when queries need to be generated programmatically.
- **Comprehensive query support:** 📚 Covers most compound, full-text, and term-level queries, with easy extension for additional types.

For more information, detailed guides, and examples, please read the [documentation](https://sdqri.github.io/effdsl).

## Getting started

### Getting effdsl

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/sdqri/effdsl/v2"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `effdsl` package:

```sh
$ go get -u github.com/sdqri/effdsl/v2
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

## 🤝 Contribution
Contributions are welcome! Whether it's fixing a bug 🐛, adding a new feature 🌟, or improving the documentation 📚, your help is appreciated. Please check out the CONTRIBUTING.md guide to get started.

## 📜 License
This project is licensed under the **MIT License**. For more details, see the [License](LICENSE.md) file. 📄 ( **In short:** You can use, modify, and distribute this software freely as long as you include the original copyright notice and license. The software is provided "as-is" without warranties or guarantees.)
