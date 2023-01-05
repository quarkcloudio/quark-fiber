package fields

import "github.com/quarkcms/quark-fiber/pkg/ui/admin/component"

type Month struct {
	Item
}

// 初始化
func (p *Month) Init() *Month {
	p.Component = "monthField"
	p.InitItem().SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}
