package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sandbox-science/online-learning-platform/configs/database"
	"github.com/sandbox-science/online-learning-platform/internal/entity"
	"github.com/sandbox-science/online-learning-platform/internal/utils"
)

// Login auths a user by checking email and password.

func Login(c *fiber.Ctx) error {

	var data entity.Login

	// Parse the request body
	if err := c.BodyParser(&data); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	// Find user with email
	var user entity.Account

	if err := database.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	// Check if the password matches

	if err := utils.CheckPasswordHash(data.Password, user.Password); err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	// Successful login

	return c.JSON(fiber.Map{

		"message": "Login successful",
		"user":    user,
	})
}
