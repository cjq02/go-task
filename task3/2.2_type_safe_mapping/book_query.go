package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

const (
	dbDriver   = "mysql"
	dbUser     = "task3_user"
	dbPassword = "123456"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "task3"
)

func getDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect(dbDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	return db, nil
}

func queryBooksByPrice(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > ? ORDER BY price DESC"

	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, fmt.Errorf("查询书籍信息失败: %v", err)
	}

	return books, nil
}

func main() {
	db, err := getDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	books, err := queryBooksByPrice(db, 50.0)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("价格大于 50 元的书籍数量: %d\n", len(books))
}
