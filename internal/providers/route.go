package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/routes"
)

// 结构体
type Route struct{}

// 注册服务
func (p *Route) Register(app *fiber.App) {

	// 注册Admin路由
	(&routes.Admin{}).Route(app)

	// 注册Mix路由
	(&routes.Mix{}).Route(app)

	// 注册Native路由
	(&routes.Native{}).Route(app)

	// 注册Web路由
	(&routes.Web{}).Route(app)
}
