---
hide:
  - navigation
---

# Release Notes

## v2.7.0 Latest
### What's New
* Add comprehensive aggregation support across metrics, bucket, and pipeline aggregations.
* Introduce aggregation APIs and helpers for composing aggregations in queries.

### Highlights
* Metrics aggregations: avg, sum, min/max, stats/extended stats, percentiles/percentile ranks, rate, scripted metric, t-test, string stats, top hits, top metrics, value count, weighted avg, boxplot, median absolute deviation, cardinality, matrix stats, cartesian/geo variants.
* Bucket aggregations: terms, filters/filter, date histogram/date range/auto date histogram, range/histogram, geo grids, missing, nested/reverse nested, composite, rare/multi terms, sampler/diversified/random sampler, significant terms/text, parent/children/global, adjacency matrix, variable width histogram.
* Pipeline aggregations: avg/sum/min/max/stats/extended stats/percentiles bucket, bucket script/selector/sort, cumulative sum/cardinality, derivative, serial diff, normalize, moving fn/percentiles, bucket correlation, bucket count K-S test, change point, inference bucket.

### What's Changed
* Expand aggregation example test coverage and add doc-style option cases.
* Update aggregation docs navigation and fix geo distance doc entry.

## v2.1.2
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
