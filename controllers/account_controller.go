package controllers

import (
	"go-crud-live/config"
	"go-crud-live/models"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Account

	// GORM way: Find all records in 'accounts' table
	if err := config.GormDB.Find(&accounts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	println("w value", json.NewEncoder(w).Encode(accounts))
	// Return accounts in JSON format
	json.NewEncoder(w).Encode(accounts)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var accounts []models.Account
	if err := config.GormDB.First(&accounts, id).Error; err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(accounts)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var account models.Account
	if err := config.GormDB.First(&account, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&account)

	if err := config.GormDB.Save(&account).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func DeleteAccout(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := config.GormDB.Delete(&models.Account{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Account is delete"})
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	_ = json.NewDecoder(r.Body).Decode(&account)

	if err := config.GormDB.Create(&account).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}
