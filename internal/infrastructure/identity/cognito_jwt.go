package identity

import (
	"context"
)

// Más adelante esto lo haremos compatible con application/auth.Claims
type Claims struct {
	Sub   string
	Email string
	Name  string
}

type CognitoJWTValidator struct {
	Issuer   string
	Audience string
	// Acá después vamos a cachear las keys (JWKS)
}

func NewCognitoJWTValidator(issuer, audience string) *CognitoJWTValidator {
	return &CognitoJWTValidator{
		Issuer:   issuer,
		Audience: audience,
	}
}

// ParseToken valida el JWT y devuelve los claims.
// Por ahora lo dejamos como stub.
func (v *CognitoJWTValidator) ParseToken(ctx context.Context, rawToken string) (*Claims, error) {
	// TODO: implementar validación real usando JWKS de Cognito
	return nil, nil
}
