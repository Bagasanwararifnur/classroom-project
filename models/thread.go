package models

type Thread struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  int       `json:"user_id"`
	Comment []Comment `gorm:"foreignKey:ThreadID" json:"comments" swaggerignore:"true"`
}
