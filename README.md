[![codecov](https://codecov.io/gh/Tufin/oasdiff/branch/master/graph/badge.svg?token=Y8BM6X77JY)](https://codecov.io/gh/Tufin/oasdiff)
[![CircleCI](https://circleci.com/gh/Tufin/oasdiff.svg?style=svg)](https://circleci.com/gh/Tufin/oasdiff)

# OpenAPI Spec Diff
A diff tool for OpenAPI Spec 3.  

## Unique features vs. other OAS3 diff tools
- go module
- deep diff into paths, parameters, responses, schemas, enums etc.

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
package main

import (
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/load"
)

func main() {
	swaggerLoader := openapi3.NewSwaggerLoader()
	swaggerLoader.IsExternalRefsAllowed = true

	loader := load.NewOASLoader(swaggerLoader)

	base, err := loader.From("https://raw.githubusercontent.com/Tufin/oasdiff/master/data/openapi-test1.yaml")
	if err != nil {
		return
	}

	revision, err := loader.From("https://raw.githubusercontent.com/Tufin/oasdiff/master/data/openapi-test2.yaml")
	if err != nil {
		return
	}

	bytes, err := json.MarshalIndent(diff.Get(base, revision, "", ""), "", " ")
	if err != nil {
		log.Errorf("failed to marshal result with '%v'", err)
		return
	}

	fmt.Printf("%s\n", bytes)
}
```
