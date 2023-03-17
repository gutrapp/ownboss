package routes

import (
	"ownboss/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getEmployees(c *fiber.Ctx) error {
	employees, err := models.GetEmployees()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Couldn't get employees",
		})
	}

	return c.Status(fiber.StatusOK).JSON(employees)
}

func getEmployeesFromCompany(c *fiber.Ctx) error {
	var employees []models.Employee
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	employees, err = models.GetEmployeesByCompanyId(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting employee" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(employees)
}

func createEmployee(c *fiber.Ctx) error {
	var employee models.Employee

	err := c.BodyParser(&employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.CreateEmployee(&employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating employee" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

func getEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	employee, err = models.GetEmployee(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting employee" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

func updateEmployee(c *fiber.Ctx) error {
	var employee models.Employee

	err := c.BodyParser(&employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.UpdateEmployee(&employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating employee" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(employee)
}

func deleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	err = models.DeleteEmployee(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting employee" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}
