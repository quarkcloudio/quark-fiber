package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/native/resources"
)

type Index struct {
	resources.Page
}

// 初始化
func (p *Index) Init() interface{} {

	// 标题
	p.Title = "首页"

	return p
}

// 内容
func (p *Index) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	return p.View(
		p.
			Navigator("跳转到测试页", "pages/test/index").
			SetOpenType("redirect"),
	).SetStyle(map[string]interface{}{
		"color":    "blue",
		"fontSize": 28,
	})
}
