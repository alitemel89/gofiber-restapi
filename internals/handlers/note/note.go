package noteHandler

import (
	"github.com/alitemel89/golang-restapi/database"
	"github.com/alitemel89/golang-restapi/internals/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Add a uuid to the note
	note.ID = uuid.New()

	// Create the Note and Return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})
}