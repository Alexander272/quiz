package models

type AttemptDetails struct {
	AttemptID string             `json:"attemptId" db:"attempt_id"`
	Questions []*AttemptQuestion `json:"questions"`
}

type AttemptQuestion struct {
	ID        string   `json:"id"`
	Answers   []string `json:"answers"`
	Correct   []string `json:"correct,omitempty"`
	IsCorrect bool     `json:"isCorrect"`
	Points    int      `json:"points"`
}

type GetAttemptDetails struct {
	AttemptID   string `json:"attemptId"`
	ShowAnswers bool   `json:"showAnswers"`
}

type CreateAttemptDetailsDTO struct {
	AttemptID string                `json:"attemptId" db:"attempt_id"`
	Questions []*AttemptQuestionDTO `json:"questions"`
}

type AttemptQuestionDTO struct {
	ID string `json:"id"`
	// Points  int      `json:"points"`
	Answers []string `json:"answers"`
}

type AttemptDetailDTO struct {
	ID         string   `json:"id" db:"id"`
	AttemptID  string   `json:"attemptId" db:"attempt_id"`
	QuestionID string   `json:"questionId" db:"question_id"`
	Answers    []string `json:"answers" db:"answers"`
}
