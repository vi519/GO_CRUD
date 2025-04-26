package controllers

import (
	"encoding/json"
	"go-crud-live/config"
	"go-crud-live/models"
	"net/http"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Account

	// GORM way: Find all records in 'accounts' table
	if err := config.GormDB.Find(&accounts).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return accounts in JSON format
	json.NewEncoder(w).Encode(accounts)
}
