package controllers

import (
	"fmt"
	"os"
	"reflect"
	"upyun_sync/models"
	"upyun_sync/upyun"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {

	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/login")
		return
	}
	fmt.Println(reflect.TypeOf(userinfo))
	user := new(models.Space)
	//Magic code:: reflect to get all field
	user.Name = reflect.ValueOf(userinfo).Elem().Field(0).String()
	user.UserName = reflect.ValueOf(userinfo).Elem().Field(1).String()
	user.PassWord = reflect.ValueOf(userinfo).Elem().Field(2).String()
	u := upyun.NewUpYun(user.Name, user.UserName, user.PassWord)
	u.Debug = false
	u.SetApiDomain(upyun.EdAuto)
	v, err := u.GetBucketUsage()
	if err != nil {
		fmt.Println(err)
		return
	}
	filepath := this.GetString("path")
	if filepath == "" {
		filepath = "/"
	}
	dirs, err := u.ReadDir("filepath")
	fmt.Printf("ReadDir: %v\n", err)
	for i, d := range dirs {
		fmt.Printf("\t%d: %v\n", i, d)
	}
	currentPath := filepath
	this.Data["UsedSize"] = v / 1024 / 1024
	files, err := u.ReadDir(currentPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(files))
	if files[0].Name != "" {
		this.Data["Files"] = files
	}
	if currentPath[len(currentPath)-1] != os.PathSeparator {
		currentPath += string(os.PathSeparator)
	}
	this.Data["User"] = user
	this.Data["currentPath"] = currentPath
	this.TplNames = "index.html"
}

func (this *MainController) Login() {
	userinfo := this.GetSession("userinfo")
	// fmt.Println(userinfo)
	if userinfo != nil {
		this.Ctx.Redirect(302, "/index")
	}
	this.TplNames = "login.html"
	spacename := this.GetString("spacename")
	username := this.GetString("username")
	fmt.Println(username)
	if username != "" && spacename != "" {
		password := this.GetString("password")
		user := new(models.Space)
		user.Name = spacename
		user.UserName = username
		user.PassWord = password
		this.SetSession("userinfo", user)
		this.Ctx.Redirect(302, "/index")
		return
	}
	return
}
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/login")
}

func (this *MainController) Upload() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/login")
	}
	// filename := this.GetString("filePath")
	file, fileheader, _ := this.GetFile("filePath")
	filename := fileheader.Filename
	file.Close()
	targetpath := this.GetString("targetPath")
	if targetpath[len(targetpath)-1] != os.PathSeparator {
		targetpath += string(os.PathSeparator)
	}
	target := targetpath + filename
	fmt.Println(target)
	user := new(models.Space)
	//Magic code:: reflect to get all field
	user.Name = reflect.ValueOf(userinfo).Elem().Field(0).String()
	user.UserName = reflect.ValueOf(userinfo).Elem().Field(1).String()
	user.PassWord = reflect.ValueOf(userinfo).Elem().Field(2).String()
	u := upyun.NewUpYun(user.Name, user.UserName, user.PassWord)
	u.Debug = true
	u.SetApiDomain(upyun.EdAuto)
	tofile := "./tmp/" + user.UserName + "_" + filename
	err := this.SaveToFile("filePath", tofile)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Redirect(302, "/index")
		return
	}

	tmp, err := os.Open(tofile)
	if err != nil {
		fmt.Println(err)
		this.Ctx.Redirect(302, "/index")
		return
	}
	defer os.Remove(tofile)
	defer tmp.Close()
	u.SetContentMD5(upyun.FileMd5(tofile))
	fmt.Printf("WriteFile: %v\n", u.WriteFile(target, tmp, true))

	this.Ctx.Redirect(302, "/index")
}

func (this *MainController) Delete() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, "/login")
	}
	fileName := this.GetString("dFilename")
	user := new(models.Space)
	//Magic code:: reflect to get all field
	user.Name = reflect.ValueOf(userinfo).Elem().Field(0).String()
	user.UserName = reflect.ValueOf(userinfo).Elem().Field(1).String()
	user.PassWord = reflect.ValueOf(userinfo).Elem().Field(2).String()
	u := upyun.NewUpYun(user.Name, user.UserName, user.PassWord)
	u.Debug = true
	u.SetApiDomain(upyun.EdAuto)
	fmt.Printf("RmDir: %v\n", u.RmDir(fileName))
	this.Ctx.Redirect(302, "/index")

}
