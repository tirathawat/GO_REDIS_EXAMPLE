package demo_handler

import (
	"context"
	"encoding/json"
	"example/models/response"
	demo_service "example/services/demo"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type demoHandlerRedis struct {
	Service     demo_service.IDemoService
	RedisClient *redis.Client
}

func NewDemoHandlerRedis(service demo_service.IDemoService, redisClient *redis.Client) demoHandlerRedis {
	return demoHandlerRedis{Service: service, RedisClient: redisClient}
}

func (h demoHandlerRedis) GetDemo(c *fiber.Ctx) error {

	key := "handler::GetDemo"

	if responseJson, err := h.RedisClient.Get(context.Background(), key).Result(); err == nil {
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	demo, err := h.Service.GetDemo()
	if err != nil {
		return err
	}

	response := response.DemoResponse{
		Status: "ok",
		Demo:   demo,
	}

	if data, err := json.Marshal(response); err == nil {
		h.RedisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
