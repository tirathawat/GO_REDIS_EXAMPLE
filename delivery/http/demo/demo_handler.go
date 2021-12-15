package demo_handler

import (
	"example/models/response"
	demo_service "example/services/demo"

	"github.com/gofiber/fiber/v2"
)

type demoHandler struct {
	Service demo_service.IDemoService
}

func NewDemoHandler(service demo_service.IDemoService) demoHandler {
	return demoHandler{Service: service}
}

func (h demoHandler) GetDemo(c *fiber.Ctx) error {

	demo, err := h.Service.GetDemo()
	if err != nil {
		return err
	}

	response := response.DemoResponse{
		Status: "ok",
		Demo:   demo,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
