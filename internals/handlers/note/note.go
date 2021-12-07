package noteHandler

import (
	"github.com/alitemel89/gofiber-restapi/database"
	"github.com/alitemel89/gofiber-restapi/internals/model"
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

// Get Notes
func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes[]model.Note

	// find all notes in the database
	db.Find(&notes)

	// if no note is present return an error
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	// Else return notes
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": notes})
}

// Update a Note
func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title string `json:"title"`
		SubTitle string `json:"subtitle"`
		Text string `json:"text"`
	}

	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Store the body containing the updated data and return error if encounterd
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message":"Review your input", "data": err })
	}

	// Edit the note
	note.Title = updateNoteData.Title
	note.Subtitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// Save the changes
	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found"})

}


// Delete note
func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Note"})
}
