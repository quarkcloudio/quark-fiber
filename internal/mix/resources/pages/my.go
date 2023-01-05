package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/navbar"
	"github.com/quarkcms/quark-go/pkg/ui/mix/resources"
)

type My struct {
	resources.Page
}

// 初始化
func (p *My) Init() interface{} {

	// 标题
	p.Title = "我的"

	return p
}

// 头部导航
func (p *My) NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{} {

	return nil
}

// 内容
func (p *My) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	return p.View("我的").SetStyle("padding:20px")
}
