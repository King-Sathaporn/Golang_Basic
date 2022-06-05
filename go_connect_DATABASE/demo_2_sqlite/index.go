package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB = constantDatabase()
	todoList := todo{Username: "admin", Title: "test1", Message: "message1"}
	// Insert
	fmt.Println("-----------------Insert-----------------")
	db.Create(&todoList)

	// Query
	fmt.Println("-----------------Query-----------------")
	query(db)

	// Update
	fmt.Println("-----------------Update-----------------")
	var tmp todo
	db.First(&tmp, 1)
	update(db, tmp)
	query(db)

	// Delete
	fmt.Println("-----------------Delete-----------------")
	var deleteTmp []todo
	db.Find(&deleteTmp, "Message like ?", "%est2%")
	delete(db, deleteTmp)
	query(db)
}

func query(_db *gorm.DB) {
	var todos []todo
	_db.Find(&todos)
	printJSON(todos)
}

func printJSON(data []todo) {
	json, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(json))
}

func update(_db *gorm.DB, todo todo) {
	_db.Model(&todo).Update("message", "test2")

}

func delete(_db *gorm.DB, todos []todo) {
	_db.Delete(&todos) //soft delete
	//_db.Unscoped().Delete(&todos) //!hard delete
}

type todo struct {
	gorm.Model
	Username string
	Title    string
	Message  string
}

type todo2 struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

func constantDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&todo{})
	return database
}
