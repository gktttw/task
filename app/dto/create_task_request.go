package dto

type TaskRequest struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}
