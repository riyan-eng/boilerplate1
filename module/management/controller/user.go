package controller

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("users")
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("user detail" + id)
}

func CreateUser(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Write into Casbin rule DB
		e.AddGroupingPolicy(fmt.Sprint("3"), "user")
		return c.JSON(fiber.Map{
			"data":    "new user registered successfully",
			"message": "ok",
		})
	}
}

func UpdateUser(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse input from request body
		e.UpdateGroupingPolicy([]string{"3", "user"}, []string{"3", "admin"})
		return c.JSON(fiber.Map{
			"data":    "update user successfully",
			"message": "ok",
		})
	}
}

func DeleteUser(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		e.RemoveGroupingPolicy("1")
		return c.JSON(fiber.Map{
			"error":   false,
			"message": "Delete user successfully!",
		})
	}
}
