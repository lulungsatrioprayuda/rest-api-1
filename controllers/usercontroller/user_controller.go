package usercontroller

import "github.com/gofiber/fiber/v2"

// Index returns a fiber.Handler that sends "hello world" as a string.
//
// c is a pointer to a fiber.Ctx object.
// It returns an error.
func Index(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
