package controllers

import (
	"fmt"
	"upyun_sync/upyun"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	u := upyun.NewUpYun("backup", "zerob13", "www.zerob13.in")

	u.Debug = true
	u.SetApiDomain(upyun.EdAuto)
	v, err := u.GetBucketUsage()
	if err != nil {
		fmt.Println(err)
		return
	}
	dirs, err := u.ReadDir("/")
	fmt.Printf("ReadDir: %v\n", err)
	for i, d := range dirs {
		fmt.Printf("\t%d: %v\n", i, d)
	}

	this.Data["UsedSize"] = v / 1024 / 1024
	files, err := u.ReadDir("/")
	if err != nil {
		fmt.Println(err)
		return
	}
	this.Data["Files"] = files
	this.TplNames = "index.html"
}
