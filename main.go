package main

import (
	"CrawlS/crawldata"
	"CrawlS/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	
	//crawl data va cho vao database
	crawldata.Crawl() 
	log.Printf("Crawl Done")

	//create new record
	router := gin.Default()
	router.POST("/", handler.CreateQuestion)

	//get all records
	router.GET("/", handler.GetAllQuestions)

	//get record by id
	router.GET("/:id", handler.GetQuestionByID)

	//update record
	router.PUT("/:id", handler.UpdateQuestionByID)

	//delete record by id
	router.DELETE("/:id", handler.DeleteQuestionByID)
	router.Run()
}
