package noteRoutes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/alitemel89/gofiber-restapi/internals/handlers/note"
)


func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
    // Create a Note
    note.Post("/", noteHandler.CreateNote)

	// Get all Notes
	note.Get("/", noteHandler.GetNotes)

	// Update a Note
	note.Put("/:noteId", noteHandler.UpdateNote)

	// Delete a Note
	note.Delete("/:noteId", noteHandler.DeleteNote)
}