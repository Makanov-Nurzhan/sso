package auth

import (
	"context"
	"grpc-service/internal/domain/models"
	"log/slog"
	"time"
)

type Auth struct {
	log         *slog.Logger
	usrSaver    UserSaver
	usrProvider UserProvider
	aaProvider  UserProvider
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(ctx context.Context, appID string) (models.App, error)
}

// New returns a new instance of the Auth service.
func New(
	log *slog.Logger,
	usrSaver UserSaver,
	usrProvider UserProvider,
	aaProvider UserProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSaver:    usrSaver,
		usrProvider: usrProvider,
		log:         log,
		aaProvider:  aaProvider,
		tokenTTL:    tokenTTL,
	}
}

// Login checks if user with given credentials exists in the system and returns access token.
//
// If user exists, but password is incorrect, returns error.
// If user doesn't exist, returns error.
func (a *Auth) Login(ctx context.Context, email, password string) (string, error) {
	panic("implement me")
}

// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (a *Auth) RegisterNewUser(ctx context.Context, email, password string) error {
	const op = "auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("Registering user")
}

// IsAdmin checks if user is admin.
func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("implement me")
}
