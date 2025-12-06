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

func transfer(db *sql.DB, fromAccountID int, toAccountID int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开始事务失败: %v", err)
	}

	var balance float64
	checkBalanceSQL := "SELECT balance FROM accounts WHERE id = ? FOR UPDATE"
	err = tx.QueryRow(checkBalanceSQL, fromAccountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("查询账户余额失败: %v", err)
	}

	if balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足，当前余额: %.2f，需要: %.2f", balance, amount)
	}

	deductSQL := "UPDATE accounts SET balance = balance - ? WHERE id = ?"
	_, err = tx.Exec(deductSQL, amount, fromAccountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("扣除账户余额失败: %v", err)
	}

	addSQL := "UPDATE accounts SET balance = balance + ? WHERE id = ?"
	_, err = tx.Exec(addSQL, amount, toAccountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("增加账户余额失败: %v", err)
	}

	insertTransactionSQL := "INSERT INTO transactions (from_account_id, to_account_id, amount) VALUES (?, ?, ?)"
	_, err = tx.Exec(insertTransactionSQL, fromAccountID, toAccountID, amount)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易信息失败: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("提交事务失败: %v", err)
	}

	return nil
}

func main() {
	db, err := getDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	fromAccountID := 1
	toAccountID := 2
	amount := 100.00

	err = transfer(db, fromAccountID, toAccountID, amount)
	if err != nil {
		log.Fatalf("转账失败: %v", err)
	}

	fmt.Printf("转账成功：从账户 %d 向账户 %d 转账 %.2f 元\n", fromAccountID, toAccountID, amount)
}
