package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func AuthorizeCasbin(enforce *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get current user
		userID, ok := c.Locals("userID").(string)
		if userID == "" || !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"data":    "current logged in user not found",
				"message": "unauthorized",
			})
		}

		// load policy
		if err := enforce.LoadPolicy(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "failed to load casbin policy",
				"message": "bad",
			})
		}

		// casbin enforce policy
		accepted, err := enforce.Enforce(userID, c.OriginalURL(), c.Method()) // userID - url - method
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "error when authorizing user's accessibility",
				"message": "bad",
			})
		}
		if !accepted {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   true,
				"message": "Unauthorized!",
			})
		}

		return c.Next()
	}
}
