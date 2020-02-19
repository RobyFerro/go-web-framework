package gwf

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Auth structure will be used to handle the authenticated user data.
type Auth struct {
	User struct {
		ID       uint
		Name     string
		Surname  string
		Username string
		Password string
	} `json:"user"`
	Key string
}

// GetUser will parse incoming request and returns the user data.
func (c *Auth) GetUser(req *http.Request) error {
	bearerSchema := "Bearer "
	tokenString := req.Header.Get("Authorization")

	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString[len(bearerSchema):], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.Key), nil
	})

	for key, val := range claims {
		if key == "user" {
			userData := val.(string)
			err := json.Unmarshal([]byte(userData), &c.User)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// NewToken will return a new JWT token
func (c *Auth) NewToken() (string, bool) {
	c.User.Password = ""
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  time.Now().Add(time.Hour * time.Duration(2)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.Key))

	if err != nil {
		return "", false
	}

	return tokenString, true
}

// RefreshToken will grefresh the a speficic token
func (c *Auth) RefreshToken(req http.ResponseWriter) bool {
	expirationTime := time.Now().Add(5 * time.Minute)
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  expirationTime,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.Key))
	req.Header().Set("refresh-token", tokenString)

	if err != nil {
		return false
	}

	return true
}
