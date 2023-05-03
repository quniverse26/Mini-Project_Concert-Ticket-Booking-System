package middleware

import (
	"github.com/quniverse26/miniproject/model"

    "github.com/dgrijalva/jwt-go"
    "time"
)

func generateToken(admin *model.Admin) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = admin.Email
    claims["password"] = admin.Password
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    // Generate encoded token and return it
    encodedToken, err := token.SignedString([]byte("sayaadmin"))
    if err != nil {
        return "", err
    }

    return encodedToken, nil
}
