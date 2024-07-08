package models

type UserAnswer struct {
	QuizID     string   `json:"quizId" db:"quiz_id"`
	QuestionID string   `json:"questionId" db:"question_id"`
	Answer     []string `json:"answer" db:"answer"`
}

type GetResults struct {
	QuizID string `json:"quizId"`
}

type Result struct {
	QuestionID    string   `json:"questionId" db:"question_id"`
	IsCorrect     bool     `json:"isCorrect" db:"is_correct"`
	UserAnswer    []string `json:"userAnswer" db:"answer"`
	CorrectAnswer []string `json:"correctAnswer" db:"correct"`
}

type UserResult struct {
	QuizID  string `json:"quizId"`
	Time    string `json:"time"`
	Points  int    `json:"points"`
	Correct int    `json:"correct"`
	Total   int    `json:"total"`
}
