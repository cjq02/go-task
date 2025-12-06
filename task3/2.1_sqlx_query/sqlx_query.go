package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
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

func queryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE department = ?"

	err := db.Select(&employees, query, department)
	if err != nil {
		return nil, fmt.Errorf("查询员工信息失败: %v", err)
	}

	return employees, nil
}

func queryHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"

	err := db.Get(&employee, query)
	if err != nil {
		return Employee{}, fmt.Errorf("查询工资最高员工失败: %v", err)
	}

	return employee, nil
}

func main() {
	db, err := getDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	employees, err := queryEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	employee, err := queryHighestSalaryEmployee(db)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("部门为 \"技术部\" 的员工数量: %d\n", len(employees))
	fmt.Printf("工资最高的员工: %s\n", employee.Name)
}
