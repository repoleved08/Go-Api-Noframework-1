package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/repoleved08/go-api-nof-1/database"
	"github.com/repoleved08/go-api-nof-1/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
	}
	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE employee SET name=$1, email=$2 WHERE id=$3`
	_, err = database.DB.Exec(query, employee.Name, employee.Email, employee.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Employee Updated successfully!"))
}
