package resources

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/form"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/grid"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/group"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/icon"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/image"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/layout"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/link"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/list"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/navbar"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/navigator"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/page"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/searchbar"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/section"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/segmentedcontrol"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/swiper"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/video"
	"github.com/quarkcms/quark-go/pkg/ui/mix/component/view"
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

// 头部导航
func (p *Page) NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{} {
	return nil
}

// 内容
func (p *Page) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// 底部导航
func (p *Page) TabBar(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// View
func (p *Page) View(body interface{}) *view.Component {
	return (&view.Component{}).Init().SetBody(body)
}

// 布局-行
func (p *Page) Row(cols []interface{}) *layout.Row {
	return (&layout.Row{}).Init().SetBody(cols)
}

// 布局-列
func (p *Page) Col(span int, body interface{}) *layout.Col {
	return (&layout.Col{}).Init().SetSpan(span).SetBody(body)
}

// 列表组件
func (p *Page) List(items []interface{}) *list.Component {
	return (&list.Component{}).Init().SetBody(items)
}

// 列表子组件
func (p *Page) ListItem(title string, url string) *list.ListItem {
	return (&list.ListItem{}).Init().SetTitle(title).SetTo(url)
}

// 标题栏组件
func (p *Page) Section(title string, body interface{}) *section.Component {
	return (&section.Component{}).Init().SetTitle(title).SetBody(body)
}

// 分组组件
func (p *Page) Group(title string, body interface{}) *group.Component {
	return (&group.Component{}).Init().SetTitle(title).SetBody(body)
}

// 宫格组件
func (p *Page) Grid(column int, body []interface{}) *grid.Component {
	return (&grid.Component{}).Init().SetColumn(column).SetBody(body)
}

// 宫格子组件
func (p *Page) GridItem(body interface{}) *grid.GridItem {
	return (&grid.GridItem{}).Init().SetBody(body)
}

// 页面跳转
func (p *Page) Navigator(content interface{}, url string) *navigator.Component {
	return (&navigator.Component{}).Init().SetBody(content).SetUrl(url)
}

// 图片
func (p *Page) Image(src string) *image.Component {
	return (&image.Component{}).
		Init().
		SetSrc(src)
}

// Icon
func (p *Page) Icon(iconType string) *icon.Component {
	return (&icon.Component{}).Init().SetType(iconType)
}

// 视频
func (p *Page) Video(src string) *video.Component {
	return (&video.Component{}).Init().SetSrc(src)
}

// Link
func (p *Page) Link(href string, body interface{}) *link.Component {
	return (&link.Component{}).Init().SetHref(href).SetBody(body)
}

// 分段器
func (p *Page) SegmentedControl(titles []interface{}, items []interface{}) *segmentedcontrol.Component {
	return (&segmentedcontrol.Component{}).Init().SetTitles(titles).SetItems(items)
}

// 轮播图
func (p *Page) Swiper(items []interface{}) *swiper.Component {
	return (&swiper.Component{}).
		Init().
		SetAutoplay(true).
		SetIndicatorDots(true).
		SetItems(items)
}

// 轮播图子组件
func (p *Page) SwiperItem(body interface{}) *swiper.SwiperItem {
	return (&swiper.SwiperItem{}).SetBody(body)
}

// 表单
func (p *Page) Form(api string, items []interface{}) *form.Component {
	return (&form.Component{}).
		Init().
		SetApi(api).
		SetBody(items)
}

// 表单项
func (p *Page) Field() *form.Field {
	return (&form.Field{})
}

// 表单项
func (p *Page) FormItem() *form.Field {
	return (&form.Field{})
}

// 行为：label按钮文字，actionType按钮的样式类型 primary | default | warn
func (p *Page) Action(label string, actionType string) *action.Component {
	return (&action.Component{}).
		Init().
		SetLabel(label).
		SetType(actionType)
}

// 搜索栏
func (p *Page) SearchBar() *searchbar.Component {
	return (&searchbar.Component{}).Init()
}

// 页面组件渲染
func (p *Page) PageComponentRender(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	navbarInstance := (&navbar.Component{}).Init()

	navBar := resourceInstance.(interface {
		NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{}
	}).NavBar(c, resourceInstance, navbarInstance)

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
		SetNavBar(navBar).
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
