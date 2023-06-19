package models

type Books struct {
	Id          int    `json:"id" form:"id" gorm:"unique;not null;PRIMARY_KEY;AUTO_INCREMENT"`
	Title       string `json:"title" form:"title" binding:"required" gorm:"not null"`
	Author      string `json:"author" form:"author" binding:"required" gorm:"not null"`
	Description string `json:"description" form:"description" binding:"required" gorm:"not null"`
	Stock       int    `json:"stock" form:"stock" binding:"required" gorm:"not null"`
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

const (
	USER     = "admin"
	PASSWORD = "password123"
	SECRET   = "secret"
)
