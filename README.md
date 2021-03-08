[![codecov](https://codecov.io/gh/tufin/oasdiff/branch/master/graph/badge.svg?token=Y8BM6X77JY)](https://codecov.io/gh/tufin/oasdiff)
[![CircleCI](https://circleci.com/gh/Tufin/oasdiff.svg?style=svg)](https://circleci.com/gh/Tufin/oasdiff)
[![Go Report Card](https://goreportcard.com/badge/github.com/tufin/oasdiff)](https://goreportcard.com/report/github.com/tufin/oasdiff)
[![GoDoc](https://godoc.org/github.com/tufin/oasdiff?status.svg)](https://godoc.org/github.com/tufin/oasdiff)

# OpenAPI Diff Go Module
This [Go](https://golang.org) module provides a diff utility for [OpenAPI Spec 3](https://swagger.io/specification/).

## Unique features vs. other diff tools
- go module
- deep diff into paths, parameters, requests, responses, schemas, enums etc.

## Build
```
git clone https://github.com/Tufin/oasdiff.git
cd oasdiff
go build
```

## Running from the command-line
```
./oasdiff -base data/openapi-test1.yaml -revision data/openapi-test2.yaml
```

## Help
```
./oasdiff --help
```

## Embedding into your Go program
```
diff.Get(&diff.Config{}, spec1, spec2)
```
See full example: [main.go](main.go)

## Notes
1. oasdiff expects [OpenAPI References](https://swagger.io/docs/specification/using-ref/) to be resolved.  
You can resolve refs using [this function](https://pkg.go.dev/github.com/getkin/kin-openapi/openapi3#SwaggerLoader.ResolveRefsIn) from the openapi3 package.

2. oasdiff ignores changes to [Examples](https://swagger.io/specification/#example-object) and [Extensions](https://swagger.io/specification/#specification-extensions) by default. You can change this behavior through [configuration](diff/config.go).

## Work in progress
While most objects of OpenAPI Spec are already supported by this diff tool, some are still missing, notably: Examples, ExternalDocs, Security, Links, Variables and a couple more.  
Pull requests are welcome!


