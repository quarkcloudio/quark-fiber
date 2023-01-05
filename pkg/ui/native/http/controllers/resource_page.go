package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/native/http/requests"
)

type ResourcePage struct{}

// Page页
func (p *ResourcePage) Handle(c *fiber.Ctx) error {

	resource := &requests.Page{}

	// 资源实例
	resourceInstance := resource.Resource(c)

	if resourceInstance == nil {
		return c.SendStatus(404)
	}

	// 断言Render方法
	component := resourceInstance.(interface {
		Render(*fiber.Ctx, interface{}) interface{}
	}).Render(c, resourceInstance)

	return c.JSON(component)
}
