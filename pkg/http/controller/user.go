package controller

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/service"
)

type UserController struct {
	s *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{s}
}

func (ctr *UserController) GetUsers(ctx *fiber.Ctx) {
	users := ctr.s.Query.GetUsers()
	ctx.JSON(users)
}

func (ctr *UserController) GetUser(ctx *fiber.Ctx) {
	id := extractID(ctx)
	user, _ := ctr.s.Query.GetUser(id)
	ctx.JSON(user)
}

func (ctr *UserController) CreateUser(ctx *fiber.Ctx) {
	newUser := new(mutation.NewUser)
	if err := ctx.BodyParser(newUser); err != nil {
		log.Fatal(err)
	}
	created := ctr.s.Mutation.CreateUser(*newUser)
	ctx.JSON(created)
}

func (ctr *UserController) DeleteUser(ctx *fiber.Ctx) {
	// user: delete
	ctx.JSON(fiber.Map{
		"ok": true,
	})
}
