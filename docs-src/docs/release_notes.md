---
hide:
  - navigation
---

# Release Notes

## v2.1.2 Latest
### What's New 
* ✨ Add support for boosting query, constant score query, disjunction max query in compound queries.
* ✨ Add support for match query, match_bool_prefix query, match_phrase query, match_phrase_prefix query, simple_query_string query in full text queries.
* ✨ Add support for ids query, prefix query, wildcard query in term-level queries.
* 📝 Add documentation for all supported queries.

### What's Changed 
* 🔄 Separate queries into packages to remove prefixed parameters.
* ✅ Complete functional options in queries.

## v1.2.0 

* 💡 Extend `MatchQuery` parameters & add `WildcardQuery` and `Suggesters`. PR [#2](https://github.com/sdqri/effdsl/pull/2) by [@moguchev](https://github.com/moguchev).

