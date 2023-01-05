package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/navbar"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/resources"
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

// 头部导航
func (p *Index) NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{} {

	return nil
}

// 内容
func (p *Index) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	var content []interface{}

	searchBar := p.SearchBar().SetStyle("margin:0px 10px")
	content = append(content, searchBar)

	imageStyle := "width: 100%; height: 200px;"
	swiper := p.Swiper([]interface{}{
		p.SwiperItem(
			p.Image("https://storage.360buyimg.com/jdc-article/NutUItaro2.jpg").
				SetStyle(imageStyle),
		),
		p.SwiperItem(
			p.Image("https://storage.360buyimg.com/jdc-article/NutUItaro34.jpg").
				SetStyle(imageStyle),
		),
		p.SwiperItem(
			p.Image("https://storage.360buyimg.com/jdc-article/fristfabu.jpg").
				SetStyle(imageStyle),
		),
	}).SetItemStyle("border-radius:10px;overflow:hidden;").SetStyle("margin:0px 20px;")

	content = append(content, swiper)

	grid := p.Grid(4, []interface{}{
		p.GridItem(p.Icon("color")),
		p.GridItem(p.Icon("wallet")),
		p.GridItem(p.Icon("shop-filled")),
		p.GridItem(p.Icon("location")),
	}).SetStyle("text-align:center;margin:20px 0px;")

	content = append(content, grid)

	list := p.List([]interface{}{
		p.ListItem("页面", "/pages/engine/withoutNavBar?api=/api/mix/page/test"),
		p.ListItem("表单", "/pages/engine/withoutNavBar?api=/api/mix/form/index"),
	})
	section := p.Section("标题", list).SetType("line").SetStyle("margin:0px 20px;")

	content = append(content, section)

	return content
}
