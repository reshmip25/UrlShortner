package model

import (
	"database/sql"
	"fmt"
	"log"

)

var db *sql.DB
//var Count uint64


func Index() {
	var err error

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/hello")


	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	 //err = db.QueryRow("SELECT COUNT(*) FROM UrlData").Scan(&Count)
}


func Create(hashedValue string ,longUrl string,custom string) (error , uint64){

	var ID uint64
	var hits int

	_, err := db.Query("Insert into UrlData (hashedValue,longURL,hits) Values (?,?,?)", hashedValue, longUrl,hits)

	if(err == nil){

		db.QueryRow("select ID from UrlData where hashedValue =?",hashedValue).Scan(&ID)

		return err,ID

	}

	return err,ID
}


func ShowLong(shortUrl string)(string,error){

	var hits int
	var long string

	err := db.QueryRow("Select longURL, hits from UrlData where shortUrl = ?",shortUrl).Scan(&long,&hits)

	Update(shortUrl,hits)
	return long,err

}


func ShowShort(hashedValue string) (string){

	var shortU string
	_  = db.QueryRow("Select shortUrl from UrlData where hashedValue = ?",hashedValue).Scan(&shortU)

	return shortU
}

func Update(shortUrl string,hits int){

	db.Query("UPDATE UrlData SET hits =? where shortUrl = ?", hits + 1  , shortUrl)


}

func UpdateShort(ID uint64, shortUrl string){


	db.Query("UPDATE UrlData SET shortUrl =? where ID = ?", shortUrl, ID)
}


