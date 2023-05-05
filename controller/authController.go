package controller

import (
	"encoding/json"
	"net/http"

	"github.com/quniverse26/miniproject/config"
	"github.com/quniverse26/miniproject/model"
	"github.com/quniverse26/miniproject/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register model.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		utils.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		utils.Response(w, 400, "Password not match", nil)
		return
	}

	passwordHash, err := utils.HashPassword(register.Password)
	if err != nil {
		utils.Response(w, 500, err.Error(), nil)
		return
	}

	buyer := model.Buyer{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := config.DB.Create(&buyer).Error; err != nil {
		utils.Response(w, 500, err.Error(), nil)
		return
	}

	utils.Response(w, 201, "Register Successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login model.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.Response(w, 500, err.Error(), nil)
		return
	}

	var buyer model.Buyer
	if err := config.DB.First(&buyer, "email = ?", login.Email).Error; err != nil {
		utils.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := utils.VerifyPassword(buyer.Password, login.Password); err != nil {
		utils.Response(w, 404, "Wrong email or password", nil)
		return
	}

	token, err := utils.CreateToken(&buyer)
	if err != nil {
		utils.Response(w, 500, err.Error(), nil)
		return
	}

	utils.Response(w, 200, "Successfully Login", token)
}