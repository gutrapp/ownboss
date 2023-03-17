package routes

import (
	"ownboss/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func createEvent(c *fiber.Ctx) error {
	var event models.Event

	err := c.BodyParser(&event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.CreateEvent(&event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating event" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func updateEvent(c *fiber.Ctx) error {
	var event models.Event

	err := c.BodyParser(&event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.UpdateEvent(&event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating event" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func deleteEvent(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	err = models.DeleteEvent(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting event" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}

func getEvents(c *fiber.Ctx) error {
	events, err := models.GetEvents()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Couldn't get events",
		})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func getEventsFromCompany(c *fiber.Ctx) error {
	var events []models.Event
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	events, err = models.GetCompanyEvents(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting event" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func getEvent(c *fiber.Ctx) error {
	var event models.Event
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	event, err = models.GetEvent(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting event" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}
