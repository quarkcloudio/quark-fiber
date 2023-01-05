package session

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
	session "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/quarkcms/quark-go/pkg/framework/config"
)

var store = session.New()

// 初始化
func init() {

	store = session.New(session.Config{
		Expiration:     config.Get("session.expiration").(time.Duration),
		Storage:        config.Get("session.storage").(fiber.Storage),
		KeyLookup:      config.Get("session.key_lookup").(string),
		CookieDomain:   config.Get("session.cookie_domain").(string),
		CookiePath:     config.Get("session.cookie_path").(string),
		CookieSecure:   config.Get("session.cookie_secure").(bool),
		CookieHTTPOnly: config.Get("session.cookie_http_only").(bool),
		CookieSameSite: config.Get("session.cookie_same_site").(string),
		KeyGenerator:   config.Get("session.key_generator").(func() string),
	})
}

// 设置值
func Set(c *fiber.Ctx, key string, value interface{}) error {

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set(key, value)

	return sess.Save()
}

// 获取值
func Get(c *fiber.Ctx, key string) interface{} {

	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	result := sess.Get(key)

	return result
}
