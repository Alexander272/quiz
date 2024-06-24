package models

type Answer struct {
	ID         string `json:"id" db:"id"`
	QuestionID string `json:"questionId" db:"question_id"`
	Text       string `json:"text" db:"text"`
	Image      string `json:"image,omitempty" db:"image"`
	IsCorrect  bool   `json:"isCorrect,omitempty" db:"is_correct"`
}
