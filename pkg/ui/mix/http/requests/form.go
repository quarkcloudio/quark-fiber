package requests

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/mix"
)

type Form struct{}

// Form资源
func (p *Form) Resource(c *fiber.Ctx) interface{} {
	var resourceInstance interface{}

	for _, provider := range mix.Providers {
		providerName := reflect.TypeOf(provider).String()

		// 包含字符串
		if find := strings.Contains(providerName, "*forms."); find {
			structName := strings.Replace(providerName, "*forms.", "", -1)
			if strings.ToLower(structName) == strings.ToLower(c.Params("resource")) {

				// 初始化实例
				resourceInstance = provider.(interface{ Init() interface{} }).Init()
			}
		}
	}

	return resourceInstance
}
