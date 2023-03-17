package routes

import (
	"ownboss/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func createCompany(c *fiber.Ctx) error {
	var company models.Company

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}
	name := c.Params("name")
	salary, err := strconv.ParseUint(c.Params("salary"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing salary" + err.Error(),
		})
	}
	err = c.BodyParser(&company)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	company, err = models.CreateCompany(company)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating company" + err.Error(),
		})
	}

	employee := models.Employee{
		CompanyID: company.ID,
		UserID:    uint(id),
		Position:  "CEO",
		Name:      name,
		Salary:    uint(salary),
		JobTitle:  "CEO",
	}

	err = models.CreateEmployee(&employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating employee" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Company created",
	})
}

func getCompany(c *fiber.Ctx) error {
	var company models.Company
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	company, err = models.GetCompany(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting company" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}

func updateCompany(c *fiber.Ctx) error {
	var company models.Company

	err := c.BodyParser(&company)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = models.UpdateCompany(&company)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating company" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(company)
}

func deleteCompany(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	err = models.DeleteCompany(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting company" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}

func getCompanies(c *fiber.Ctx) error {
	companies, err := models.GetCompanies()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Couldn't get companies",
		})
	}

	return c.Status(fiber.StatusOK).JSON(companies)
}
