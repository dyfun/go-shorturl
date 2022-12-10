package routes

import (
	"fmt"
	"github.com/dyfun/go-shorturl/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"time"
)

func ResolveUrl(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)

	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short url not found in db"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to db"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	fmt.Println("Redirecting...")
	time.Sleep(5 * time.Second)

	return c.Redirect(value, 301)
}
