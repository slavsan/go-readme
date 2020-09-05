# go-readme

A cli app for creating better project READMEs

### Usage
In order to generate a README.md file from a readme.yaml, run
```
go-readme create
```

### Setup

#### Clone
```
git clone git@github.com:slavsan/go-readme.git
```

#### Build
```
go build main.go
```

#### Start
```
go run main.go
```

#### Testing
To run the BDD (Cucumber) tests, execute
```
godog
```
If you want to run them with `progress` formatting (shorter), execute
```
godog --format=progress
```
In order to run a specific feature (test), you'll need to specify the feature file and line, like so
```
godog features/generate_markdown_readme.feature:197
```

#### Linting
```
golangci-lint run
```

### License
MIT

### Author
Slavcho Slavchev