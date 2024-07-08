package models

import (
	"time"
)

type Question struct {
	ID          string        `json:"id" db:"id"`
	Number      int           `json:"number" db:"number"`
	QuizID      string        `json:"quizId" db:"quiz_id"`
	Text        string        `json:"text" db:"text"`
	Description string        `json:"description" db:"description"`
	Image       string        `json:"image,omitempty" db:"image"`
	HasShuffle  bool          `json:"hasShuffle" db:"has_shuffle"` // перемешивать ответы
	Level       string        `json:"level" db:"level"`            //? сложность вопроса
	Points      int           `json:"points" db:"points"`          //? очки которые начисляются за правильный ответ
	Time        time.Duration `json:"time" db:"time"`              //? время для дачи ответа
	// GroupID     string        `json:"groupId" db:"group_id"`       //? чтобы можно было формировать группы с вопросами и уже в этих группах перемешивать вопросы и ограничивать количество выводимых вопросов из группы
	Answers []*Answer `json:"answers"`
}

type GetQuestionsDTO struct {
	QuizID     string `json:"quizId" db:"quiz_id"`
	HasAnswers bool   `json:"hasAnswers"`
	HasShuffle bool   `json:"hasShuffle" db:"has_shuffle"`
}
type GetQuestionDTO struct {
	ID string `json:"id" db:"id"`
}

type QuestionDTO struct {
	ID          string `json:"id" db:"id"`
	Number      int    `json:"number" db:"number" binding:"required,min=1"`
	QuizID      string `json:"quizId" db:"quiz_id"`
	Text        string `json:"text" db:"text" binding:"required,min=3"`
	Description string `json:"description" db:"description"`
	// Image       *multipart.FileHeader `json:"image" db:"image_file"`
	Image      string        `json:"image" db:"image"`
	HasShuffle bool          `json:"hasShuffle" db:"has_shuffle"`
	Level      string        `json:"level" db:"level"`   //? сложность вопроса
	Points     int           `json:"points" db:"points"` //? очки которые начисляются за правильный ответ
	Time       time.Duration `json:"time" db:"time"`     //? время для дачи ответа
	// GroupID     string        `json:"groupId"` //? чтобы можно было формировать группы с вопросами и уже в этих группах перемешивать вопросы и ограничивать количество выводимых вопросов из группы
	Answers []*AnswerDTO `json:"answers"`
}

type DeleteQuestionDTO struct {
	ID     string `json:"id"`
	QuizID string `json:"quizId"`
}

type UserQuestion struct {
	ID      string   `json:"id"`
	Answers []string `json:"answers"`
}
