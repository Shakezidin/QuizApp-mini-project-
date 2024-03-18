package model

import "gorm.io/gorm"

// Quiz represents a quiz that can be created by an admin.
type Quiz struct {
	gorm.Model
	Title       string `json:"title" form:"title" gorm:"not null"`
	Description string `json:"description" form:"description" gorm:"not null"`
	Questions   int    `json:"questions" form:"questions" gorm:"not null"`
	Duration    int    `json:"duration" form:"duration" gorm:"not null"` // Duration in minutes
	Status      string `json:"status" form:"status" gorm:"default:pending"`
}

// Question represents a quiz question.
type Question struct {
	gorm.Model
	QuizID        uint   `json:"quiz_id" form:"quiz_id" gorm:"not null"`
	Quiz          Quiz   `gorm:"ForeignKey:QuizID"`
	Text          string `json:"text" form:"text" gorm:"not null"`
	OptionA       string `json:"option_a" form:"option_a" gorm:"not null"`
	OptionB       string `json:"option_b" form:"option_b" gorm:"not null"`
	OptionC       string `json:"option_c" form:"option_c" gorm:"not null"`
	OptionD       string `json:"option_d" form:"option_d" gorm:"not null"`
	CorrectAnswer int    `json:"correct_answer" form:"correct_answer" gorm:"not null"`
}
