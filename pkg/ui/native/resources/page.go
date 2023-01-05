package resources

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/ui/native/component/navigator"
	"github.com/quarkcms/quark-fiber/pkg/ui/native/component/page"
	"github.com/quarkcms/quark-fiber/pkg/ui/native/component/view"
)

// 结构体
type Page struct {
	Title string
	Style string
}

// 初始化
func (p *Page) Init() *Page {

	return p
}

// 内容
func (p *Page) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// 页面跳转
func (p *Page) Navigator(content interface{}, url string) *navigator.Component {
	return (&navigator.Component{}).Init().SetBody(content).SetUrl(url)
}

// 底部导航
func (p *Page) TabBar(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// View
func (p *Page) View(body interface{}) *view.Component {
	return (&view.Component{}).Init().SetBody(body)
}

// 页面组件渲染
func (p *Page) PageComponentRender(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	content := resourceInstance.(interface {
		Content(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).Content(c, resourceInstance)

	tabBar := resourceInstance.(interface {
		TabBar(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).TabBar(c, resourceInstance)

	value := reflect.ValueOf(resourceInstance).Elem()
	title := value.FieldByName("Title").String()
	style := value.FieldByName("Style").String()

	return (&page.Component{}).
		Init().
		SetTitle(title).
		SetContent(content).
		SetTabBar(tabBar).
		SetStyle(style).
		JsonSerialize()
}

// 组件渲染
func (p *Page) Render(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	return resourceInstance.(interface {
		PageComponentRender(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).PageComponentRender(c, resourceInstance)
}
