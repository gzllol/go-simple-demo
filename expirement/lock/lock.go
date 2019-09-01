package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "test"
)

func Connect() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	fmt.Println(dsn)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	return DB
}

type User struct {
	ID      int64          `db:"id"`
	L       sql.NullString `db:"l"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Version sql.NullInt64  `db:"version"`
}

func QueryOne(id int) *User {
	user := new(User)
	row := DB.QueryRow("select * from test where id=?", id)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.L, &user.Version); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return nil
	}
	return user
}

func Clear(id int) {
	DB.Exec("update test set l='' where id=?", id)
}

func Add(id, num int) {
	old := QueryOne(id)
	DB.Exec("update test set l=?", old.L.String+","+strconv.Itoa(num))
}

func Check(id, max int) bool {
	u := QueryOne(id)
	strs := strings.Split(u.L.String, ",")
	flags := make([]bool, max)
	for _, s := range strs {
		if len(s) == 0 {
			continue
		}
		index, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err.Error())
		}
		flags[index] = true
	}
	for i := 0; i < max; i++ {
		if !flags[i] {
			return false
		}
	}
	return true
}

func AddOptimisticLock(id, num int, sleep int) {
	old := QueryOne(id)
	for {
		result, err := DB.Exec("update test set l=? where id=? and version=?", old.L.String+","+strconv.Itoa(num), id, old.Version)
		if err != nil {
			time.Sleep(sleep)
		} else if result != nil {
			raf, err2 := result.RowsAffected()
			if err2 != nil || raf != 1 {
				time.Sleep(sleep)
			}
		}
	}
}

func AddToMaxOptimisticLock(id, max int) {
	for i := 0; i < max; i++ {
		Add(id, i)
	}
}

func main() {
	id := 1
	max := 1000
	DB = Connect()
	Clear(id)
	start := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		Add(id, i)
	}
	end := time.Now().UnixNano()
	fmt.Printf("%v\n", float64(max)/float64((end-start))*float64(1000000000))
	if !Check(id, max) {
		fmt.Printf("id - %v, max - %v check failed\n", id, max)
	}
}
