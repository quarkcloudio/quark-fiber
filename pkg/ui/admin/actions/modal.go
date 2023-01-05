package actions

import "github.com/gofiber/fiber/v2"

type Modal struct {
	Action
	Width          int  `json:"width"`
	DestroyOnClose bool `json:"destroyOnClose"`
}

// 初始化
func (p *Modal) ParentInit() interface{} {
	p.ActionType = "modal"
	p.Width = 520

	return p
}

// 宽度
func (p *Modal) GetWidth() int {
	return p.Width
}

// 关闭时销毁 Modal 里的子元素
func (p *Modal) GetDestroyOnClose() bool {
	return p.DestroyOnClose
}

// 内容
func (p *Modal) GetBody(c *fiber.Ctx, resourceInstance interface{}) interface{} {
	return nil
}

// 弹窗行为
func (p *Modal) GetActions(c *fiber.Ctx, resourceInstance interface{}) []interface{} {
	return []interface{}{}
}
