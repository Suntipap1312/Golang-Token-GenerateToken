package main

import "time"

func main() {
	GenerateToken()
}

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JwtPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
