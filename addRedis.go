package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//http.HandleFunc("/mysql/select",HandleDB)
	//开启监听
	//http.ListenAndServe(":8803", nil)
	HandleDB()
}

func HandleDB() {
	//连接redis数据库
	var conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis连接失败,错误信息：",err)
		return
	}
	defer conn.Close()
	fmt.Println("redis连接成功")

	//验证redis中是否有数据
	reply, err := conn.Do("lrange", "mlist", 0, -1)
	pkeys, _ := redis.Strings(reply, err)
	fmt.Println(pkeys)

	if len(pkeys) > 0 {
		//如果有
		fmt.Println("从redis获得数据")
		// 从redis里直接读取
		for _, key := range pkeys {
			retStrs, _ := redis.Strings(conn.Do("hgetall", key))
			//fmt.Println(retStrs)
			fmt.Printf("{%s %s %s}\n", retStrs[1], retStrs[3], retStrs[5])
		}
	} else {
		//如果没有
		fmt.Println("从mysql获得数据")
	}

	//连接mysql数据库
	db,err:= sql.Open("mysql","root:0322@tcp(127.0.0.1:3306)/testgo?charset=utf8")
	if err!=nil{
		fmt.Println("错误信息：",err)
		return
	}
	fmt.Println("连接成功：",db)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM test_table_name WHERE id=3")
	defer rows.Close()
	if err != nil {
		fmt.Println("查询失败，错误信息为：",err)
		return
	}
	fmt.Println("查询成功")
	var id int
	var name int
	var text string
	for rows.Next() {
		rows.Columns()
		err = rows.Scan(&id, &name, &text)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(text)
	}
	//写入redis并且设置过期时间
		//将p以hash形式写入redis
		_, e1 := conn.Do("id", id, "name", name, "text", text)

		//将这个hash的key加入mlist
		_, e2 := conn.Do("rpush", "mlist", id)

		//设置过期时间
		_, e3 := conn.Do("expire", id, 60)
		_, e4 := conn.Do("expire", "mlist", 60)

		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			fmt.Println("写入失败", e1, e2, e3, e4)
		} else {
			fmt.Println("写入成功")
		}
}
