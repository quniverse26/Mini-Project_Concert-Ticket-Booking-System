package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Buyer struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var buyers []Buyer

type BuyerRegister struct {
	Name 	 string `json:"name"`
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type BuyerLogin struct {
	Email 	string `json:"email"`
	Password string `json:"password"`
}

type BuyerJWTDecode struct {
	ID uint
	Name string
}

type JWTClaims struct {
	ID    uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// func (b *Buyer) BeforeCreate(tx *gorm.DB) (err error) {
// 	hash, _ := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
// 	b.Password = string(hash)
// 	return
// }

func (b *Buyer) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	b.Password = string(hash)
	return nil
}

func (b *Buyer) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(b.Password), []byte(plain))
	return err == nil
}

func (b *Buyer) GenerateToken() (string, error) {
	claims := &JWTClaims{
		ID: b.ID,
		Name: b.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}