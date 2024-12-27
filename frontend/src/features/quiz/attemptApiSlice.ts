import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IAttempt, IAttemptDTO } from './types/attempt'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'

const attemptApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getActiveAttempt: builder.query<{ data: IAttempt[] }, { quiz: string }>({
			query: data => ({
				url: API.Attempt,
				params: new URLSearchParams({ quiz: data.quiz, active: 'true' }),
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
			// providesTags: (_res, _err, arg) => [{ type: 'Question', id: arg }],
		}),

		createAttempt: builder.mutation<{ id: string }, IAttemptDTO>({
			query: data => ({
				url: API.Attempt,
				method: 'POST',
				body: data,
			}),
		}),
		updateAttempt: builder.mutation({
			query: data => ({
				url: `${API.Attempt}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
		}),
	}),
})

export const { useGetActiveAttemptQuery, useCreateAttemptMutation } = attemptApiSlice
