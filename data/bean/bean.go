package bean

import (
	"database/sql"
	"fmt"
	"github.com/xujianhuaap/seek/data"
)

type User struct {
	Id int8
	Mobile string `json:"mobile"`
	PassWord string `json:"pass_word"`
}
func DeleteUserFromDB(id int64) bool {
	db,err:= sql.Open(data.DataBaseType,data.DataBaseLoginInfo)
	defer db.Close()
	if err == nil {
		err=db.Ping()
		if err == nil {
			_,err := db.Exec("DELETE FROM user where id = ?",id)
			if err == nil{
				return true
			}
		}
	}
	return false
}
func WriteUserToDB(mobile,password string) bool {
	db,err:= sql.Open(data.DataBaseType,data.DataBaseLoginInfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err=db.Ping()
	if err != nil {
		panic(err)
	}
	stmIn,err:=db.Prepare("INSERT INTO user VALUES (?,?,?,?,?)")
	defer stmIn.Close()
	if err != nil {
		panic(err)
	}
	_,err = stmIn.Exec(nil,mobile,"unkown",password,nil)

	if err != nil {
		return false
	}
	return true

}
func LoadUserFromDB(mobile,password string) (bool,string) {
	fmt.Printf("LoadUserFromDB mobile\t%v\tpassword\t%v\t\n",mobile,password)

	var aMobile,aPassword string
	 var aId string;
	db, err := sql.Open(data.DataBaseType,data.DataBaseLoginInfo)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = db.QueryRow("SELECT mobile,password,id FROM user where mobile = ?",mobile).Scan(&aMobile,&aPassword,&aId)
	fmt.Printf("query result Mobile %v\t,Password %v \t\n",aMobile,aPassword)
	if err != nil {
		return false,"waitting for a moment retry"
	}else {
		if(aMobile != ""){
			if(aPassword != password){
				return false,"password not conrrect"
			}else {
				return true,aId;
			}
		}else {
			return false,"has not reistered"
		}

	}

}

func HasLogin(mobile string) bool {
	fmt.Printf("HasLogin mobile\t%v\n",mobile)

	var aMobile string
	db, err := sql.Open(data.DataBaseType,data.DataBaseLoginInfo)
	if err != nil{
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = db.QueryRow("SELECT mobile FROM user where mobile = ?",mobile).Scan(&aMobile)

	if err == nil && aMobile != "" {
		fmt.Printf("query result Mobile %v\t\n",aMobile)
		return true
	}
	return false

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
