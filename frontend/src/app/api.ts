export const API = Object.freeze({
	Auth: {
		SignIn: `auth/sign-in` as const,
		SignOut: `auth/sign-out` as const,
		Refresh: `auth/refresh` as const,
	},
	Quiz: {
		Index: 'quizzes' as const,
		My: 'quizzes/my' as const,
	},
	Question: 'questions' as const,
	Attempt: {
		Base: 'attempts' as const,
		Save: 'attempts/save' as const,
		Details: 'attempts/details' as const,
	},
	// AttemptDetails: 'attempts/details' as const,
	Media: 'media' as const,
})
