package fiber

import (
	"github.com/gofiber/fiber/v2"
	svc "github.com/macgeargear/gmg/services"
)

type UserHandler interface {
	UserCtx(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userhandler struct {
	userService svc.UserService
}

func NewUserHandler(userService svc.UserService) UserHandler {
	return &userhandler{userService: userService}
}

// Delete implements UserHandler.
func (h *userhandler) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Get implements UserHandler.
func (h *userhandler) Get(c *fiber.Ctx) error {
	return nil
}

// Post implements UserHandler.
func (h *userhandler) Post(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements UserHandler.
func (h *userhandler) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UserCtx implements UserHandler.
func (h *userhandler) UserCtx(c *fiber.Ctx) error {
	panic("unimplemented")
}
