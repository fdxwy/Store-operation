package main
import (
   "database/sql"
   "fmt"
   _ "github.com/go-sql-driver/mysql"
)

type TestMysql struct {
   db *sql.DB
}

/* 初始化数据库引擎 */
func Init() (*TestMysql,error){
   test := new(TestMysql)
   db,err := sql.Open("mysql","root:0322@tcp(127.0.0.1:3306)/testgo?charset=utf8");
   if err!=nil {
      fmt.Println("database initialize error : ",err.Error());
      return nil,err
   }
   test.db = db
   return test,nil
}

/* 测试数据库数据添加 */
func (test *TestMysql)Create(){
   if test.db==nil {
      return
   }
   stmt,err := test.db.Prepare("insert into test_table_name(id,name,text)values(?,?,?)");
   if err!=nil {
      fmt.Println(err.Error());
      return
   }
   defer stmt.Close();
   if result,err := stmt.Exec(4,88,"北京");err==nil {
      if id,err := result.LastInsertId();err==nil {
         fmt.Println("insert id : ",id)
      }
   }
   if result,err := stmt.Exec(5,88,"上海");err==nil {
      if id,err := result.LastInsertId();err==nil {
         fmt.Println("insert id : ",id)
      }
   }
   if result,err := stmt.Exec(6,99,"杭州");err==nil {
      if id,err := result.LastInsertId();err==nil {
         fmt.Println("insert id : ",id)
      }
   }

}

func (test *TestMysql)Close(){
   if test.db!=nil {
      test.db.Close()
   }
}

func main(){
   if test,err := Init();err==nil {
      test.Create()
      test.Close()
   }
}
