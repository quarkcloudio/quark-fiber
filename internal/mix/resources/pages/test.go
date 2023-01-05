package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/navbar"
	"github.com/quarkcms/quark-go/pkg/ui/mix/resources"
)

type Test struct {
	resources.Page
}

// 初始化
func (p *Test) Init() interface{} {

	// 标题
	p.Title = "测试页"

	return p
}

// 头部导航
func (p *Test) NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{} {

	navbar.SetTitle(p.Title).SetLeftIcon("left")

	return navbar
}

// 内容
func (p *Test) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	var actions []interface{}

	actions = append(actions, p.Action("抽屉", "primary").
		SetOpenType("drawer").
		SetDrawer(func(drawer *action.Drawer) interface{} {
			return drawer.SetBody("ABC")
		}))

	return p.View(actions).SetStyle("padding:20px")
}
