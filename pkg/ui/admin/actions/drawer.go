package actions

import "github.com/gofiber/fiber/v2"

type Drawer struct {
	Action
	Width          int  `json:"width"`
	DestroyOnClose bool `json:"destroyOnClose"`
}

// 初始化
func (p *Drawer) ParentInit() interface{} {
	p.ActionType = "drawer"
	p.Width = 520

	return p
}

// 宽度
func (p *Drawer) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *Drawer) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 内容
func (p *Drawer) GetBody(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// 弹窗行为
func (p *Drawer) GetActions(c *fiber.Ctx, resourceInstance interface{}) []interface{} {
	return []interface{}{}
}
