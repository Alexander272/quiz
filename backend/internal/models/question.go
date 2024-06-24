package models

import "time"

type Question struct {
	ID          string        `json:"id" db:"id"`
	QuizID      string        `json:"quizId" db:"quiz_id"`
	Text        string        `json:"text" db:"text"`
	Description string        `json:"description,omitempty" db:"description"`
	Image       string        `json:"image,omitempty" db:"image"`
	HasShuffle  bool          `json:"hasShuffle" db:"has_shuffle"` // перемешивать ответы
	Level       string        `json:"level" db:"level"`            //? сложность вопроса
	Points      int           `json:"points" db:"points"`          //? очки которые начислятся за правильный ответ
	GroupID     string        `json:"groupId" db:"group_id"`       //? чтобы можно было формировать группы с вопросами и уже в этих группах перемешивать вопросы и ограничивать количество выводимых вопросов из группы
	Time        time.Duration `json:"time" db:"time"`              //? время для дачи ответа
}
