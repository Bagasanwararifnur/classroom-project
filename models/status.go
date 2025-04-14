package models

type Status struct {
	ID      int  `json:"id"`
	ClassID int  `json:"class_id"`
	UserID  int  `json:"user_id"`
	Status  bool `json:"status"`
}
