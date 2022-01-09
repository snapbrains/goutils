This go project provides a set of utilities for working with Go

# Build:
go build


# Run all Tests:
gotest ./... -v

# Find bad stuff
go vet

# Run benchmarks:
go test -bench=. -benchmem


# Update packages
go get -u

# Build lib:
go mod tidy
go build
git push
