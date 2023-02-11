package main

import (
	"SocialNetHTTPService/internal/app"
)

// @title SocialNet Service
// @version 1.0
// @description API Server for SocialNet Application
// @host localhost:8080
// @BasePath /
// @query.collection.format	multi
func main() {
	app.Run()
}
