package jwt_test

import (
	jwt_lib "github.com/golang-jwt/jwt/v5"
	"grpc-service/internal/domain/models"
	"grpc-service/internal/lib/jwt"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	user := models.User{
		ID:    1,
		Email: "test@example.com",
	}
	app := models.App{
		ID:     123,
		Secret: "mysecretkey",
	}
	duration := time.Hour

	tokenStr, err := jwt.NewToken(user, app, duration)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if tokenStr == "" {
		t.Fatal("expected non-empty token string")
	}

	// Проверка валидности токена
	parsedToken, err := jwt_lib.Parse(tokenStr, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(app.Secret), nil
	})
	if err != nil || !parsedToken.Valid {
		t.Fatalf("token should be valid, got err: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwt_lib.MapClaims)
	if !ok {
		t.Fatal("expected map claims")
	}

	if claims["uid"] != float64(user.ID) { // jwt возвращает числа как float64
		t.Errorf("expected uid %v, got %v", user.ID, claims["uid"])
	}
	if claims["email"] != user.Email {
		t.Errorf("expected email %v, got %v", user.Email, claims["email"])
	}
	if claims["app_id"] != float64(app.ID) {
		t.Errorf("expected app_id %v, got %v", app.ID, claims["app_id"])
	}
	if exp, ok := claims["exp"].(float64); !ok || int64(exp) < time.Now().Unix() {
		t.Error("expiration time is invalid or already expired")
	}
}
