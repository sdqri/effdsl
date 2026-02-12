---
hide:
  - navigation
---

# API Coverage

This page tracks the current coverage of effdsl APIs. Status indicates whether a feature is implemented in this repository.

## Queries

### Compound queries

| Item | Status | Docs |
| --- | --- | --- |
| Bool query | Supported | [bool_query.md](bool_query.md) |
| Boosting query | Supported | [boosting_query.md](boosting_query.md) |
| Constant score query | Supported | [constant_score.md](constant_score.md) |
| Disjunction max query | Supported | [dis_max_query.md](dis_max_query.md) |
| Function score query | Not yet | — |

### Full text queries

| Item | Status | Docs |
| --- | --- | --- |
| Intervals query | Not yet | — |
| Match query | Supported | [match_query.md](match_query.md) |
| Match bool prefix query | Supported | [match_bool_prefix.md](match_bool_prefix.md) |
| Match phrase query | Supported | [match_phrase_query.md](match_phrase_query.md) |
| Match phrase prefix query | Supported | [match_phrase_prefix.md](match_phrase_prefix.md) |
| Multi match query | Supported | [multi_match_query.md](multi_match_query.md) |
| Combined fields query | Not yet | — |
| Query string query | Supported | [query_string.md](query_string.md) |
| Simple query string query | Supported | [simple_query_string.md](simple_query_string.md) |

### Term-level queries

| Item | Status | Docs |
| --- | --- | --- |
| Exists query | Supported | [exists_query.md](exists_query.md) |
| Fuzzy query | Supported | [fuzzy_query.md](fuzzy_query.md) |
| IDs query | Supported | [ids_query.md](ids_query.md) |
| Prefix query | Supported | [prefix_query.md](prefix_query.md) |
| Range query | Supported | [range_query.md](range_query.md) |
| Regexp query | Supported | [regexp_query.md](regexp_query.md) |
| Term query | Supported | [term_query.md](term_query.md) |
| Terms query | Supported | [terms_query.md](terms_query.md) |
| Terms set query | Supported | [terms_set_query.md](terms_set_query.md) |
| Wildcard query | Supported | [wildcard_query.md](wildcard_query.md) |

### Joining queries

| Item | Status | Docs |
| --- | --- | --- |
| Nested query | Supported | [nested_query.md](nested_query.md) |

### Suggesters

| Item | Status | Docs |
| --- | --- | --- |
| Completion suggester | Supported | [completion_suggester.md](completion_suggester.md) |
| Phrase suggester | Supported | [phrase_suggester.md](phrase_suggester.md) |
| Term suggester | Supported | [term_suggester.md](term_suggester.md) |

## Aggregations

### Metrics

| Item | Status | Docs |
| --- | --- | --- |
| Avg aggregation | Supported | [avg_aggregation.md](avg_aggregation.md) |
| Boxplot aggregation | Supported | [boxplot_aggregation.md](boxplot_aggregation.md) |
| Cartesian bounds aggregation | Supported | [cartesian_bounds_aggregation.md](cartesian_bounds_aggregation.md) |
| Cartesian centroid aggregation | Supported | [cartesian_centroid_aggregation.md](cartesian_centroid_aggregation.md) |
| Cardinality aggregation | Supported | [cardinality_aggregation.md](cardinality_aggregation.md) |
| Extended stats aggregation | Supported | [extended_stats_aggregation.md](extended_stats_aggregation.md) |
| Geo bounds aggregation | Supported | [geo_bounds_aggregation.md](geo_bounds_aggregation.md) |
| Geo centroid aggregation | Supported | [geo_centroid_aggregation.md](geo_centroid_aggregation.md) |
| Geo line aggregation | Supported | [geo_line_aggregation.md](geo_line_aggregation.md) |
| Matrix stats aggregation | Supported | [matrix_stats_aggregation.md](matrix_stats_aggregation.md) |
| Max aggregation | Supported | [max_aggregation.md](max_aggregation.md) |
| Median absolute deviation aggregation | Supported | [median_absolute_deviation_aggregation.md](median_absolute_deviation_aggregation.md) |
| Min aggregation | Supported | [min_aggregation.md](min_aggregation.md) |
| Percentile ranks aggregation | Supported | [percentile_ranks_aggregation.md](percentile_ranks_aggregation.md) |
| Percentiles aggregation | Supported | [percentiles_aggregation.md](percentiles_aggregation.md) |
| Rate aggregation | Supported | [rate_aggregation.md](rate_aggregation.md) |
| Scripted metric aggregation | Supported | [scripted_metric_aggregation.md](scripted_metric_aggregation.md) |
| Stats aggregation | Supported | [stats_aggregation.md](stats_aggregation.md) |
| String stats aggregation | Supported | [string_stats_aggregation.md](string_stats_aggregation.md) |
| Sum aggregation | Supported | [sum_aggregation.md](sum_aggregation.md) |
| T-test aggregation | Supported | [t_test_aggregation.md](t_test_aggregation.md) |
| Top hits aggregation | Supported | [top_hits_aggregation.md](top_hits_aggregation.md) |
| Top metrics aggregation | Supported | [top_metrics_aggregation.md](top_metrics_aggregation.md) |
| Value count aggregation | Supported | [value_count_aggregation.md](value_count_aggregation.md) |
| Weighted avg aggregation | Supported | [weighted_avg_aggregation.md](weighted_avg_aggregation.md) |

### Bucket

| Item | Status | Docs |
| --- | --- | --- |
| Adjacency matrix aggregation | Supported | [adjacency_matrix_aggregation.md](adjacency_matrix_aggregation.md) |
| Auto date histogram aggregation | Supported | [auto_date_histogram_aggregation.md](auto_date_histogram_aggregation.md) |
| Composite aggregation | Supported | [composite_aggregation.md](composite_aggregation.md) |
| Date histogram aggregation | Supported | [date_histogram_aggregation.md](date_histogram_aggregation.md) |
| Date range aggregation | Supported | [date_range_aggregation.md](date_range_aggregation.md) |
| Filter aggregation | Supported | [filter_aggregation.md](filter_aggregation.md) |
| Filters aggregation | Supported | [filters_aggregation.md](filters_aggregation.md) |
| Geo distance aggregation | Supported | [geo_distance_aggregation.md](geo_distance_aggregation.md) |
| Geohash grid aggregation | Supported | [geohash_grid_aggregation.md](geohash_grid_aggregation.md) |
| Geo-tile grid aggregation | Supported | [geo_tile_grid_aggregation.md](geo_tile_grid_aggregation.md) |
| Geohex grid aggregation | Supported | [geohex_grid_aggregation.md](geohex_grid_aggregation.md) |
| Global aggregation | Supported | [global_aggregation.md](global_aggregation.md) |
| Histogram aggregation | Supported | [histogram_aggregation.md](histogram_aggregation.md) |
| IP range aggregation | Supported | [ip_range_aggregation.md](ip_range_aggregation.md) |
| Missing aggregation | Supported | [missing_aggregation.md](missing_aggregation.md) |
| Multi terms aggregation | Supported | [multi_terms_aggregation.md](multi_terms_aggregation.md) |
| Nested aggregation | Supported | [nested_aggregation.md](nested_aggregation.md) |
| Parent aggregation | Supported | [parent_aggregation.md](parent_aggregation.md) |
| Range aggregation | Supported | [range_aggregation.md](range_aggregation.md) |
| Rare terms aggregation | Supported | [rare_terms_aggregation.md](rare_terms_aggregation.md) |
| Reverse nested aggregation | Supported | [reverse_nested_aggregation.md](reverse_nested_aggregation.md) |
| Sampler aggregation | Supported | [sampler_aggregation.md](sampler_aggregation.md) |
| Diversified sampler aggregation | Supported | [diversified_sampler_aggregation.md](diversified_sampler_aggregation.md) |
| Random sampler aggregation | Supported | [random_sampler_aggregation.md](random_sampler_aggregation.md) |
| Significant terms aggregation | Supported | [significant_terms_aggregation.md](significant_terms_aggregation.md) |
| Significant text aggregation | Supported | [significant_text_aggregation.md](significant_text_aggregation.md) |
| Terms aggregation | Supported | [terms_aggregation.md](terms_aggregation.md) |
| Children aggregation | Supported | [children_aggregation.md](children_aggregation.md) |
| Variable width histogram aggregation | Supported | [variable_width_histogram_aggregation.md](variable_width_histogram_aggregation.md) |

### Pipeline

| Item | Status | Docs |
| --- | --- | --- |
| Avg bucket aggregation | Supported | [avg_bucket_aggregation.md](avg_bucket_aggregation.md) |
| Bucket script aggregation | Supported | [bucket_script_aggregation.md](bucket_script_aggregation.md) |
| Bucket selector aggregation | Supported | [bucket_selector_aggregation.md](bucket_selector_aggregation.md) |
| Bucket sort aggregation | Supported | [bucket_sort_aggregation.md](bucket_sort_aggregation.md) |
| Cumulative cardinality aggregation | Supported | [cumulative_cardinality_aggregation.md](cumulative_cardinality_aggregation.md) |
| Cumulative sum aggregation | Supported | [cumulative_sum_aggregation.md](cumulative_sum_aggregation.md) |
| Derivative aggregation | Supported | [derivative_aggregation.md](derivative_aggregation.md) |
| Extended stats bucket aggregation | Supported | [extended_stats_bucket_aggregation.md](extended_stats_bucket_aggregation.md) |
| Max bucket aggregation | Supported | [max_bucket_aggregation.md](max_bucket_aggregation.md) |
| Min bucket aggregation | Supported | [min_bucket_aggregation.md](min_bucket_aggregation.md) |
| Stats bucket aggregation | Supported | [stats_bucket_aggregation.md](stats_bucket_aggregation.md) |
| Sum bucket aggregation | Supported | [sum_bucket_aggregation.md](sum_bucket_aggregation.md) |
| Percentiles bucket aggregation | Supported | [percentiles_bucket_aggregation.md](percentiles_bucket_aggregation.md) |
| Serial diff aggregation | Supported | [serial_diff_aggregation.md](serial_diff_aggregation.md) |
| Normalize aggregation | Supported | [normalize_aggregation.md](normalize_aggregation.md) |
| Moving function aggregation | Supported | [moving_fn_aggregation.md](moving_fn_aggregation.md) |
| Moving percentiles aggregation | Supported | [moving_percentiles_aggregation.md](moving_percentiles_aggregation.md) |
| Bucket count K-S test aggregation | Supported | [bucket_count_ks_test_aggregation.md](bucket_count_ks_test_aggregation.md) |
| Bucket correlation aggregation | Supported | [bucket_correlation_aggregation.md](bucket_correlation_aggregation.md) |
| Change point aggregation | Supported | [change_point_aggregation.md](change_point_aggregation.md) |
| Inference bucket aggregation | Supported | [inference_bucket_aggregation.md](inference_bucket_aggregation.md) |

## Search

### Customize search results

| Item | Status | Docs |
| --- | --- | --- |
| Collapse search results | Supported | — |
| Filter search results | Not yet | — |
| Highlighting | Not yet | — |
| Long-running searches | Not yet | — |
| Near real-time search | Not yet | — |
| Paginate search results | Supported | [paginate_search.md](paginate_search.md) |
| Retrieve inner hits | Not yet | — |
| Retrieve selected fields (source filtering) | Supported | — |
| Search across clusters | Not yet | — |
| Search multiple data streams and indices | Not yet | — |
| Search shard routing | Not yet | — |
| Search templates | Not yet | — |
| Sort search results | Supported | [sort_search.md](sort_search.md) |
| kNN search | Not yet | — |

## Search APIs

| Item | Status | Docs |
| --- | --- | --- |
| Point in time | Supported | — |
