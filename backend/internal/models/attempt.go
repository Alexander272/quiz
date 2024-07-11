package models

type Attempt struct {
	ID          string `json:"id" db:"id"`
	ScheduleID  string `json:"scheduleId" db:"schedule_id"`
	UserID      string `json:"userId" db:"user_id"`
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
type GetAttemptByID struct {
	ID string `json:"id" db:"id"`
}

type AttemptDTO struct {
	ID          string `json:"id" db:"id"`
	ScheduleID  string `json:"scheduleId" db:"schedule_id"`
	UserID      string `json:"userId" db:"user_id"`
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
