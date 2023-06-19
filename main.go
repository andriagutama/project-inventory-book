package main

import (
	apps "project-inventory-book/app"
	"project-inventory-book/auth"
	"project-inventory-book/db"
	"project-inventory-book/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	handler := apps.New(conn)

	// all routes

	// home
	router.GET("/", auth.HomeHandler)

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// all book
	router.GET("/books", middleware.AuthValid, handler.GetBooks)
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	// add book
	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	// update book
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBook)

	// delete book
	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
