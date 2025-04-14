package models

type Class struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Grade  int    `json:"grade"`
	Topic  string `json:"topic"`
	Status Status `gorm:"foreignKey:ClassID" json:"status" swaggerignore:"true"`
}
