import { IAnswer, IAnswerForm } from './answer'

export interface IQuestion {
	id: string
	number: number
	quizId: string
	text: string
	description?: string
	image?: string
	hasShuffle: boolean
	level: string
	points: number
	time: string
	answers?: (IAnswer | undefined)[]
}

export interface IQuestionDTO {
	id?: string
	number: number
	quizId: string
	text: string
	description: string
	image: string
	hasShuffle: boolean
	level: string
	points: number
	time: string
}

export interface IGetQuestionDTO {
	quizId: string
	hasAnswers: boolean
	hasShuffle: boolean
}

export interface IQuestionForm {
	number: number
	text: string
	description: string
	image?: File
	hasShuffle: boolean
	level: string
	points: number
	time: string
	answers: IAnswerForm[]
}
