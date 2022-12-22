module example.com/cyoa/v2

go 1.19

require (
	example.com/cyoa/decoder v0.0.0-00010101000000-000000000000
	github.com/kr/pretty v0.3.1
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
)

replace example.com/cyoa => ./cmd/cyoa/

replace example.com/cyoa/decoder => ./internal/json/
