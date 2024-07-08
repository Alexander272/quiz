export interface IAnswer {
	id: string
	questionId: string
	number: number
	text: string
	image?: string
	isCorrect: boolean
}

export interface IAnswerDTO {
	id?: string
	// questionId: string
	number: number
	text: string
	// image?: string
	isCorrect: boolean
}

export interface IAnswerForm {
	id?: string
	number: number
	text: string
	// image?: File | string
	isCorrect: boolean
}
