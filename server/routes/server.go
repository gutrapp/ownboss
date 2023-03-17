package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func SetupAndListen() {
	server := fiber.New()
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 24,
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	server.Use(Middleware())

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Server Running...")
	})

	auth := server.Group("/auth")
	auth.Post("/register", register)
	auth.Post("/login", login)
	auth.Delete("/logout", logout)
	auth.Get("/", logged)

	user := server.Group("/user")
	user.Get("/", getJobs)

	employee := server.Group("/employee")
	employee.Get("/", getEmployees)
	employee.Get("/:id", getEmployee)
	employee.Get("/company/:id", getEmployeesFromCompany)
	employee.Post("/", createEmployee)
	employee.Put("/", updateEmployee)
	employee.Delete("/", deleteEmployee)

	company := server.Group("/company")
	company.Get("/", getCompanies)
	company.Get("/:id", getCompany)
	company.Post("/", createCompany)
	company.Put("/", updateCompany)
	company.Delete("/", deleteCompany)

	service := server.Group("/service")
	service.Get("/", getServices)
	service.Get("/:id", getService)
	service.Get("/company/:id", getServicesFromCompany)
	service.Post("/", createService)
	service.Put("/", updateService)
	service.Delete("/", deleteService)

	event := server.Group("/service")
	event.Get("/", getEvents)
	event.Get("/:id", getEvent)
	event.Get("/company/:id", getEventsFromCompany)
	event.Post("/", createEvent)
	event.Put("/", updateEvent)
	event.Delete("/", deleteEvent)

	server.Listen(":4000")
}
