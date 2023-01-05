package resources

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-fiber/pkg/framework/msg"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/action"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/form"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/grid"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/group"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/icon"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/image"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/layout"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/link"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/list"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/navbar"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/navigator"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/page"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/searchbar"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/section"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/segmentedcontrol"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/swiper"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/video"
	"github.com/quarkcms/quark-fiber/pkg/ui/mix/component/view"
)

// 结构体
type From struct {
	Title     string
	Style     string
	FromStyle string
	Api       string
}

// 初始化
func (p *From) Init() *From {

	return p
}

// 底部导航
func (p *From) TabBar(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// View
func (p *From) View(body interface{}) *view.Component {
	return (&view.Component{}).Init().SetBody(body)
}

// 布局-行
func (p *From) Row(cols []interface{}) *layout.Row {
	return (&layout.Row{}).Init().SetBody(cols)
}

// 布局-列
func (p *From) Col(span int, body interface{}) *layout.Col {
	return (&layout.Col{}).Init().SetSpan(span).SetBody(body)
}

// 列表组件
func (p *From) List(items []interface{}) *list.Component {
	return (&list.Component{}).Init().SetBody(items)
}

// 列表子组件
func (p *From) ListItem(title string, url string) *list.ListItem {
	return (&list.ListItem{}).Init().SetTitle(title).SetTo(url)
}

// 标题栏组件
func (p *From) Section(title string, body interface{}) *section.Component {
	return (&section.Component{}).Init().SetTitle(title).SetBody(body)
}

// 分组组件
func (p *From) Group(title string, body interface{}) *group.Component {
	return (&group.Component{}).Init().SetTitle(title).SetBody(body)
}

// 宫格组件
func (p *From) Grid(column int, body []interface{}) *grid.Component {
	return (&grid.Component{}).Init().SetColumn(column).SetBody(body)
}

// 宫格子组件
func (p *From) GridItem(body interface{}) *grid.GridItem {
	return (&grid.GridItem{}).Init().SetBody(body)
}

// 页面跳转
func (p *From) Navigator(content interface{}, url string) *navigator.Component {
	return (&navigator.Component{}).Init().SetBody(content).SetUrl(url)
}

// 搜索栏
func (p *From) SearchBar() *searchbar.Component {
	return (&searchbar.Component{}).Init()
}

// 图片
func (p *From) Image(src string) *image.Component {
	return (&image.Component{}).
		Init().
		SetSrc(src)
}

// Icon
func (p *From) Icon(iconType string) *icon.Component {
	return (&icon.Component{}).Init().SetType(iconType)
}

// 视频
func (p *From) Video(src string) *video.Component {
	return (&video.Component{}).Init().SetSrc(src)
}

// Link
func (p *From) Link(href string, body interface{}) *link.Component {
	return (&link.Component{}).Init().SetHref(href).SetBody(body)
}

// 分段器
func (p *From) SegmentedControl(titles []interface{}, items []interface{}) *segmentedcontrol.Component {
	return (&segmentedcontrol.Component{}).Init().SetTitles(titles).SetItems(items)
}

// 轮播图
func (p *From) Swiper(items []interface{}) *swiper.Component {
	return (&swiper.Component{}).
		Init().
		SetAutoplay(true).
		SetIndicatorDots(true).
		SetItems(items)
}

// 轮播图子组件
func (p *From) SwiperItem(body interface{}) *swiper.SwiperItem {
	return (&swiper.SwiperItem{}).SetBody(body)
}

// 表单
func (p *From) Form(api string, items []interface{}) *form.Component {
	return (&form.Component{}).
		Init().
		SetApi(api).
		SetBody(items)
}

// 表单项
func (p *From) Field() *form.Field {
	return (&form.Field{})
}

// 表单项
func (p *From) FormItem() *form.Field {
	return (&form.Field{})
}

// 行为
func (p *From) Action(label string, actionType string) *action.Component {
	return (&action.Component{}).
		Init().
		SetLabel(label).
		SetType(actionType)
}

// 头部导航
func (p *From) NavBar(c *fiber.Ctx, resourceInstance interface{}, navbar *navbar.Component) interface{} {
	return nil
}

// 字段
func (p *From) Fields(c *fiber.Ctx, resourceInstance interface{}) []interface{} {
	return nil
}

// 行为
func (p *From) Actions(c *fiber.Ctx, resourceInstance interface{}) []interface{} {
	return []interface{}{
		p.Action("提交", "primary").SetOpenType("submit"),
	}
}

// 表单数据
func (p *From) Data(c *fiber.Ctx, resourceInstance interface{}) map[string]interface{} {
	return nil
}

// 内容
func (p *From) Content(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	fields := resourceInstance.(interface {
		Fields(c *fiber.Ctx, resourceInstance interface{}) []interface{}
	}).Fields(c, resourceInstance)

	data := resourceInstance.(interface {
		Data(c *fiber.Ctx, resourceInstance interface{}) map[string]interface{}
	}).Data(c, resourceInstance)

	actions := resourceInstance.(interface {
		Actions(c *fiber.Ctx, resourceInstance interface{}) []interface{}
	}).Actions(c, resourceInstance)

	value := reflect.ValueOf(resourceInstance).Elem()
	api := value.FieldByName("Api").String()
	fromStyle := value.FieldByName("FromStyle").String()

	if api == "" {
		api = "/api/mix/form/" + strings.ToLower(c.Params("resource")) + "/submit"
	}

	return p.Form(api, fields).
		SetStyle(fromStyle).
		SetModel(data).
		SetActions(actions)
}

// 执行表单
func (p *From) Handle(c *fiber.Ctx, resourceInstance interface{}) error {

	return msg.Error(c, "请自行处理表单逻辑", msg.DEFAULT_URL)
}

// 页面组件渲染
func (p *From) PageComponentRender(c *fiber.Ctx, resourceInstance interface{}) interface{} {
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
func (p *From) Render(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	return resourceInstance.(interface {
		PageComponentRender(c *fiber.Ctx, resourceInstance interface{}) interface{}
	}).PageComponentRender(c, resourceInstance)
}
