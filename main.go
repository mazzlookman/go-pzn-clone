package main

import "go-pzn-clone/app"

func main() {
	server := app.RouterInitialized()
	server.Run(":8080")
}
