package demo_service

import (
	"context"
	"encoding/json"
	"example/models/entities"
	"example/repositories"
	"time"

	"github.com/go-redis/redis/v8"
)

type demoServiceRedis struct {
	Repo        repositories.IDemoRepository
	RedisClient *redis.Client
}

func NewDemoServiceRedis(repo repositories.IDemoRepository, redisClient *redis.Client) demoServiceRedis {
	return demoServiceRedis{Repo: repo, RedisClient: redisClient}
}

func (s demoServiceRedis) GetDemo() (demo []entities.Demo, err error) {

	key := "service::GetDemo"

	if demoJson, err := s.RedisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(demoJson), &demo) == nil {
			return demo, nil
		}
	}

	demoDB, err := s.Repo.GetDemo()
	if err != nil {
		return nil, err
	}

	for _, d := range demoDB {
		demo = append(demo, entities.Demo{
			ID:       d.ID,
			Name:     d.Name,
			Priority: d.Priority,
		})
	}

	if data, err := json.Marshal(demo); err == nil {
		s.RedisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return demo, nil
}
