package service

import (
  "crypto/sha1"
  "fmt"
  "errors"
  "github.com/dgrijalva/jwt-go"
  "github.com/Sedokovka/simple-to-do-app"
  "github.com/Sedokovka/simple-to-do-app/pkg/repository"
  "time"
)

const (
  salt = "sdksdnb12e2fvs90sdcq632"
  signingKey = "hesdbgkiw82793#$%^JHBJGVF"
  tokenTTL = 12 * time.Hour
)

type AuthService struct {
  repo repository.Authorization
}

type tokenClaims struct {
  jwt.StandardClaims
  UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
  return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user crud.User) (int, error) {
  user.Password = s.generatePasswordHash(user.Password)
  return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
  hash := sha1.New()
  hash.Write([]byte(password))

  return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}


func(s *AuthService) GenerateToken(username, password string) (string, error) {
  password = s.generatePasswordHash(password)
  user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
  token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

  if err != nil {
    return 0, err
  }

  claims, ok := token.Claims.(*tokenClaims)
  if !ok {
    return 0, errors.New("token claims are not of type *tokenClaims")
  }

  return claims.UserId, nil
}
