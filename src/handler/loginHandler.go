package handler

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

<<<<<<< HEAD
func (this *LoginController) POST() {
	username := this.GetStrings("username")
	password := this.GetStrings("password")


=======
func (this *LoginController) Get() {
>>>>>>> 758e9a85b04aff84e5a32f825f4d8ac01b320c5f
	this.Ctx.WriteString("hello world")
}