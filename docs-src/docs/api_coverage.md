---
hide:
  - navigation
---

# API Coverage

## Compound queries

- [x] bool query
- [x] boosting query
- [x] constant score query 
- [x] disjunction max query
- [ ] function_score query

## Full text queries

- [ ] intervals query
- [x] match query
- [x] match_bool_prefix query
- [x] match_phrase query
- [x] match_phrase_prefix query
- [x] multi_match query
- [ ] combined_fields query
- [x] query_string query 
- [x] simple_query_string query

## Term-level queries

- [x] exists query
- [x] fuzzy query
- [x] ids query
- [x] prefix query
- [x] range query
- [x] regexp query
- [x] term query
- [x] terms query
- [x] terms_set query
- [x] wildcard query

## Aggregations

### Metrics

- [x] avg aggregation
- [x] boxplot aggregation
- [x] cartesian-bounds aggregation
- [x] cartesian-centroid aggregation
- [x] cardinality aggregation
- [x] extended stats aggregation
- [x] geo-bounds aggregation
- [x] geo-centroid aggregation
- [x] geo-line aggregation
- [x] sum aggregation

## Customize Search results (options)

- [x] Collapse search results
- [ ] Filter search results
- [ ] Highlighting
- [ ] Long-running searches
- [ ] Near real-time search
- [x] Paginate search results (Supported methods: Simple paginating, Search after)
- [ ] Retrieve inner hits
- [x] Retrieve selected fields (Supported methods: Source filtering)
- [ ] Search across clusters
- [ ] Search multiple data streams and indices
- [ ] Search shard routing
- [ ] Search templates
- [x] Sort search results (Suppoerted Parameters : value, order)
- [ ] kNN search

## Search APIs

- [x] Point in time
