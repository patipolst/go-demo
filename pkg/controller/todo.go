package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type TodoController struct {
	q query.TodoQuery
	m mutation.TodoMutation
}

func NewTodoController(q query.TodoQuery, m mutation.TodoMutation) TodoController {
	return TodoController{q, m}
}

func (ctr *TodoController) GetTodos(ctx *fiber.Ctx) {
	todos := ctr.q.GetTodos()
	ctx.JSON(todos)
}

func (ctr *TodoController) GetTodo(ctx *fiber.Ctx) {
	id := extractID(ctx)
	todo, _ := ctr.q.GetTodo(id)
	ctx.JSON(todo)
}

func (ctr *TodoController) CreateTodo(ctx *fiber.Ctx) {
	newTodo := new(mutation.NewTodo)
	if err := ctx.BodyParser(newTodo); err != nil {
		log.Fatal(err)
	}
	created := ctr.m.CreateTodo(*newTodo)
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
