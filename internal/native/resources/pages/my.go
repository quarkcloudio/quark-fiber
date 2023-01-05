package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/ui/native/resources"
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

// 内容
func (p *My) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	return p.View("我的").SetStyle(map[string]interface{}{
		"color":    "red",
		"fontSize": 28,
	})
}
