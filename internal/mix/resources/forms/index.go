package forms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/navbar"
	"github.com/quarkcms/quark-go/pkg/ui/mix/resources"
)

type Index struct {
	resources.From
}

// 初始化
func (p *Index) Init() interface{} {

	// 标题
	p.Title = "表单页"

	// 表单样式
	p.FromStyle = "margin:20px;background:#ffffff;padding:10px;border-radius:10px;"

	return p
}

// 头部导航
func (p *Index) NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{} {

	navbar.SetTitle(p.Title).SetLeftIcon("left")

	return navbar
}

// 字段
func (p *Index) Fields(c *fiber.Ctx, resourceInstance interface{}) []interface{} {

	return []interface{}{
		p.Field().EasyInput("username", "用户名"),
		p.Field().Radio("sex", "性别").SetOptions([]map[string]interface{}{
			{
				"text":  "男",
				"value": 1,
			},
			{
				"text":  "女",
				"value": 2,
			},
		}).SetValue(1),
		p.Field().DatetimePicker("created_at", "创建日期"),
		p.Field().Switch("status", "状态").SetValue(true),
	}
}

// 执行表单
func (p *Index) Handle(c *fiber.Ctx, resourceInstance interface{}) error {

	return msg.Success(c, "提交成功！", msg.DEFAULT_URL, msg.DEFAULT_DATA)
}
