export interface IAttemptDetails {
	attemptId: string
	questions: IAttemptQuestion[]
}
export interface IAttemptQuestion {
	id: string
	answers: string[]
	correct: string[]
	isCorrect: boolean
	points: number
}

export interface IGetAttemptDetailsDTO {
	attemptId: string
	showAnswers: boolean
}

export interface IAttemptDetailDTO {
	id?: string
	attemptId: string
	questionId: string
	answers: string[]
}
