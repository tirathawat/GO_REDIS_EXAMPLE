package repositories

import (
	"context"
	"encoding/json"
	"example/database"
	"example/models/storages"
	"time"

	"github.com/go-redis/redis/v8"
)

type demoRepositoryRedis struct {
	Database    database.IDatabase
	RedisClient *redis.Client
}

func NewDemoRepositoryRedis(db database.IDatabase, redisClient *redis.Client) demoRepositoryRedis {
	db.GetDB().AutoMigrate(storages.DemoDB{})
	db.FeedData()
	return demoRepositoryRedis{Database: db, RedisClient: redisClient}
}

func (r demoRepositoryRedis) GetDemo() (demo []storages.DemoDB, err error) {
	key := "repository::GetDemo"

	if demoJson, err := r.RedisClient.Get(context.Background(), key).Result(); err == nil {
		if err = json.Unmarshal([]byte(demoJson), &demo); err == nil {
			return demo, nil
		}
	}

	err = r.Database.GetDB().
		Order("priority desc").
		Limit(30).
		Find(&demo).
		Error

	if data, err := json.Marshal(demo); err != nil {
		return nil, err
	} else {
		if err = r.RedisClient.Set(context.Background(), key, string(data), time.Second*10).Err(); err != nil {
			return nil, err
		}
	}

	return demo, err
}
