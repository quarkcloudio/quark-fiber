package metrics

import (
	"github.com/quarkcms/quark-fiber/internal/models"
	"github.com/quarkcms/quark-fiber/pkg/framework/db"
	"github.com/quarkcms/quark-fiber/pkg/ui/admin/component/statistic"
	"github.com/quarkcms/quark-fiber/pkg/ui/admin/metrics"
)

type TotalPicture struct {
	metrics.Value
}

// 初始化
func (p *TotalPicture) Init() *TotalPicture {
	p.Title = "图片数量"
	p.Col = 6

	return p
}

// 计算数值
func (p *TotalPicture) Calculate() *statistic.Component {

	return p.
		Init().
		Count((&db.Model{}).Model(&models.Picture{})).
		SetValueStyle(map[string]string{"color": "#cf1322"})
}
