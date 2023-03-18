package middleware

import (
	"boilerplate/config"
	"boilerplate/util"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/errgroup"
)

func AuthorizeJwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data":    "authorization header is required",
				"message": "bad",
			})
		}
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"data":    "undefined token",
				"message": "bad",
			})
		}

		tokenString := splitToken[1]
		claims, err := util.ParseToken(tokenString, "AllYourBaseAccess")
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// Malformed token -> Delete Cookie
				c.ClearCookie("jwt")
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   true,
					"message": "missing or malformed jwt!",
				})
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error":   true,
					"message": "expired token!",
				})
			} else {
				// Cannot handle -> Delete Cookie
				c.ClearCookie("jwt")
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error":   true,
					"message": "error when processing identity!",
				})
			}
		}

		if err := util.ValidateToken(claims, false, c.Context()); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "not authorized!",
			})
		}

		c.Locals("userID", claims.UserID)
		c.Locals("companyID", claims.CompanyID)

		g := new(errgroup.Group)
		g.Go(func() (err error) {
			err = config.Redis.Expire(c.Context(), fmt.Sprintf("token-%s", claims.UserID), time.Minute*time.Duration(1)).Err()
			return
		})
		g.Wait()

		return c.Next()
	}
}
