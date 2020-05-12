package controller

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/service"
)

type User struct {
	s *service.UserService
}

func NewUser(s *service.UserService) *User {
	return &User{s}
}

func (ctr *User) GetUsers(ctx *fiber.Ctx) {
	users := ctr.s.Query.GetUsers()
	ctx.JSON(users)
}

func (ctr *User) GetUser(ctx *fiber.Ctx) {
	id := extractID(ctx)
	user, _ := ctr.s.Query.GetUser(id)
	ctx.JSON(user)
}

func (ctr *User) CreateUser(ctx *fiber.Ctx) {
	newUser := new(mutation.User)
	if err := ctx.BodyParser(newUser); err != nil {
		log.Fatal(err)
	}
	created := ctr.s.Mutation.CreateUser(*newUser)
	ctx.JSON(created)
}

func (ctr *User) DeleteUser(ctx *fiber.Ctx) {
	// user: delete
	ctx.JSON(fiber.Map{
		"ok": true,
	})
}
