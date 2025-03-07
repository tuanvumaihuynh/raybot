package raybot

import (
	_ "embed"
)

//go:embed api/openapi/gen/openapi.yml
var OpenapiSpec []byte
