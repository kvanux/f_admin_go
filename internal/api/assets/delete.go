package assets

import (
	"encoding/json"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/db"
	"f_admin_go/internal/models"
	"net/http"
)

func handleDeleteAsset(w http.ResponseWriter, r *http.Request) {
	var req models.AssetDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	_, err := db.DB.Exec("DELETE FROM assets WHERE id = $1", req.ID)
	if err != nil {
		shared.WriteError(w, http.StatusInternalServerError, "Database delete error")
		return
	}

	w.WriteHeader(http.StatusOK)
}
