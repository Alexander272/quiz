import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IGetQuestionDTO, IQuestion, IQuestionDTO } from './types/question'
import { API } from '@/app/api'
import { apiSlice } from '@/app/apiSlice'

const questionApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getQuestions: builder.query<{ data: IQuestion[] }, IGetQuestionDTO>({
			query: data => ({
				url: API.Question,
				params: new URLSearchParams(
					Object.entries(data).filter(item => {
						if (item[0] == 'answers' && item[1] == false) return false
						if (item[0] == 'shuffle' && item[1] == true) return false

						return true
					})
				),
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
			providesTags: (_res, _err, arg) => [{ type: 'Question', id: arg.quiz }],
		}),
		getQuestionById: builder.query<{ data: IQuestion }, string>({
			query: id => ({
				url: `${API.Question}/${id}`,
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
			providesTags: (_res, _err, arg) => [{ type: 'Question', id: arg }],
		}),

		createQuestion: builder.mutation<void, IQuestionDTO>({
			query: data => ({
				url: API.Question,
				method: 'POST',
				body: data,
			}),
			invalidatesTags: (_res, _err, arg) => [{ type: 'Question', id: arg.quizId }],
		}),
		updateQuestion: builder.mutation<void, IQuestionDTO>({
			query: data => ({
				url: `${API.Question}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
			invalidatesTags: (_res, _err, arg) => [
				{ type: 'Question', id: arg.id },
				{ type: 'Question', id: arg.quizId },
			],
		}),
		deleteQuestion: builder.mutation<void, { quizId: string; id: string }>({
			query: data => ({
				url: `${API.Question}/${data.id}`,
				method: 'DELETE',
				params: new URLSearchParams({ quizId: data.quizId }),
			}),
			invalidatesTags: (_res, _err, arg) => [
				{ type: 'Question', id: arg.id },
				{ type: 'Question', id: arg.quizId },
			],
		}),
	}),
})

export const {
	useGetQuestionsQuery,
	useGetQuestionByIdQuery,
	useCreateQuestionMutation,
	useUpdateQuestionMutation,
	useDeleteQuestionMutation,
} = questionApiSlice
