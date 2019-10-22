# Go tutorial from tutorialedge.net

https://tutorialedge.net/golang/getting-started-with-go/

## Skipped

- Beginner/
  - Getting started
  - Basic types
  - Composite types (except maps)
  - Functions
  - Methods
  - Interfaces
  - Init
  - Go modules

## Testing

Testing command:

```sh
# move to this director
cd test_intro/

# Simple test
go test

# Verbose
go test -v

# Coverage display
go test --cover

# Display what's covered by testing
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

If not located in the proper package:

```sh
go test github.com/al-un/tutos-go/lang_tutorialedge/testingpractice
```

[Go cannot test a single file](https://stackoverflow.com/a/14723658/4906586)
