package pq_models

import "github.com/lib/pq"

type AttemptDetails struct {
	ID         string         `json:"id" db:"id"`
	AttemptID  string         `json:"attemptId" db:"attempt_id"`
	QuestionID string         `json:"questionId" db:"question_id"`
	Answers    pq.StringArray `json:"answers" db:"answers"`
	Correct    pq.StringArray `json:"correct" db:"correct"`
	Points     int            `json:"points" db:"points"`
}

type AttemptDetailsDTO struct {
	ID         string         `json:"id" db:"id"`
	AttemptID  string         `json:"attemptId" db:"attempt_id"`
	QuestionID string         `json:"questionId" db:"question_id"`
	Answers    pq.StringArray `json:"answers" db:"answers"`
}
