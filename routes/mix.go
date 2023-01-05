package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/mix/controllers"
	"github.com/quarkcms/quark-go/internal/mix/middleware"
	mixcontrollers "github.com/quarkcms/quark-go/pkg/ui/mix/http/controllers"
)

type Mix struct{}

// mix路由
func (p *Mix) Route(app *fiber.App) {

	ag := app.Group("/api/mix")
	ag.Get("/captcha", (&mixcontrollers.Captcha{}).Make)
	ag.Get("/page/:resource", (&mixcontrollers.ResourcePage{}).Handle)
	ag.Get("/form/:resource", (&mixcontrollers.ResourceForm{}).Handle)
	ag.All("/form/:resource/submit", (&mixcontrollers.ResourceForm{}).Submit)

	// 图片上传、下载
	ag.Get("/picture/getLists", (&mixcontrollers.Picture{}).GetLists)
	ag.Post("/picture/upload", (&mixcontrollers.Picture{}).Upload)
	ag.Get("/picture/download", (&mixcontrollers.Picture{}).Download)
	ag.All("/picture/delete", (&mixcontrollers.Picture{}).Delete)
	ag.Post("/picture/crop", (&mixcontrollers.Picture{}).Crop)

	// 文件上传、下载
	ag.Post("/file/upload", (&mixcontrollers.File{}).Upload)
	ag.Get("/file/download", (&mixcontrollers.File{}).Download)

	amg := app.Group("/api/mix", (&middleware.ApiMiddleware{}).Handle)
	amg.Get("/user/index", (&controllers.User{}).Index)
	amg.All("/index/index", (&controllers.Index{}).Index)
}
