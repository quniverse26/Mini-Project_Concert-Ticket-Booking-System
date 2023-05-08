package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var admins []Admin

type AdminRegister struct {
	Name 	 string `json:"name"`
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type AdminJWTDecode struct {
	ID uint
	Name string
}

// type JWTClaims struct {
// 	ID    uint   `json:"id"`
// 	Name string `json:"name"`
// 	jwt.RegisteredClaims
// }

func (a *Admin) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hash)
	return nil
}

func (a *Admin) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
	return err == nil
}

func (a *Admin) GenerateToken() (string, error) {
	claims := &JWTClaims{
		ID: a.ID,
		Name: a.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}