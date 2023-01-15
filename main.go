package main

import (
	"quiz-3/handlers"
	"database/sql"
	"quiz-3/database"
	"quiz-3/controllers"
	"github.com/gin-gonic/gin"
	"fmt"
	//"net/http"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "mirza"
	password = "12Dwiana!"
	dbname = "db_quiz3"
)

var (
	DB *sql.DB
	err error
)


func main() {
	
	// Database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB connection failed")
		panic(err)
	} else {
		fmt.Println("DB connection success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()
	

	// router query
	router.GET("/segitiga-sama-sisi", handlers.QuerySegitiga)
	router.GET("/persegi", handlers.QueryPersegi)
	router.GET("/persegi-panjang", handlers.QueryPersegiPanjang)
	router.GET("/lingkaran", handlers.QueryLingkaran)


	// router category
	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
	router.GET("/categories/:id/books", controllers.GetCategoryBook)
	
	// router books
	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", controllers.InsertBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)


	router.Run()
}