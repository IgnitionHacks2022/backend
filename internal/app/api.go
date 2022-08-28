package app

import (
	"backend/internal/entity"
	"backend/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func StatisticsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	fmt.Println(userID)

	conn, err := db.Connection()
	items, err := db.GetList(conn, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	payload := entity.StatisticsResponse{
		Items: items,
	}

	jsonPayload, err := json.Marshal(payload)

	w.Write(jsonPayload)

}
