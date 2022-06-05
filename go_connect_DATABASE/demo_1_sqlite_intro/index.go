package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model //? generated ID, created_at, updated_at, deleted_at
	Code       string
	Price      uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{}) //? create table if not exists

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	//? .First() return first object in the table.
	//fmt.Println(product)
	// json, _ := json.MarshalIndent(product, "", "  ")
	// fmt.Println(string(json))
	db.First(&product, "code = ?", "D42") // find product with code D42

	//query with find array
	var products []Product
	db.Find(&products)
	json, _ := json.MarshalIndent(products, "", "  ")
	fmt.Println(string(json))

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)
}
