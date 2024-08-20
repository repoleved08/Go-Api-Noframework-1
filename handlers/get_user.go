package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/repoleved08/go-api-nof-1/database"
	"github.com/repoleved08/go-api-nof-1/models"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get the user ID from query parameters
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Query the database for the user
	var employee models.Employee
	query := "SELECT id, name, email FROM employee WHERE id=$1"
	err = database.DB.QueryRow(query, id).Scan(&employee.ID, &employee.Name, &employee.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
