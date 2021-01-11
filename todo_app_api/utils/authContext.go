package utils

type contextKey struct {
	name string
}

// TokenClaimsKey is the key that stores the access token claims
var TokenClaimsKey = &contextKey{"token"}

// TokenClaims defines the access token claims
type TokenClaims struct {
	UserID         int
	RefreshTokenID int
}
