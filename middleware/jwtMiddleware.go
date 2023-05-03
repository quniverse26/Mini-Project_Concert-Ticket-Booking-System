package middleware

import (
    "net/http"

    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    //"github.com/labstack/echo/v4/middleware"
)

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // get token from Authorization header
        token := c.Request().Header.Get("Authorization")

        // parse token
        parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            // validate signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
            }

            // validate secret
            secret := "sayaadmin" // replace with your secret key
            return []byte(secret), nil
        })

        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
        }

        if parsedToken == nil || !parsedToken.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
        }

        return next(c)
    }
}
