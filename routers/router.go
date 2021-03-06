package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

var FilterUser = func(ctx *context.Context) {
	//username, ok := ctx.Input.Session("username").(string)
	session := controllers.GetSession()
	sess, _ := session.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)
	username, ok := sess.Get("username").(string)
	if !ok && ctx.Request.RequestURI != "/school/login" && ctx.Request.RequestURI!="/school/toGrade" {
		//ctx.Redirect(302, "/school/login")
	}
	fmt.Println("用户名:" + username)
}

func init() {
	beego.Router("/school/login", &controllers.MainController{}, "get:Login;post:Craw")
	beego.Router("/school/toCredit", &controllers.MainController{}, "get:ToCredit;post:QueryCredit")
	beego.Router("/school/toEvaluate", &controllers.MainController{}, "get:ToEvaluate;post:Evaluate")
	beego.Router("/checkCode", &controllers.MainController{}, "get:CheckCode")
	beego.Router("/getCodeUrl", &controllers.MainController{}, "get:GetCodeUrl")
	beego.Router("/school/toGrade", &controllers.MainController{}, "get:ToGrade;post:QueryGrade")
	beego.Router("/school/cet", &controllers.MainController{}, "get:Cet;post:GetCetGrade")
	beego.InsertFilter("/school/*", beego.BeforeRouter, FilterUser)
}
