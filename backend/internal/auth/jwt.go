package auth

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret     string
	jwtTTLSeconds int
)

// Init initialisiert das Auth Paket mit Secret und TTL (in Sekunden).
// Sollte Init nicht aufgerufen werden, wird ein Fehler geworfen beim Token-Generieren.
func Init(secret string, ttlSeconds int) {
	jwtSecret = secret
	jwtTTLSeconds = ttlSeconds
}

// GenerateToken erstellt einen JWT Token mit dem gegebenen userID.
// TTL wird in Sekunden aus der Umgebungsvariable JWT_TTL gelesen (fallback 3600).
func GenerateToken(userID uint) (string, error) {
	secret := jwtSecret
	if secret == "" {
		return "", fmt.Errorf("JWT secret not initialized (call auth.Init)")
	}

	ttl := 3600
	if jwtTTLSeconds > 0 {
		ttl = jwtTTLSeconds
	}

	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"exp": time.Now().Add(time.Duration(ttl) * time.Second).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}

// ParseToken gibt die userID (subject) zurück, falls gültig
func ParseToken(tkn string) (uint, error) {
	secret := jwtSecret
	if secret == "" {
		return 0, fmt.Errorf("JWT secret not initialized (call auth.Init)")
	}

	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(string); ok {
			var id uint
			fmt.Sscanf(sub, "%d", &id)
			return id, nil
		}
	}
	return 0, fmt.Errorf("invalid token claims")
}
