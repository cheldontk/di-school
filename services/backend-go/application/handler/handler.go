package handler

import (
	"database/sql"
	"errors"
	"log"

	"github.com/cheldontk/di-school/di-go/application/database"
	"github.com/cheldontk/di-school/di-go/domain/model"
	"github.com/gofiber/fiber/v2"
)

const (
	all_products   string = "SELECT name, description, category, amount FROM products order by name"
	one_product    string = "SELECT * FROM products WHERE id = $id"
	insert_product string = "INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)"
	delete_product string = "DELETE FROM products WHERE id = $id"
)

func GetAllProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query(all_products)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			c.SendStatus(500)
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		result.Products = append(result.Products, product)
	}

	//Return all products in JSON
	if err := c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All products return successfully",
	}); err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.SendStatus(200)
}

func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query(one_product, id)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&id, &product.Name, &product.Description, &product.Category, &product.Amount); err {
		case sql.ErrNoRows:
			log.Println("No rows were returned!")
			c.SendStatus(500)
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		case nil:
			log.Println(product.Name, product.Description, product.Category, product.Amount)
		default:
			c.SendStatus(500)
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
	}

	if err := c.JSON(&fiber.Map{
		"success": false,
		"message": "Successfully fetched product",
		"product": product,
	}); err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.SendStatus(200)
}

func CreateProduct(c *fiber.Ctx) error {

	//instance
	p := new(model.Product)

	//Parse body
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.SendStatus(400)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	//Validate
	if err := model.NewProduct(p); err != nil {
		log.Println(err)
		c.SendStatus(400)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	//insert product
	res, err := database.DB.Query(insert_product, p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	//result
	log.Println(res)

	//return JSON Product
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": p,
	}); err != nil {
		c.SendStatus(400)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   errors.New("error creating product"),
		})
	}
	return c.SendStatus(200)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := database.DB.Query(delete_product, id)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	//Print result
	log.Println(res)

	//return JSON product
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product delete successfully",
	}); err != nil {
		c.SendStatus(500)
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.SendStatus(200)
}
