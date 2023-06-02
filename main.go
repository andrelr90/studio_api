package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"studio_api_project/main/api"
	"studio_api_project/main/repositories"
)

func main() {
	router := gin.Default()

	repositories.PopulateClassesWithExamples();
	repositories.PopulateBookingsWithExamples();
	api.StartClassesAPI(router)
	api.StartBookingsAPI(router)
	
	log.Fatal(router.Run(":8000"))
}
