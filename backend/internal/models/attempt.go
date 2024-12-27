package models

type Attempt struct {
	ID          string `json:"id" db:"id"`
	ScheduleID  string `json:"scheduleId" db:"schedule_id"`
	UserID      string `json:"userId" db:"user_id"`
	Username    string `json:"user" db:"username"`
	StartTime   int64  `json:"startTime" db:"start_time"`
	EndTime     int64  `json:"endTime" db:"end_time"`
	Correct     int    `json:"correct" db:"correct"`
	Total       int    `json:"total" db:"total"`
	Points      int    `json:"points" db:"points"`
	TotalPoints int    `json:"totalPoints" db:"total_points"`
}

type GetAttempt struct {
	ScheduleID string `json:"scheduleId" db:"schedule_id"`
	UserID     string `json:"userId" db:"user_id"`
}
type GetAttemptByQuiz struct {
	UserID string `json:"userId" db:"user_id"`
	QuizID string `json:"quizId" db:"quiz_id"`
	Time   int64  `json:"time" db:"time"`
}
type GetAttemptByID struct {
	ID string `json:"id" db:"id"`
}

type AttemptDTO struct {
	ID          string `json:"id" db:"id"`
	Token       string // для получения username из keycloak
	ScheduleID  string `json:"scheduleId" db:"schedule_id"`
	UserID      string `json:"userId" db:"user_id"`
	Username    string `json:"user" db:"username"`
	StartTime   int64  `json:"startTime" db:"start_time"`
	EndTime     int64  `json:"endTime" db:"end_time"`
	Correct     int    `json:"correct" db:"correct"`
	Total       int    `json:"total" db:"total"`
	Points      int    `json:"points" db:"points"`
	TotalPoints int    `json:"totalPoints" db:"total_points"`
}

type DeleteAttemptDTO struct {
	ID string `json:"id" db:"id"`
}

type FinishAttempt struct {
	ID          string                `json:"id" db:"id"`
	QuizID      string                `json:"quizId" db:"quiz_id"`
	ScheduleID  string                `json:"scheduleId" db:"schedule_id"`
	StartTime   int64                 `json:"startTime" db:"start_time"`
	EndTime     int64                 `json:"endTime" db:"end_time"`
	Correct     int                   `json:"correct" db:"correct"`
	Total       int                   `json:"total" db:"total"`
	Points      int                   `json:"points" db:"points"`
	TotalPoints int                   `json:"totalPoints" db:"total_points"`
	ShowAnswers bool                  `json:"showAnswers" db:"show_answers"`
	Questions   []*AttemptQuestionDTO `json:"questions"`
}
