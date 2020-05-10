package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/service"
)

type TodoController struct {
	s *service.TodoService
}

func NewTodoController(s *service.TodoService) TodoController {
	return TodoController{s}
}

func (ctr *TodoController) GetTodos(ctx *fiber.Ctx) {
	todos := ctr.s.Query.GetTodos()
	ctx.JSON(todos)
}

func (ctr *TodoController) GetTodo(ctx *fiber.Ctx) {
	id := extractID(ctx)
	todo, _ := ctr.s.Query.GetTodo(id)
	ctx.JSON(todo)
}

func (ctr *TodoController) CreateTodo(ctx *fiber.Ctx) {
	newTodo := new(mutation.NewTodo)
	if err := ctx.BodyParser(newTodo); err != nil {
		log.Fatal(err)
	}
	created := ctr.s.Mutation.CreateTodo(*newTodo)
	ctx.JSON(created)
}

func (ctr *TodoController) DeleteTodo(ctx *fiber.Ctx) {
	// todo: delete
	ctx.JSON(fiber.Map{
		"ok": true,
	})
}

func extractID(ctx *fiber.Ctx) int {
	param := ctx.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	return id
}
