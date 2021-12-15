package router

import (
	demo_handler "example/delivery/http/demo"

	"github.com/gofiber/fiber/v2"
)

type router struct {
	App         *fiber.App
	Router      fiber.Router
	DemoHandler demo_handler.IDemoHandler
}

var instantiated *router = nil

func New(
	app *fiber.App,
	demoHandler demo_handler.IDemoHandler,
) *router {
	if instantiated == nil {
		instantiated = &router{
			App:         app,
			Router:      app.Group("api/v1"),
			DemoHandler: demoHandler,
		}
		instantiated.init()
	}
	return instantiated
}

func (r *router) init() {
	r.setupDemo()
}

func (r *router) setupDemo() {
	group := r.Router.Group("demo")
	{
		group.Get("/list", r.DemoHandler.GetDemo)
	}
}
