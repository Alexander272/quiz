export interface IQuiz {
	id: string
	title: string
	description: string
	image?: string
	isDrawing: boolean
	numberOfAttempts: number
	categoryId: string
	startTime: number
	endTime: number
	hasShuffle: boolean
	hasSkippable: boolean
	showList: boolean
	showAnswers: boolean
	showResults: boolean
	time: number
	authorId: string
}

export interface IQuizDTO {
	id?: string
	title: string
	description: string
	image?: File
	imageLink?: string
	numberOfAttempts: number
	// categoryId: string
	// startTime: number
	// endTime: number
	hasShuffle: boolean
	hasSkippable: boolean
	showList: boolean
	showAnswers: boolean
	showResults: boolean
	time: number
	// authorId: string
}

export interface IQuizForm {
	title: string
	description: string
	image?: File | string
	numberOfAttempts: number
	hasShuffle: boolean
	hasSkippable: boolean
	showList: boolean
	showAnswers: boolean
	showResults: boolean
	time: number
}
