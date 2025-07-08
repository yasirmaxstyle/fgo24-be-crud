package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	godotenv.Load()
	db, err := strconv.Atoi(os.Getenv("RDDB"))
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDADDRESS"),
		Password: os.Getenv("RDPASSWORD"),
		DB:       db,
	})

	return rdb
}

// var ctx = context.Background()

// func TrackingRedisKey(id int, key string) {
// 	trKey := fmt.Sprintf("contact_set:%d", id)
// 	RedisClient().SAdd(ctx, trKey, key)
// }

// func SetContactCache(id int, key string, value interface{}, expiration time.Duration) {
// 	RedisClient().Set(ctx, key, value, expiration)
// 	TrackingRedisKey(id, key)
// }

// func CleanupTrackedKeys(id int) {
// 	trKey := fmt.Sprintf("contact_set:%d", id)
// 	keys, err := RedisClient().SMembers(ctx, trKey).Result()
// 	if err != nil {
// 		log.Printf("failed to get tracked key for contact: %d", id)
// 	}

// 	if len(keys) > 0 {
// 		_, err := RedisClient().Del(ctx, keys...).Result()
// 		if err != nil {
// 			log.Printf("failed to delete tracked key: %v", err)
// 		}
// 	}

// 	RedisClient().Del(ctx, trKey)

// 	commonKeys := []string{
// 		"/contacts",
// 		fmt.Sprintf("/%d", id),
// 	}

// 	RedisClient().Del(ctx, commonKeys...)
// }
