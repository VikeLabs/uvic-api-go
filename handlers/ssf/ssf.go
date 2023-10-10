package ssf

import "github.com/gofiber/fiber/v2"

func Router(r fiber.Router) {
	r.Get("/buildings", Buildings)
}
