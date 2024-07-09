import { IAnswer, IAnswerDTO, IAnswerForm } from './answer'

export interface IQuestion {
	id: string
	number: number
	quizId: string
	text: string
	description: string
	image?: string
	hasShuffle: boolean
	level: string
	points: number
	time: number
	answers: IAnswer[]
}

export interface IQuestionDTO {
	id?: string
	number: number
	quizId: string
	text: string
	description?: string
	// image?: File | string
	image?: string
	hasShuffle: boolean
	level: string
	points: number
	time: number
	answers: IAnswerDTO[]
}

export interface IGetQuestionDTO {
	quiz: string
	answers?: boolean
	shuffle?: boolean
}

export interface IQuestionForm {
	number: number
	text: string
	description: string
	image?: File | string
	hasShuffle: boolean
	level: string
	points: number
	time: number
	answers: IAnswerForm[]
}
