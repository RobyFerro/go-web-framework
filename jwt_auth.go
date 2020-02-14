package go_web_framework

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type Auth struct {
	User struct {
		ID       uint
		Name     string
		Surname  string
		Username string
		Password string
	} `json:"user"`
	Token string
	Conf  Conf
}

// Prepare Auth structure for Service Container
func SetAuth(conf Conf) *Auth {
	return &Auth{Conf: conf}
}

// Get user struct from authentication token (JWT)
func (c *Auth) GetUser(req *http.Request) error {
	bearerSchema := "Bearer "
	tokenString := req.Header.Get("Authorization")

	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString[len(bearerSchema):], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.Conf.App.Key), nil
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

// Issue new JWT token
func (c *Auth) NewToken() bool {
	c.User.Password = ""
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  time.Now().Add(time.Hour * time.Duration(2)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.Conf.App.Key))
	c.Token = tokenString

	if err != nil {
		return false
	}

	return true
}

// Refresh JWT token
func (c *Auth) RefreshToken() bool {
	expirationTime := time.Now().Add(5 * time.Minute)
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  expirationTime,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.Conf.App.Key))
	c.Token = tokenString

	if err != nil {
		return false
	}

	return true
}
