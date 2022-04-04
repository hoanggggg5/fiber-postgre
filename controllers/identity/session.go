package identity

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/hoanggggg5/fiber-postgre/config"
	"github.com/hoanggggg5/fiber-postgre/models"
)

func Login(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `validate:"email|minLen:7" message:"email is invalid"`
		Password string `validate:"required|minLen:7"  message:"password is invalid"`
	}

	payload := new(Payload)

	// if err := c.BodyParser(payload); err != nil {
	// 	return c.JSON("LỖI body")
	// }

	v := validate.Struct(payload)

	if v.Validate() {
		// do something ...
		var user *models.User
		config.Database.First(&user)

		if !user.CheckPasswordHash(payload.Password) {
			return c.JSON("Sai password")
		}

		// session, _ := config.Store.Get(c)

		// session.Set("email", user.Email)
		// if err := session.Save(); err != nil {
		// 	return c.JSON("LỖI lưu cookie")
		// }

		return c.Status(202).JSON(201)
	} else {
		fmt.Println(v.Errors) // all error messages
	}
	return c.Status(202).JSON(201)
}
