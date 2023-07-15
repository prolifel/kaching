package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prolifel/kaching/config"
	"github.com/prolifel/kaching/models"
	"github.com/redis/go-redis/v9"
)

func validateAPIKey(c *fiber.Ctx, key string) (isValid bool, err error) {
	if len(key) == 0 {
		err = keyauth.ErrMissingOrMalformedAPIKey
		return
	}

	var (
		appNamePrefix = fmt.Sprintf("%s:", os.Getenv(models.EnvAppName))
		userIDPrefix  = "user_id="
	)

	if !strings.HasPrefix(key, appNamePrefix) {
		err = keyauth.ErrMissingOrMalformedAPIKey
		return
	}

	key = key[len(appNamePrefix):]

	if !strings.HasPrefix(key, userIDPrefix) {
		err = keyauth.ErrMissingOrMalformedAPIKey
		return
	}

	key = key[len(userIDPrefix):]

	userID, errx := strconv.ParseInt(key, 10, 0)
	if errx != nil {
		err = keyauth.ErrMissingOrMalformedAPIKey
		return
	}

	c.Context().SetUserValue("user_id", userID)

	return true, nil
}

func main() {
	app := config.New()

	config.Catch(app.InitProgresql())

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	appNew := fiber.New()

	appNew.Use(
		keyauth.New(keyauth.Config{
			Validator: validateAPIKey,
		}),
		logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path} - ${reqHeader:Authorization}\n",
		}),
	)

	appNew.Get("/user", func(c *fiber.Ctx) error {
		ctxUserID := c.Context().UserValue("user_id")
		userID, ok := ctxUserID.(int64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON("gak valid user id mu bos 💀")
		}

		var (
			user models.UserResponse
		)

		userRedis := client.HGetAll(c.UserContext(), fmt.Sprintf("users:%d", userID)).Val()
		if _, ok := userRedis["user_id"]; ok {
			user.UserID, _ = strconv.ParseInt(userRedis["user_id"], 10, 64)
			user.Email = userRedis["email"]
			user.Name = userRedis["name"]
			user.DataSource = "redis"

			return c.Status(fiber.StatusOK).JSON(user)
		}

		app.DB.QueryRowxContext(c.UserContext(), `
			select
				user_id,
				email,
				name
			from users
			where user_id = $1
			limit 1;
		`, userID).StructScan(&user)

		go func() {
			userMap := map[string]string{
				"user_id": fmt.Sprint(user.UserID),
				"email":   user.Email,
				"name":    user.Name,
			}
			for k, v := range userMap {
				err := client.HSet(context.Background(), fmt.Sprintf("users:%d", userID), k, v).Err()
				if err != nil {
					log.Printf("error while hset: %v", err.Error())
				}
			}
			client.Expire(context.Background(), fmt.Sprintf("users:%d", userID), time.Second*10)
		}()

		user.DataSource = "db"

		return c.Status(fiber.StatusOK).JSON(user)
	})

	appNew.Listen(":3000")
}
