package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/http/requests"
)

type ResourceForm struct{}

// Form页
func (p *ResourceForm) Handle(c *fiber.Ctx) error {

	resource := &requests.Form{}

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

// Form提交方法
func (p *ResourceForm) Submit(c *fiber.Ctx) error {

	resource := &requests.Form{}

	// 资源实例
	resourceInstance := resource.Resource(c)

	if resourceInstance == nil {
		return c.SendStatus(404)
	}

	// 断言Handle方法
	return resourceInstance.(interface {
		Handle(*fiber.Ctx, interface{}) error
	}).Handle(c, resourceInstance)
}
