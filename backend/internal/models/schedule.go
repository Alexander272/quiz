package models

type Schedule struct {
	ID               string `json:"id" db:"id"`
	QuizID           string `json:"quizId" db:"quiz_id"`
	StartTime        int64  `json:"startTime" db:"start_time"`
	EndTime          int64  `json:"endTime" db:"end_time"`
	NumberOfAttempts int    `json:"numberOfAttempts" db:"number_of_attempts"`
	// NumberOfTries    int    `json:"numberOfTries" db:"number_of_tries"`
}

type GetSchedule struct {
	Time int64 `json:"time"`
}

type GetScheduleByQuiz struct {
	QuizID string `json:"quizId" db:"quiz_id"`
}

type ScheduleDTO struct {
	ID               string `json:"id" db:"id"`
	QuizID           string `json:"quizId" db:"quiz_id"`
	StartTime        int64  `json:"startTime" db:"start_time"`
	EndTime          int64  `json:"endTime" db:"end_time"`
	NumberOfAttempts int    `json:"numberOfAttempts" db:"number_of_attempts"`
}

type DeleteScheduleDTO struct {
	ID string `json:"id" db:"id"`
}
