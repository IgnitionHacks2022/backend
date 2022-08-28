package entity

import "backend/pkg/db"

type StatisticsResponse struct {
	Items []db.Item `json:"items"`
}
