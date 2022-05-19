package handler

import (
	"CrawlS/dbconnect"
	"CrawlS/entity"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db, _ = dbconnect.DBConnect("crawldata") //mo connection

func CreateQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	question := entity.Questions{ID: id, Author: c.PostForm("author"), Title: c.PostForm("title"), Link: c.PostForm("link"), Votes: c.PostForm("votes"), Answers: c.PostForm("answers"), Views: c.PostForm("views")}
	db.Save(&question)
	fmt.Println(question.ID)
	fmt.Println("id: ", c.PostForm("id"))
	fmt.Println("author: ", c.PostForm("author"))
	fmt.Println("title: ", c.PostForm("title"))
	fmt.Println("link: ", c.PostForm("link"))
	fmt.Println("votes: ", c.PostForm("votes"))
	fmt.Println("answers: ", c.PostForm("answers"))
	fmt.Println("views: ", c.PostForm("views"))
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Question created successfully!", "resourceId": question.ID})
}

func GetAllQuestions(c *gin.Context) {
	var questions []entity.Questions
	db.Find(&questions)
	if len(questions) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No question found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": questions})
}

func GetQuestionByID(c *gin.Context) {
	var question entity.Questions
	questionID := c.Param("id")
	db.First(&question, questionID)
	if question.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No question found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": question})
}

func UpdateQuestionByID(c *gin.Context) {
	var question entity.Questions
	questionID := c.Param("id")
	db.First(&question, questionID)
	if question.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No question found!"})
		return
	}
	db.Model(&question).Update("author", c.PostForm("author"))
	db.Model(&question).Update("title", c.PostForm("title"))
	db.Model(&question).Update("link", c.PostForm("link"))
	db.Model(&question).Update("votes", c.PostForm("votes"))
	db.Model(&question).Update("answer", c.PostForm("answer"))
	db.Model(&question).Update("views", c.PostForm("views"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Question updated successfully!"})
}

func DeleteQuestionByID(c *gin.Context) {
	var question entity.Questions
	todoID := c.Param("id")
	db.First(&question, todoID)
	if question.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No question found!"})
		return
	}
	db.Delete(&question)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Question deleted successfully!"})
}
