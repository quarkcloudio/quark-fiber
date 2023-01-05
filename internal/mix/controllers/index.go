package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
)

type Index struct{}

// Index
func (p *Index) Index(c *fiber.Ctx) error {

	return msg.Success(c, "Hello World", msg.DEFAULT_URL, msg.DEFAULT_DATA)
}
