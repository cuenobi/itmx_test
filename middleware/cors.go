package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CORSMiddleware(whiteList []string) fiber.Handler {
    whiteListMap := make(map[string]bool)
    for _, origin := range whiteList {
        whiteListMap[origin] = true
    }

    return func(c *fiber.Ctx) error {
        if c.Method() == "OPTIONS" {
            c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            c.Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, Accept")
            c.Set("Access-Control-Max-Age", "86400")

            origin := c.Get("Origin")
            if whiteListMap[origin] || whiteListMap["*"] {
                c.Set("Access-Control-Allow-Origin", origin)
                c.Set("Access-Control-Allow-Credentials", "true")
            }

            return c.SendStatus(fiber.StatusNoContent)
        }

        origin := c.Get("Origin")
        if !whiteListMap[origin] && !whiteListMap["*"] {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "message": "Not allowed by CORS",
            })
        }

        c.Set("Access-Control-Allow-Origin", origin)
        c.Set("Access-Control-Allow-Credentials", "true")

        return c.Next()
    }
}
