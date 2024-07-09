import type { RouteObject } from 'react-router-dom'
import { createBrowserRouter } from 'react-router-dom'

import { AppRoutes } from '@/constants/routes'
import { Layout } from '@/components/Layout/Layout'
import { NotFound } from '@/pages/notFound/NotFoundLazy'
import { Auth } from '@/pages/auth/AuthLazy'
import { Home } from '@/pages/home/HomeLazy'
import { Quizzes } from '@/pages/quiz/QuizzesLazy'
import { MyQuizzes } from '../quiz/my/MyQuizzesLazy'
import { Quiz } from '../quiz/start/QuizLazy'
import { Create } from '@/pages/quiz/create/CreateLazy'
import { Edit } from '@/pages/quiz/edit/EditLazy'
import PrivateRoute from './PrivateRoute'

const config: RouteObject[] = [
	{
		element: <Layout />,
		errorElement: <NotFound />,
		children: [
			{
				path: AppRoutes.Auth,
				element: <Auth />,
			},
			{
				element: <PrivateRoute />,
				children: [
					{
						path: AppRoutes.Home,
						element: <Home />,
					},
					{
						path: AppRoutes.Home + '/:id',
						element: <Quiz />,
					},
					{
						path: AppRoutes.Quizzes.Index,
						element: <Quizzes />,
					},
					{
						path: AppRoutes.Quizzes.My,
						element: <MyQuizzes />,
					},
					{
						path: AppRoutes.Quizzes.Create,
						element: <Create />,
					},
					{
						path: AppRoutes.Quizzes.Edit + ':id',
						element: <Edit />,
					},
				],
			},
		],
	},
]

export const router = createBrowserRouter(config)
