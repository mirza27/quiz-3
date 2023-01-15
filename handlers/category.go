package handlers

import (
	
	"database/sql"
	"quiz-3/database"
)

// router menampilkan semua kategori
func GetAllCategory(db *sql.DB) (err error, results []database.Category)  {
	sql := "SELECT id, name FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = database.Category{}
		 
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

// mengedit kategori
func InsertCategory(db *sql.DB, category *database.Category) (err error){
	sql := "INSERT INTO category (id, name) VALUES ($1, $2)"

	errs := db.QueryRow(sql, category.ID, category.Name)

	return errs.Err()

}

// mengupdate category
func UpdateCategory(db *sql.DB, category *database.Category) (err error){
	sql := "UPDATE category SET name = $1 WHERE id = $3"

	errs := db.QueryRow(sql, category.Name, category.ID)

	return errs.Err()

}

// mengahpus kategory
func DeleteCategory(db *sql.DB, category *database.Category) (err error){
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}

func GetCategoryBook(db *sql.DB, id_book *database.Book) (err error, results []database.Book){
	sql := "SELECT id, title, description, image_url, release_year, price, total_page, thickness FROM book WHERE category_id = $1"

	rows, err := db.Query(sql, id_book.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = database.Book{}
		 
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}


