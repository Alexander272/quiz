export const API = Object.freeze({
	Auth: {
		SignIn: `auth/sign-in`,
		SignOut: `auth/sign-out`,
		Refresh: `auth/refresh`,
	},
	Quiz: {
		Index: 'quizzes',
		My: 'quizzes/my',
	},
	Question: 'questions',
	Attempt: 'attempts',
	AttemptDetails: 'attempts/details',
	Media: 'media',
})
