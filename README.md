# First error
```
# envoy-filter-bro
./config.go:17:38: cannot use &parser{} (value of type *parser) as api.StreamFilterConfigParser value in argument to http.RegisterHttpFilterConfigParser: *parser does not implement api.StreamFilterConfigParser (wrong type for method Parse)
		have Parse(*anypb.Any) (interface{}, error)
		want Parse(*anypb.Any) interface{

``` 

