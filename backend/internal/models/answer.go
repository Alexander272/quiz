package models

type AnswerList struct {
	QuestionID string    `json:"questionId" db:"question_id"`
	List       []*Answer `json:"list"`
}

type Answer struct {
	ID         string `json:"id" db:"id"`
	QuestionID string `json:"questionId" db:"question_id"`
	Number     int    `json:"number" db:"number"`
	Text       string `json:"text" db:"text"`
	Image      string `json:"image,omitempty" db:"image"`
	IsCorrect  bool   `json:"isCorrect,omitempty" db:"is_correct"`
}

type GetAnswersDTO struct {
	QuizID     string `json:"quizId"`
	QuestionID string `json:"questionId"`
	HasCorrect bool   `json:"hasCorrect"`
}

type AnswerDTO struct {
	ID         string `json:"id" db:"id"`
	QuestionID string `json:"questionId" db:"question_id"`
	Number     int    `json:"number" db:"number" binding:"required,min=1"`
	Text       string `json:"text" db:"text" binding:"required,min=1"`
	Image      string `json:"image" db:"image"`
	IsCorrect  bool   `json:"isCorrect" db:"is_correct"`
}

type DeleteAnswerDTO struct {
	ID string `json:"id" db:"id"`
}
