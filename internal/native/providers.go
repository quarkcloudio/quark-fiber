package native

import (
	"github.com/quarkcms/quark-go/internal/native/resources/pages"
)

// 注册服务
var Providers = []interface{}{
	&pages.Index{},
	&pages.My{},
}
