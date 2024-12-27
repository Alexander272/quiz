export interface IAttempt {
	id: string
	scheduleId: string
	userId: string
	startTime: number
	endTime: number
	correct: number
	total: number
	points: number
	totalPoints: number
}

export interface IAttemptDTO {
	id?: string
	scheduleId: string
	startTime?: number
	endTime?: number
}
