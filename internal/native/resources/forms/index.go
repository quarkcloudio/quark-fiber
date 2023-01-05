package forms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/ui/native/resources"
)

type Index struct {
	resources.Page
}

// 初始化
func (p *Index) Init() interface{} {

	// 标题
	p.Title = "表单页"

	return p
}

// 内容
func (p *Index) Body(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return "这里是表单页"
}
