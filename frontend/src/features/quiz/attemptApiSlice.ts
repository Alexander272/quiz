import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IAttempt, IAttemptDTO } from './types/attempt'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'
import { IAttemptDetailDTO } from './types/attemptDetails'

const attemptApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getActiveAttempt: builder.query<{ data: IAttempt[] }, { quiz: string }>({
			query: data => ({
				url: API.Attempt.Base,
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
				url: API.Attempt.Base,
				method: 'POST',
				body: data,
			}),
		}),
		updateAttempt: builder.mutation({
			query: data => ({
				url: `${API.Attempt.Base}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
		}),

		saveAttempt: builder.mutation<null, IAttemptDetailDTO[]>({
			query: data => ({
				url: API.Attempt.Save,
				method: 'POST',
				body: data,
			}),
			// invalidatesTags
		}),
	}),
})

export const { useGetActiveAttemptQuery, useCreateAttemptMutation, useSaveAttemptMutation } = attemptApiSlice
