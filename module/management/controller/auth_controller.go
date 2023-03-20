package controller

import (
	"boilerplate/module/management/controller/dto"
	"boilerplate/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	userID := c.Query("user_id", "1")
	companyID := c.Query("company_id", "comp_123")

	// Create JWT token with userID.
	accessToken, refreshToken, err := util.GenerateJWT(string(c.Request().Host()), userID, companyID, 1)

	// accessToken, refreshToken, exp, err := util.GenerateTokenPair(string(c.Request().Host()), userID, "comp1", c.Context())
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"data":    "internal server error",
			"message": "bad",
		})
	}

	// Create cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Minute * 1),
		HTTPOnly: true,
	}

	// save cookie
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"exp":          "exp",
		},
		"message": "ok",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	refreshRequest := new(dto.RefreshRequest)
	if err := c.BodyParser(refreshRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "bad request",
			"message": "bad",
		})
	}

	claims, err := util.ParseToken(refreshRequest.RefreshToken, "AllYourBaseRefresh")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorized!",
		})
	}

	if err := util.ValidateToken(claims, true, c.Context()); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "not authorized!",
		})
	}

	accessToken, refreshToken, err := util.GenerateJWT(string(c.Request().Host()), claims.UserID, claims.CompanyID, 3)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":    "internal server error",
			"message": "bad",
		})
	}

	// Create cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Minute * 1),
		HTTPOnly: true,
	}

	// save cookie
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"exp":          "exp",
		},
		"message": "ok",
	})
}
