package main

import (
	"fmt"
	"xorm.io/xorm"
)

type User struct{
	Name string `yaml: "username"`
	Password string `yaml: "password"`
	Permissions []string `yaml : "userpermissions"`
}

func Createuser(eng *xorm.Engine) {
	var user User;
	fmt.Println("UserName:")
	fmt.Scanf("%s",&user.Name)
	fmt.Println("Password:")
	fmt.Scanf("%s",&user.Password)
	queryString := fmt.Sprintf("create role %s login password '%s';",user.Name,user.Password)
	_, err := eng.QueryString(queryString)
	if err != nil {
		panic(err)
	}
}