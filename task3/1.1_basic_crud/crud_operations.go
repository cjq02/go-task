package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver   = "mysql"
	dbUser     = "task3_user"
	dbPassword = "123456"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "task3"
)

func getDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("打开数据库连接失败: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接测试失败: %v", err)
	}

	return db, nil
}

func main() {
	db, err := getDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	insertSQL := "INSERT INTO students (name, age, grade) VALUES (?, ?, ?)"
	result, err := db.Exec(insertSQL, "张三", 20, "三年级")
	if err != nil {
		log.Fatalf("插入记录失败: %v", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("1. 插入成功，新记录ID: %d\n", id)

	querySQL := "SELECT id, name, age, grade FROM students WHERE age > ?"
	rows, err := db.Query(querySQL, 18)
	if err != nil {
		log.Fatalf("查询记录失败: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n2. 年龄大于 18 岁的学生信息：")
	for rows.Next() {
		var id int
		var name, grade string
		var age int
		if err := rows.Scan(&id, &name, &age, &grade); err != nil {
			log.Fatalf("扫描记录失败: %v", err)
		}
		fmt.Printf("%-2d %-6s %-3d %-6s\n", id, name, age, grade)
	}

	updateSQL := "UPDATE students SET grade = ? WHERE name = ?"
	result, err = db.Exec(updateSQL, "四年级", "张三")
	if err != nil {
		log.Fatalf("更新记录失败: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\n3. 更新成功，受影响行数: %d\n", rowsAffected)

	deleteSQL := "DELETE FROM students WHERE age < ?"
	result, err = db.Exec(deleteSQL, 15)
	if err != nil {
		log.Fatalf("删除记录失败: %v", err)
	}
	rowsAffected, _ = result.RowsAffected()
	fmt.Printf("\n4. 删除成功，受影响行数: %d\n", rowsAffected)
}
