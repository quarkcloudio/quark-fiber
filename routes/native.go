package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/native/controllers"
	"github.com/quarkcms/quark-go/internal/native/middleware"
	nativecontrollers "github.com/quarkcms/quark-go/pkg/ui/native/http/controllers"
)

type Native struct{}

// Native路由
func (p *Native) Route(app *fiber.App) {

	ag := app.Group("/api/native")
	ag.Get("/captcha", (&nativecontrollers.Captcha{}).Make)
	ag.Get("/page/:resource", (&nativecontrollers.ResourcePage{}).Handle)

	// 图片上传、下载
	ag.Get("/picture/getLists", (&nativecontrollers.Picture{}).GetLists)
	ag.Post("/picture/upload", (&nativecontrollers.Picture{}).Upload)
	ag.Get("/picture/download", (&nativecontrollers.Picture{}).Download)
	ag.All("/picture/delete", (&nativecontrollers.Picture{}).Delete)
	ag.Post("/picture/crop", (&nativecontrollers.Picture{}).Crop)

	// 文件上传、下载
	ag.Post("/file/upload", (&nativecontrollers.File{}).Upload)
	ag.Get("/file/download", (&nativecontrollers.File{}).Download)

	amg := app.Group("/api/native", (&middleware.ApiMiddleware{}).Handle)
	amg.Get("/user/index", (&controllers.User{}).Index)
}
