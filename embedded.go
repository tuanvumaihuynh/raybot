package raybot

import (
	_ "embed"
)

//go:embed bin/oas/openapi.yml
var OpenapiSpec []byte
