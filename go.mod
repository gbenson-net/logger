module gbenson.net/go/logger

go 1.23.0

require (
	github.com/rs/zerolog v1.33.0
	gotest.tools/v3 v3.5.2
)

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.31.0 // indirect
)

retract (
	v1.0.2 // Misversioned
	v0.0.1 // Test release
)
