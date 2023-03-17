package routes

import (
	"ownboss/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func createService(c *fiber.Ctx) error {
	var service models.Service

	err := c.BodyParser(&service)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.CreateService(&service)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating service" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(service)
}

func updateService(c *fiber.Ctx) error {
	var service models.Service

	err := c.BodyParser(&service)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.UpdateService(&service)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating service" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(service)
}

func deleteService(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	err = models.DeleteService(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting service" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}

func getServices(c *fiber.Ctx) error {
	services, err := models.GetEmployees()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Couldn't get services",
		})
	}

	return c.Status(fiber.StatusOK).JSON(services)
}

func getServicesFromCompany(c *fiber.Ctx) error {
	var services []models.Service
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	services, err = models.GetCompanyServices(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting service" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(services)
}

func getService(c *fiber.Ctx) error {
	var service models.Service
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	service, err = models.GetService(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting service" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(service)
}
