This go project provides the access to dynamodb for the snap projects

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
