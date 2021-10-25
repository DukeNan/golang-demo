package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	ID   int
	Age  int
	Name string
}

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func initDB() (err error) {
	dsn := "root:yd_db_pass@tcp(192.168.100.61:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 3)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select * from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	ret, err := db.Exec(sqlStr, "王冰冰", "20")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入的ID
	if err != nil {
		fmt.Printf("get lastinsert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 19, 8)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", n)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id =?"
	ret, err := db.Exec(sqlStr, 5)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响行数
	if err != nil {
		fmt.Printf("get  RowAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected raws:%d\n", n)
}

// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段。
func insertUserDemo() (err error) {
	sqlStr := "insert into user(name, age) values(:name, :age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "张国荣",
			"age":  50,
		})
	return
}

// 与DB.NamedExec同理，这里是支持查询。
func namedQuery() {
	sqlStr := "select * from user where name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "王冰冰"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
	u := user{
		Name: "王力宏",
	}
	// 使用结构体命名查询，根据结构体字段的db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%v\n", u)
		fmt.Printf("user:%+v\n", u)
	}
}

// 事务操作
func transactionDemo2() (err error) {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit")
		}
	}()
	sqlStr1 := "update user set age=20 where id=?"
	rs, err := tx.Exec(sqlStr1, 6)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "update user set age=55 where id=?"
	rs, err = tx.Exec(sqlStr2, 9)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return err
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("insert into user (name, age) values %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"insert into user (name, age) values (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err := db.Exec(query, args...)
	return err
}

// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := db.NamedExec("insert into user (name, age) values (:name, :age)", users)
	return err
}

// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("select * from user where id in (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return users, err
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充ID
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("select id, name, age from user where id in (?) order by FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	// qlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return users, err
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// queryRowDemo()
	// queryMultiRowDemo()
	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()
	// insertUserDemo()
	// namedQuery()
	// transactionDemo2()
	fmt.Println("===================")
	// 方式1
	u1 := User{Name: "梅艳芳", Age: 18}
	u2 := User{Name: "张敏", Age: 18}
	u3 := User{Name: "邱淑贞", Age: 18}
	// users := []*User{&u1, &u2, &u3}
	// err = BatchInsertUsers(users)
	// if err != nil {
	// 	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	// }
	// 方式2
	// users2 := []interface{}{u1, u2, u3}
	// err = BatchInsertUsers2(users2)
	// if err != nil {
	// 	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	// }
	// 方式3
	users3 := []*User{&u1, &u2, &u3}
	err = BatchInsertUsers3(users3)
	if err != nil {
		fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	}
	fmt.Println("=============sqlx.In的查询===============")
	// ids := []int{1, 3, 4}
	// users, err := QueryByIDs(ids)
	// if err != nil {
	// 	fmt.Printf("err:%v\n", err)
	// }
	// fmt.Printf("users:%#v\n", users)
	fmt.Println("============QueryAndOrderByIDs==============")
	ids := []int{1, 29, 4}
	users, err := QueryAndOrderByIDs(ids)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("users:%#v\n", users)

}
