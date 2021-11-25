package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func main(){
	dnsName := "localhost"
	cnnstr := fmt.Sprintf("user=%s password=%s host=%s port=5432 connect_timeout=3 dbname=postgres sslmode=disable", "postgres","12345", dnsName)
	eng, err := xorm.NewEngine("postgres", cnnstr)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("choose any 'create' 'alter' 'delete' 'list' 'grant' 'stop'")
		var choice string
		fmt.Scanf("%s",&choice)
		if choice == "create" {
			fmt.Println("Create User")
			Createuser(eng)
		} else if choice == "grant" {
			fmt.Println("Grant user")
			GrantDb(eng)
		} else if choice == "list" {
			fmt.Println("List db")
			ListDb(eng)
		} else if choice == "alter" {
			fmt.Println("Alter db")
			AlterDb(eng)
		} else if choice == "delete" {
			fmt.Println("Delete db")
			DeleteDb(eng)
		} else {
			break
		}

		
		//var cache string
		//fmt.Scanf("%s",cache)
	}
	defer eng.Close()
}