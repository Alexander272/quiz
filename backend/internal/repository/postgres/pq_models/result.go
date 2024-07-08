package pq_models

import "github.com/lib/pq"

type Result struct {
	QuestionID    string         `json:"questionId" db:"question_id"`
	IsCorrect     bool           `json:"isCorrect" db:"is_correct"`
	UserAnswer    pq.StringArray `json:"userAnswer" db:"answer"`
	CorrectAnswer pq.StringArray `json:"correctAnswer" db:"correct"`
}
