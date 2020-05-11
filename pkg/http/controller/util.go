package controller

import (
	"strconv"

	"github.com/gofiber/fiber"
)

func extractID(ctx *fiber.Ctx) int {
	param := ctx.Params("id")
	id, _ := strconv.Atoi(param) // user: error handling
	return id
}
