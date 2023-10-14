package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/database"
	"github.com/gofiber/fiber/v2"
)

type Building struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func routeBuildings(c *fiber.Ctx) error {
	db := database.New()

	var building []Building
	rows, err := db.Queryx(`SELECT id,name FROM buildings ORDER BY name ASC;`)
	if err != nil {
		return err
	}

	var b Building
	for rows.Next() {
		rows.StructScan(&b)
		building = append(building, b)
	}

	return c.Status(http.StatusOK).JSON(building)
}
