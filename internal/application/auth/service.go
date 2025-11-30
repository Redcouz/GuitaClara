package auth

import (
	"context"
	"time"

	"guitaclara/internal/domain/user"
)

type Claims struct {
	Sub   string
	Email string
	Name  string
}

// UserRepository define lo que necesitamos de la capa de persistencia
type UserRepository interface {
	GetByGoogleSub(ctx context.Context, sub string) (*user.User, error)
	Create(ctx context.Context, u *user.User) error
}

// IdentityProvider define lo que necesitamos de Cognito / JWT
type IdentityProvider interface {
	ParseToken(ctx context.Context, rawToken string) (*Claims, error)
}

type Service struct {
	users UserRepository
	idp   IdentityProvider
	now   func() time.Time
}

func NewService(users UserRepository, idp IdentityProvider, now func() time.Time) *Service {
	if now == nil {
		now = time.Now
	}
	return &Service{
		users: users,
		idp:   idp,
		now:   now,
	}
}

// EnsureUserFromToken:
// 1) parsea el token
// 2) busca el usuario por sub
// 3) si no existe, lo crea
func (s *Service) EnsureUserFromToken(ctx context.Context, rawToken string) (*user.User, error) {
	// TODO: implementación real
	// por ahora dejamos un stub para compilar
	return nil, nil
}
