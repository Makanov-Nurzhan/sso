package auth

import (
	"context"
	"grpc-service/internal/domain/models"
	"log/slog"
)

type Auth struct {
	log *slog.Logger
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
