package main

import (
   "database/sql"
   "fmt"
   _ "github.com/go-sql-driver/mysql"
)

func main() {
   //http.HandleFunc("/mysql/select",HandleDB)
   //开启监听
   //http.ListenAndServe(":8803", nil)
   HandleDB()

}

func HandleDB() {
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
}
