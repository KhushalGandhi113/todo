package jwtAuth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
	jwt.RegisteredClaims
}
type Info struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

var Infos = []Info{}

func Auth(email string, pass string) string {

	mySigningKey := []byte("thisismysecret")

	// create the claims
	claims := &MyCustomClaims{
		email,
		pass,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Ishan",
			Subject:   "Jwt token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	return ss
}

// singup router
func SignUp(c *gin.Context) {
	var info Info
	if err := c.BindJSON(&info); err != nil {
		return
	}
	Infos = append(Infos, info)
	token := Auth(info.Email, info.Pass)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
