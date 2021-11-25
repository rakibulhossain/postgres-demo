package main

import (
	"fmt"
	"xorm.io/xorm"
)

type Database struct {
	DbName string `yam : "dbname"`
	Username string `yaml: "dbusername"`
	Permissions []string `yaml: "dbpermissions"`
}

func GrantDb(eng *xorm.Engine) {
	var db Database
	fmt.Println("Dbname:")
	fmt.Scanf("%s",&db.DbName)
	fmt.Println("Username:")
	fmt.Scanf("%s",&db.Username)
	fmt.Println("Permissions:")
	for {
		var s string
		fmt.Scanf("%s",&s);
		if s == "stop" {
			break
		}
		db.Permissions = append(db.Permissions, s)
	}
	querystring := fmt.Sprintf("select %s from pg_database;",db.DbName)
	_ , err := eng.QueryString(querystring)

	if err != nil {
		fmt.Println(1)
		querystring = fmt.Sprintf("create database %s;",db.DbName)
		_, err = eng.QueryString(querystring)
		if err != nil {
			panic(err)
		}
	}
	var permissions string
	for _, perm := range db.Permissions {
		permissions+=" "+perm
	}
	querystring = fmt.Sprintf("grant %s on database %s to %s;",permissions,db.DbName,db.Username)
	_ ,err = eng.QueryString(querystring)
	if err != nil {
		panic(err)
	}
}

func AlterDb(eng *xorm.Engine){
	fmt.Println("Dbname:")
	var dbname string
	fmt.Scanf("%s",&dbname)
	fmt.Println("newname:")
	var newName string
	fmt.Scanf("%s",&newName)
	querystring := fmt.Sprintf("Alter database %s rename to %s;",dbname,newName);
	_ ,err := eng.QueryString(querystring)
	if err != nil {
		panic(err)
	}
}

func DeleteDb(eng *xorm.Engine){
	fmt.Println("Dbname:")
	var dbname string
	fmt.Scanf("%s",&dbname)
	querystring := fmt.Sprintf("drop database %s;",dbname);
	_ ,err := eng.QueryString(querystring)
	if err != nil {
		panic(err)
	}
}

func ListDb(eng *xorm.Engine){
	querystring := fmt.Sprintf("select datname from pg_database;")
	res ,err := eng.QueryString(querystring)
	if err != nil {
		panic(err)
	}
	for _ , element := range res {
		for _, v := range element {
			fmt.Println(v)
		}
	}
}