package main

import (
	"html/template"
	"os"
	"time"
	_ "upyun_sync/routers"

	"github.com/astaxie/beego"
)

func convertTypetoIcon(in string) (out string) {
	if in == "folder" {
		return "glyphicon glyphicon-folder-close"
	}
	return "glyphicon glyphicon-file"
}

func converTime(in int64) (out string) {
	ti := time.Unix(in, 0)
	return ti.Format("Mon Jan 2 15:04:05 -0700 MST 2006")

}

func genLink(filetype string, filename string, currentpath string) (out template.HTML) {
	if currentpath[len(currentpath)-1] != os.PathSeparator {
		currentpath += string(os.PathSeparator)
	}
	if filetype == "folder" {
		return template.HTML("<a href=\"/index?path=" + currentpath + filename + "\" >" + filename + "</a>")
	}
	return template.HTML(filename)
}
func main() {
	beego.AddFuncMap("fileConv", convertTypetoIcon)
	beego.AddFuncMap("genLink", genLink)
	beego.AddFuncMap("converTime", converTime)
	beego.Run()
}
