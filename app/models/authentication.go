package models

type Response struct {
	Message string `json:"message"`
}

// Jwks stores a slice of JSON Web Keys.
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// These keys contain the public keys, which will be used to verify JWTs.
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}
