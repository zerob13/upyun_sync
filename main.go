package main

import (
	"html/template"
	"os"
	"path/filepath"
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
	return template.HTML(filename + "   <a href=\"/delete?dFilename=" + currentpath + filename + "\" class=\"btn-sm btn-danger pull-right\" >删除</a>")
}
func genBackBtn(currentpath string) (out template.HTML) {
	if currentpath != "/" {
		subpath := beego.Substr(currentpath, 0, len(currentpath)-len(filepath.Base(currentpath))-1)

		return template.HTML("<a href=\"/index?path=" + subpath + "\"  class=\"btn-sm btn-primary \">后退</a>")
	}
	return

}
func main() {
	beego.AddFuncMap("fileConv", convertTypetoIcon)
	beego.AddFuncMap("genBackBtn", genBackBtn)
	beego.AddFuncMap("genLink", genLink)
	beego.AddFuncMap("converTime", converTime)
	beego.Run()
}
