package controllers

import(
	"github.com/gin-gonic/gin"
	"quiz-3/database"
	"quiz-3/handlers"
	"net/http"
	"strconv"
)

func GetAllCategory(c *gin.Context){
	var (
		result gin.H
	)

	err, categories := handlers.GetAllCategory(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result" : err,
		}
	} else {
		result = gin.H{
			"result" : categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context){
	var category database.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = handlers.InsertCategory(database.DbConnection, &category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result" : "Success Insert Category",
	})

}


func UpdateCategory(c *gin.Context){
	var category database.Category

	id,_ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = handlers.UpdateCategory(database.DbConnection, &category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result" : "Success Update Category",
	})
}


func DeleteCategory(c *gin.Context) {
	var category database.Category

	id,err := strconv.Atoi(c.Param("id"))

	category.ID = int64(id)
	

	err = handlers.DeleteCategory(database.DbConnection, &category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result" : "Success Delete Category",
	})
}


func GetCategoryBook(c *gin.Context){
	var (
		result gin.H
		book database.Book
	)


	id, err := strconv.Atoi(c.Param("id"))


	err = c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = int64(id)

	err, categories := handlers.GetCategoryBook(database.DbConnection, &book)
	if err != nil {
		result = gin.H{
			"result" : err,
		}
	} else {
		result = gin.H{
			"result" : categories,
		}
	}

	c.JSON(http.StatusOK, result)

}