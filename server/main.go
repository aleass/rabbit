package main
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type User struct{
	Id int `xorm:"id"`
	Pwd string `xorm:"pwd"`
}
func main3() {
	db,err := MysqlInit()
	fmt.Println("15 err:",err)
	U := make([]User,0)
	db.SQL("select * from user ").Find(&U)
	fmt.Println(U)

	n := new(User)
	n.Id = 6
	n.Pwd = "1111"
	db.Insert(n)


	//c := make(chan int,100)
	//for i:=0;i<100;i++{
	//	c <- i
	//}
	//fmt.Println(<-c)
	//	fmt.Println(len(c))
	//fmt.Println(<-c)

	//r := gin.Default()
	//r.GET("/user/login", func(c *gin.Context){
	//	c.Header("Content-Type", "text/html; charset=utf-8")
	//	c.String(200, `<p>html代码</p>`)
	//})
	//_ = r.Run(":8181")
}

func MysqlInit() (*xorm.Engine,error) {
	var err error
	DsName := fmt.Sprintf("%v:%v@(%v)/%v", "root1", "root", "127.0.0.1:3306", "g")
	fmt.Println(DsName)
	db, err := xorm.NewEngine("mysql", "root1:root@tcp(127.0.0.1:3306)/g?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return db,err
	}
	db.ShowSQL(true)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(240*time.Second)
	//匹配表名
	db.Sync2()
	if err != nil {
		return db,err
	}
	return db,err
}


func main() {
	r1:= NewRabbitMQTopic("hxbExc","huxiaobai.one")
	r2:= NewRabbitMQTopic("hxbExc","huxiaobai.two.cs")
	//for i:=0;i<20;i++{
	//}
	//for {
	r1.PublishTopic("1:今夜不回家")
	r2.PublishTopic("2:今夜回家")
	time.Sleep(1*time.Second)
	//}
}