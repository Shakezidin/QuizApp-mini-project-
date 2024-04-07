package handler

import (
	"net/http"

	"github.com/Shakezidin/pkg/model"
	"github.com/gin-gonic/gin"
)

func (a *AdminHandler) GetQuizLanguages(c *gin.Context) {
	languages := a.AdminSVC.GetQuizLanguagesSVC()

	c.JSON(http.StatusOK, gin.H{"languages": languages})
}

func (a *AdminHandler) StartQuiz(c *gin.Context) {
	language := c.Param("language")

	if language == "" {
		c.JSON(400, gin.H{
			"error": "Language not selected",
		})
		return
	}

	quizQuestion, err := a.AdminSVC.StartQuizSVC(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"question": quizQuestion})
}

func (a *AdminHandler) CheckAnswer(c *gin.Context) {
	var userAnswer model.UserAnswer
	if err := c.BindJSON(&userAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := a.AdminSVC.CheckAnswerSVC(userAnswer)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if response {
		c.JSON(http.StatusOK, gin.H{"question": userAnswer.Question, "is_correct": "correct"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"question": userAnswer.Question, "is_correct": "wrong"})
}
