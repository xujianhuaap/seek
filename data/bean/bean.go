package bean

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id int8
	Mobile string `json:"mobile"`
	PassWord string `json:"pass_word"`
}

func LoadUserFromDB(mobile,password string) bool {
	fmt.Printf("LoadUserFromDB mobile\t%v\tpassword\t%v\t\n",mobile,password)

	var aMobile,aPassword string
	db, err := sql.Open("mysql", "root:123456@/seek")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = db.QueryRow("SELECT mobile,password FROM user where mobile = ? and password = ?",mobile,password).Scan(&aMobile,&aPassword)
	fmt.Printf("query result Mobile %v\t,Password %v \t\n",aMobile,aPassword)
	if err != nil {
		return false
	}else {
		return true
	}

}
type UserInfo struct {
	Id int8 `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Gender bool `json:"gender"`
}

type BorrowRecord struct {
	Id int32 ``
}
type SysMessage struct {
	Id int16 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Status string `json:"status"`
}
