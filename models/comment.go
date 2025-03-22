package models

type Comment struct {
	ID            int    `json:"id"`
	Comment       string `json:"content"`
	UserID        int    `json:"user_id"`
	ThreadID      int    `json:"thread_id"`
	ThreadDetails Thread `gorm:"foreignKey:ThreadID" swaggerignore:"true" json:"thread_detail"`
}
