package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/satori/uuid.go"
)

const (
	CategorySingle string = "single"
	CategoryOffer  string = "offer"
	CategoryAddon  string = "addon"
)

type Product struct {
	Base
	Name        string `json:"name" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	Category    string `json:"category" valid:"notnull"`
	Amount      int    `json:"amount" valid:"notnull,numeric"`
}

type Products struct {
	Products []Product `json: "products"`
}

func (p *Product) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	if p.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	switch p.Category {
	case CategorySingle:
	case CategoryOffer:
	case CategoryAddon:
	default:
		return errors.New("invalid category")
	}

	return nil
}

func NewProduct(product *Product) error {

	product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()
	err := product.isValid()
	if err != nil {
		return err
	}

	return nil
}
