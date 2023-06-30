package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lulungsatrioprayuda/rest-api-1/database/connection"
	"github.com/lulungsatrioprayuda/rest-api-1/models"
)

// Index returns a fiber.Handler that sends "hello world" as a string.
//
// c is a pointer to a fiber.Ctx object.
// It returns an error.
func Index(c *fiber.Ctx) error {

	var users []models.User
	if err := connection.DB.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	// Parse the request body to get the user data
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// Checking user if the email already exists
	var existsUser models.User
	if err := connection.DB.Where("email = ?", user.Email).First(&existsUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Sorry but email has already been taken",
		})
	}

	// Save the user in the database
	err := connection.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func ShowUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	err := connection.DB.First(&user, id).Error
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := connection.DB.First(&user, id).Error; err != nil {
		return err
	}

	// Update the user's data based on the request body
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// Save the updated user in the database
	if err := connection.DB.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := connection.DB.First(&user, id).Error; err != nil {
		return err
	}

	// Delete the user from the database
	if err := connection.DB.Delete(&user).Error; err != nil {
		return err
	}

	return c.JSON(user)
}
