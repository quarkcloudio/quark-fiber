package requests

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/native"
)

type Page struct{}

// Page资源
func (p *Page) Resource(c *fiber.Ctx) interface{} {
	var resourceInstance interface{}

	for _, provider := range native.Providers {
		providerName := reflect.TypeOf(provider).String()

		// 包含字符串
		if find := strings.Contains(providerName, "*pages."); find {
			structName := strings.Replace(providerName, "*pages.", "", -1)
			if strings.ToLower(structName) == strings.ToLower(c.Params("resource")) {

				// 初始化实例
				resourceInstance = provider.(interface{ Init() interface{} }).Init()
			}
		}
	}

	return resourceInstance
}
