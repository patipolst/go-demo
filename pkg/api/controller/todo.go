package controller

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/service"
)

type Todo struct {
	s *service.Todo
}

func NewTodo(s *service.Todo) *Todo {
	return &Todo{s}
}

func (ctr *Todo) GetTodos(ctx *fiber.Ctx) {
	todos := ctr.s.Query.GetTodos()
	ctx.JSON(todos)
}

func (ctr *Todo) GetTodo(ctx *fiber.Ctx) {
	id := extractID(ctx)
	todo, _ := ctr.s.Query.GetTodo(id)
	ctx.JSON(todo)
}

func (ctr *Todo) CreateTodo(ctx *fiber.Ctx) {
	newTodo := new(mutation.Todo)
	if err := ctx.BodyParser(newTodo); err != nil {
		log.Fatal(err)
	}
	created := ctr.s.Mutation.CreateTodo(*newTodo)
	ctx.JSON(created)
}

func (ctr *Todo) DeleteTodo(ctx *fiber.Ctx) {
	// todo: delete
	ctx.JSON(fiber.Map{
		"ok": true,
	})
}
