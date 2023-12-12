package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	// 引入所需的数据库驱动
)

// Database 是对数据库连接的封装
type Database struct {
	db *sql.DB
}

// NewDatabase 创建并初始化一个新的Database实例
func NewDatabase(password string) (*Database, error) {

	// TODO 加解密
	dbURL := buildDBURL(password)

	// 建立数据库连接
	db, err := sql.Open("driver-name", dbURL)
	if err != nil {
		return nil, err
	}

	configureDBPool(db)

	return &Database{db: db}, nil
}

func (d *Database) NewTable(tableName string) *TableAccessor {
	return &TableAccessor{
		db:    d.db,
		table: tableName,
	}
}

// TableAccessor 提供对特定表的访问
type TableAccessor struct {
	db    *sql.DB
	table string
}

// Query 在特定表上执行查询
func (ta *TableAccessor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// 确保查询只针对指定表
	// 注意：这里的实现需要根据实际的SQL语句和数据库类型进行适配
	query = "SELECT * FROM " + ta.table + " WHERE " + query
	return ta.db.Query(query, args...)
}

// ... 其他数据库操作方法 ...

// buildDBURL 构建数据库连接字符串
func buildDBURL(password string) string {
	// 使用环境变量中的其他数据库配置信息构建连接字符串
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	// ... 其他设置 ...

	return fmt.Sprintf("postgres://%s:%s@%s/%s", user, password, host, dbName)
}

// configureDBPool 配置数据库连接池
func configureDBPool(db *sql.DB) {
	// 设置连接池参数
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
}
