package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/repoleved08/go-api-nof-1/database"
	"github.com/repoleved08/go-api-nof-1/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO employee (name, email) VALUES ($1, $2) RETURNING id`
	err = database.DB.QueryRow(query, employee.Name, employee.Email).Scan(&employee.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created with id %d", employee.ID)

}
