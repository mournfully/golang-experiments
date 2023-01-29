module example.com/urlshort/v2

go 1.19

require example.com/handlers v0.0.0-00010101000000-000000000000

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace example.com/handlers => ./handlers
