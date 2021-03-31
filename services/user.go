package services

import (
	"login-microservice/database"
	"login-microservice/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Function to hash a string, particularly a password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Get User by ID
func GetUserByID(c *fiber.Ctx) error {
	db := database.DB;
	param := c.Params("id")

	var user model.User
	db.Find(&user, param)

	if(user.Username == "") {
		return c.Status(404).JSON("No user found")
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(404).JSON("Check user body")
	}
	
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON("Can't hash password")
	}
	
	user.Password = hash
	if err := db.Create(&user); err != nil {
		return c.Status(500).JSON(err)
	}

	userResponse := model.UserResponse{
		Name: user.Name,
		Username: user.Username,
	}

	return c.JSON(userResponse)

}