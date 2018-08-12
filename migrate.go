package main

import "fmt"

// Migrate .
func Migrate() {
	db.AutoMigrate(&User{}, &Task{})
	db.Model(&Task{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Migrations runned.")
}
