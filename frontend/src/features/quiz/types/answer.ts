export interface IAnswer {
	id: string
	questionId: string
	number: number
	text: string
	image?: string
	isCorrect?: boolean
}

export interface IAnswerDTO {
	id?: string
	questionId: string
	number: number
	text: string
	image: string
	isCorrect: boolean
}

export interface IAnswerForm {
	number: number
	text: string
	image?: File
	isCorrect: boolean
}
