package main

import (
	"os"
	_"github.com/joaquinicolas/FormAndroidMigrator/DataBase"
	"github.com/joaquinicolas/FormAndroidMigrator/libs"
	"fmt"
	"github.com/joaquinicolas/FormAndroidMigrator/DataBase"
)

func main()  {
	if(os.Getenv("DATABASE_URL") == ""){
		os.Setenv("DATABASE_URL","postgres://postgres:teamanalysis@ncrapps.com:5432/formcloud")
	}
	store,err := DataBase.CreateDataStore(map[string]string{
		"DATASTORE":"postgres",
		"DATASTORE_POSTGRES_DSN":os.Getenv("DATABASE_URL"),
	})

	checkErr(err)
	csv_formatter,err := libs.CreateFormatter(map[string]string{
		"FORMAT":"CSV",
		"LOCATION":os.Args[1],
		"SEPARATOR":",",
	})
fmt.Println(csv_formatter)
	checkErr(err)
	client := libs.NewClient("","","")
	fmt.Println(client)
	client.CreateFromFormatter(csv_formatter)(store)
}



func checkErr(e error)  {
	if e != nil{
		panic(e)
	}
}