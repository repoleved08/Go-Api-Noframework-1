package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/repoleved08/go-api-nof-1/database"
	"github.com/repoleved08/go-api-nof-1/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}

	rows, err := database.DB.Query("SELECT id, name, email FROM employee")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, employee)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
