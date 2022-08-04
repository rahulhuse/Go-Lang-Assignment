package main

import "gorm-test/routes"

func main() {
	r := routes.SetupRouter()
	_ = r.Run(":8080")
}
