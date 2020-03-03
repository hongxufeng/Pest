package main

import (
	"datahelper/db"
	"fmt"
	"module/imports"
	"module/upload"
	"module/user"
	"os"
	"utils/config"
	"utils/service"
)

var conf config.Config

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("invalid args : %s [config]\n", os.Args[0])
		return
	}
	err := conf.LoadPath(os.Args[1])
	if err != nil {
		fmt.Println("Load Config error :", err.Error())
		return
	}
	fmt.Println("begin init db")
	e := db.Init(&conf)
	if e == nil {
		fmt.Println("init db ok")
	} else {
		fmt.Println(e.Error())
		return
	}
	//fmt.Printf("\n%v\n\n", conf)
	server, err := service.New(&conf, 2, false, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	server.AddModule("user", &user.UserModule{})
	server.AddModule("upload", &upload.UploadModule{})
	server.AddModule("imports", &imports.ImportsModule{})
	fmt.Println(server.StartService())
}
