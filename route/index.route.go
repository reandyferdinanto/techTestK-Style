package route

import (
	"techTestK-Style/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/member", handler.MemberHandlerCreate)
	r.Get("/member",handler.MemberHandlerGetAll)
	r.Put("/member/:id", handler.MemberHandlerUpdate)
	r.Delete("/member/:id", handler.MemberHandlerDelete)
}