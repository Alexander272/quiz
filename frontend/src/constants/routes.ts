export const AppRoutes = {
	Home: '/' as const,
	Auth: '/auth' as const,
	Quiz: '/:id' as const,
	Quizzes: {
		Index: '/quizzes' as const,
		My: '/quizzes/my' as const,
		Create: '/quizzes/create' as const,
		Edit: '/quizzes/edit/' as const,
	},
}
