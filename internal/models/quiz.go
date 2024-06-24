package models

import "time"

type Quiz struct {
	ID               string        `json:"id" db:"id"`
	Title            string        `json:"title" db:"title"`
	Description      string        `json:"description,omitempty" db:"description"`
	Image            string        `json:"image,omitempty" db:"image"`
	IsDrawing        bool          `json:"isDrawing" db:"is_drawing"`
	NumberOfAttempts uint8         `json:"numberOfAttempts" db:"number_of_attempts"`
	CategoryID       string        `json:"categoryId" db:"category_id"`
	StartTime        int64         `json:"startTime" db:"start_time"`
	EndTime          int64         `json:"endTime" db:"end_time"`
	HasShuffle       bool          `json:"hasShuffle" db:"has_shuffle"`     // перемешивать вопросы
	HasSkippable     bool          `json:"hasSkippable" db:"has_skippable"` // можно пропускать вопросы
	ShowList         bool          `json:"showList" db:"show_list"`         // показывать все вопросы, а не по одному
	Time             time.Duration `json:"time" db:"time"`                  //? время для выполнения теста
	AuthorID         string        `json:"authorId" db:"author_id"`
}

type GetQuizDTO struct {
	ID string `json:"id" db:"id"`
}

type QuizDTO struct {
	ID               string        `json:"id"`
	Title            string        `json:"title" binding:"required,min=3,max=255"`
	Description      string        `json:"description,omitempty"`
	Image            string        `json:"image,omitempty"`
	NumberOfAttempts uint8         `json:"numberOfAttempts" binding:"min=0,max=250"`
	CategoryID       string        `json:"categoryId"`
	StartTime        int64         `json:"startTime"`
	EndTime          int64         `json:"endTime"`
	HasShuffle       bool          `json:"hasShuffle"`   // перемешивать вопросы
	HasSkippable     bool          `json:"hasSkippable"` // можно пропускать вопросы
	ShowList         bool          `json:"showList"`     // показывать все вопросы, а не по одному
	Time             time.Duration `json:"time"`         //? время для выполнения теста
	AuthorID         string        `json:"authorId"`
}

type DeleteQuizDTO struct {
	ID string `json:"id"`
}
