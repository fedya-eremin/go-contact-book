package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Service ...
type Service struct{}

// New ...
func New() *Service {
	return new(Service)
}

type Contact struct {
	ID         int
	Username   string
	GivenName  string
	FamilyName string
	Phone      []struct {
		TypeID      int
		CountryCode int
		Operator    int
		Number      int
	}
	Email     []string
	Birthdate time.Time
}

type Group struct {
	ID          int
	Title       string
	Description string
	Contacts    []int
}

func (s *Service) GetContact(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	contact := Contact{
		ID:         id,
		Username:   "testuser",
		GivenName:  "Иван",
		FamilyName: "Иванов",
		Phone: []struct {
			TypeID      int
			CountryCode int
			Operator    int
			Number      int
		}{
			{TypeID: 1, CountryCode: 7, Operator: 900, Number: 1234567},
		},
		Email:     []string{"ivan@example.com"},
		Birthdate: time.Now(),
	}
	return ctx.JSON(contact)
}

func (s *Service) PostContact(ctx *fiber.Ctx) error {
	var contact Contact
	if err := ctx.BodyParser(&contact); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(contact)
}

func (s *Service) PutContact(ctx *fiber.Ctx) error {
	if _, err := ctx.ParamsInt("id"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	var contact Contact
	if err := ctx.BodyParser(&contact); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные"})
	}
	return ctx.JSON(contact)
}

func (s *Service) DeleteContact(ctx *fiber.Ctx) error {
	if _, err := ctx.ParamsInt("id"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.Status(fiber.StatusNoContent).SendString("")
}

func (s *Service) GetGroup(ctx *fiber.Ctx) error {
	if _, err := ctx.ParamsInt("id"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	group := Group{
		ID:          1,
		Title:       "Друзья",
		Description: "Группа близких друзей",
		Contacts:    []int{1, 2, 3},
	}
	return ctx.JSON(group)
}

func (s *Service) PostGroup(ctx *fiber.Ctx) error {
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(group)
}

func (s *Service) PutGroup(ctx *fiber.Ctx) error {
	if _, err := ctx.ParamsInt("id"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Некорректные данные"})
	}
	return ctx.JSON(group)
}

func (s *Service) DeleteGroup(ctx *fiber.Ctx) error {
	if _, err := ctx.ParamsInt("id"); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.Status(fiber.StatusNoContent).SendString("")
}
