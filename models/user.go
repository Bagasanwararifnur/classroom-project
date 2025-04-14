package models

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Thread   []Thread  `gorm:"foreignKey:UserID" swaggerignore:"true" json:"threads"`
	Comment  []Comment `gorm:"foreignKey:UserID" swaggerignore:"true" json:"comments"`
	Status   []Status  `gorm:"foreignKey:UserID" swaggerignore:"true" json:"-"`
}
