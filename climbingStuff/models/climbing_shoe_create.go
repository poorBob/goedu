package models

type ClimbingShoeCreate struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Size  int    `json:"size"`
}
