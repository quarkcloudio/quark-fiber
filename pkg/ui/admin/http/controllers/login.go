package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/config"
	"github.com/quarkcms/quark-go/pkg/framework/hash"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/framework/token"
	"github.com/quarkcms/quark-go/pkg/ui/admin/component/login"
)

type Login struct{}

// 请求结构体
type Request struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Captcha  string `json:"captcha" form:"captcha"`
}

// 登录页面
func (p *Login) Show(c *fiber.Ctx) error {

	loginComponent := &login.Component{}

	component := loginComponent.
		SetApi("admin/login").
		SetRedirect("/index?api=admin/dashboard/index").
		SetLogo(config.Get("admin.logo")).
		SetTitle(config.Get("admin.name").(string)).
		SetDescription(config.Get("admin.description").(string)).
		SetCaptchaUrl("/api/admin/captcha").
		SetCopyright(config.Get("admin.copyright").(string)).
		SetLinks(config.Get("admin.links").([]map[string]interface{})).
		JsonSerialize()

	return c.JSON(component)
}

// 登录方法
func (p *Login) Login(c *fiber.Ctx) error {
	request := new(Request)

	if err := c.BodyParser(request); err != nil {
		return err
	}

	if !(&Captcha{}).Check(c, request.Captcha) {
		return msg.Error(c, "验证码错误", msg.DEFAULT_URL)
	}

	if request.Username == "" || request.Password == "" {
		return msg.Error(c, "用户名或密码不能为空", msg.DEFAULT_URL)
	}

	model := &models.Admin{}
	admin := model.GetAdminViaUsername(request.Username)

	// 检验账号和密码
	if !hash.Check(admin.Password, request.Password) {
		return msg.Error(c, "用户名或密码错误", msg.DEFAULT_URL)
	}

	data := make(map[string]interface{})
	data["id"] = admin.Id
	data["avatar"] = admin.Avatar
	data["nickname"] = admin.Nickname
	data["username"] = admin.Username
	data["guard_name"] = "admin"

	// 创建token
	getToken, _ := token.Make(data)
	data["token"] = getToken

	// 更新最后登录信息
	model.UpdateLastLogin(admin.Id, c.IP(), time.Now())

	return msg.Success(c, "登录成功", msg.DEFAULT_URL, data)
}

// 用户退出方法
func (p *Login) Logout(c *fiber.Ctx) error {

	return msg.Success(c, "已退出", msg.DEFAULT_URL, msg.DEFAULT_DATA)
}
