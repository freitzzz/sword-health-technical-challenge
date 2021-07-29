package http

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type UserAuth struct {
	Identifier string `json:"id"`
	Secret     string `json:"secret"`
}
