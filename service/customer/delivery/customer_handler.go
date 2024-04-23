package delivery

import (
	"itmx_test/domain"
	"itmx_test/middleware"
	"itmx_test/service/entity"
	"itmx_test/service/customer/usecase"

	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Message string `json:"message"`
}

type CustomerHandler struct {
	cu usecase.CustomerUsecase
}

func NewCustomerHandler(f *fiber.App, cu usecase.CustomerUsecase) {
	handler := &CustomerHandler{cu}

	// group name
	customer := f.Group("/customers")

	// Create
	customer.Post("", handler.CreateCustomer)

	// GetByID
	customer.Get("/:id", handler.GetCustomer)

	// Update
	customer.Put("/:id", handler.UpdateCustomer)

	// Delete By ID
	customer.Delete("/:id", handler.DeleteCustomer)
}

type CustomerBody struct {
	Name string `json:"name" validate:"required,max=100"`
	Age  int    `json:"age" validate:"required,numeric,min=1,max=110"`
}

func (ch *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	var input CustomerBody

	// Parser input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Validate input
	if err := middleware.Validate(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(middleware.ErrorResponse(err))
	}

	cutomer := &entity.Customer{
		Name: input.Name,
		Age:  input.Age,
	}

	// create customer usecase
	if err := ch.cu.CreateCustomer(cutomer); err != nil {
		return c.Status(domain.GetStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Message": "create customer successful",
	})
}

func (ch *CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	customer, err := ch.cu.GetCustomerByID(id)
	if err != nil {
		return c.Status(domain.GetStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

type CustomerUpdateBody struct {
	Name string `json:"name" validate:"max=100"`
	Age  int    `json:"age" validate:"numeric,min=1,max=110"`
}

func (ch *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var input CustomerUpdateBody

	// Parser input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Validate input
	if err := middleware.Validate(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(middleware.ErrorResponse(err))
	}

	cutomerUpdate := &entity.Customer{
		Name: input.Name,
		Age:  input.Age,
	}

	if err := ch.cu.UpdateCustomerByID(cutomerUpdate, id); err != nil {
		return c.Status(domain.GetStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (ch *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ch.cu.DelCustomerByID(id); err != nil {
		return c.Status(domain.GetStatusCode(err)).JSON(ResponseError{Message: err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
