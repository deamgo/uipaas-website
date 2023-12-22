package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var MySecret = getSecret()

type MyClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var TokenBlacklist = make(map[string]bool)

// defineTheExpirationTime
const TokenExpireDuration = time.Hour * 2

// generate jwt
func GenToken(id string) (string, error) {
	c := MyClaims{

		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}
	// Creates a signature object using the specified signature method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Sign with the specified secret and obtain the completed encoded string token
	return token.SignedString(MySecret)
}

// parsingJWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// analysis token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func IsExpireToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return time.Now().Unix() < claims.ExpiresAt, nil
	}

	return false, errors.New("invalid token")
}

func ExtractIDFromToken(tokenString string) (string, error) {
	parts := strings.SplitN(tokenString, " ", 2)
	claim, err := ParseToken(parts[1])
	if err != nil {
		return "", err
	}
	return claim.ID, nil
}

func getSecret() []byte {
	// set The Name And Path Of The Profile
	viper.SetConfigName("conf")
	viper.AddConfigPath("./auth/jwt")

	// read The Configuration File
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}
	MySecret := viper.GetString("jwt_secret")
	return []byte(MySecret)
}

// Log out of the token
func RevokeToken(tokenString string) {
	TokenBlacklist[tokenString] = true
}
