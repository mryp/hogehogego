package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" //dbrで使用する
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

//---------------------------------------------------------------------
//定数
//---------------------------------------------------------------------

//DBアクセス関連定数
const (
	DBUserID     = "root"
	DBPassword   = "root"
	DBHostName   = "127.0.0.1"
	DBPortNumber = "3306"
	DBName       = "hogehogego"
)

//---------------------------------------------------------------------
//構造体
//---------------------------------------------------------------------

//ConnectResponce 接続構造体
type ConnectResponce struct {
	Result int `json:"result"`
}

//ConnectTable 接続テーブル
type ConnectTable struct {
	ID     int64     `db:"id"`
	Termid string    `db:"termid"`
	Status int       `db:"status"`
	Uptime time.Time `db:"uptime"`
}

//---------------------------------------------------------------------
//関数
//---------------------------------------------------------------------

//ConnectHandler 接続記録を行う
func ConnectHandler(c echo.Context) error {
	termid := c.QueryParam("termid")
	status, err := strconv.ParseInt(c.QueryParam("status"), 10, 0)
	if termid == "" || err != nil {
		fmt.Printf("ConnectHandler param error err=%v, termid=%v\n", err, termid)
		return c.JSON(http.StatusOK, ConnectResponce{Result: -1})
	}

	record := ConnectTable{Termid: termid, Status: int(status), Uptime: time.Now()}

	session := connectDB()
	if session == nil {
		return c.JSON(http.StatusOK, ConnectResponce{Result: -1})
	}
	_, errinsert := session.InsertInto("connect_table").
		Columns("termid", "status", "uptime").
		Record(record).
		Exec()

	result := 0
	if errinsert != nil {
		fmt.Printf("ConnectHandler insert error err=%v\n", errinsert)
		result = -1
	}

	return c.JSON(http.StatusOK, ConnectResponce{Result: result})
}

func connectDB() *dbr.Session {
	db, err := dbr.Open("mysql", DBUserID+":"+DBPassword+"@tcp("+DBHostName+":"+DBPortNumber+")/"+DBName, nil)
	if err != nil {
		fmt.Printf("connectDB err=%v\n", err)
		return nil
	}

	dbsession := db.NewSession(nil)
	return dbsession
}
