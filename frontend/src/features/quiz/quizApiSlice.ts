import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IQuiz, IQuizDTO } from './types/quiz'
import { API } from '@/app/api'
import { apiSlice } from '@/app/apiSlice'

const quizApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getQuizzes: builder.query<{ data: IQuiz[] }, null>({
			query: () => API.Quiz.Index,
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					console.log(error)
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
			providesTags: [{ type: 'Quiz', id: 'ALL' }],
		}),
		getMyQuizzes: builder.query<{ data: IQuiz[] }, null>({
			query: () => API.Quiz.My,
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					console.log(error)
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
			providesTags: [{ type: 'Quiz', id: 'ALL' }],
		}),
		getQuizById: builder.query<{ data: IQuiz }, string>({
			query: id => ({
				url: `${API.Quiz.Index}/${id}`,
			}),
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					console.log(error)
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
			providesTags: (_res, _err, arg) => [{ type: 'Quiz', id: arg }],
		}),
		// getQuestions: builder.query()

		createQuiz: builder.mutation<{ id: string; message: string }, IQuizDTO>({
			query: data => ({
				url: API.Quiz.Index,
				method: 'POST',
				body: Object.entries(data).reduce((d, e) => (d.append(...e), d), new FormData()),
			}),
			invalidatesTags: [{ type: 'Quiz', id: 'ALL' }],
		}),
		updateQuiz: builder.mutation<{ message: string }, IQuizDTO>({
			query: data => ({
				url: `${API.Quiz.Index}/${data.id}`,
				method: 'PUT',
				body: Object.entries(data).reduce((d, e) => (d.append(...e), d), new FormData()),
			}),
			invalidatesTags: (_res, _err, arg) => [
				{ type: 'Quiz', id: arg.id },
				{ type: 'Quiz', id: 'ALL' },
				{ type: 'Question', id: arg.id },
			],
		}),
		deleteQuiz: builder.mutation<void, string>({
			query: id => ({
				url: `${API.Quiz.Index}/${id}`,
				method: 'DELETE',
			}),
			invalidatesTags: [{ type: 'Quiz', id: 'ALL' }],
		}),
	}),
})

export const {
	useGetQuizzesQuery,
	useGetMyQuizzesQuery,
	useGetQuizByIdQuery,
	useCreateQuizMutation,
	useUpdateQuizMutation,
	useDeleteQuizMutation,
} = quizApiSlice
