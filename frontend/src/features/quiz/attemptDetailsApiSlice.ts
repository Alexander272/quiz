import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IAttemptDetailDTO, IAttemptDetails, IGetAttemptDetailsDTO } from './types/attemptDetails'
import { API } from '@/app/api'
import { apiSlice } from '@/app/apiSlice'
export const attemptDetailsApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getAttemptDetails: builder.query<{ data: IAttemptDetails[] }, IGetAttemptDetailsDTO>({
			query: data => ({
				url: `${API.Attempt}/${data.attemptId}/details`,
				params: data.showAnswers ? new URLSearchParams({ showAnswers: 'true' }) : undefined,
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
		}),

		createAttemptDetail: builder.mutation<{ data: IAttemptDetails }, IAttemptDetailDTO>({
			query: data => ({
				url: API.Attempt.Details,
				method: 'POST',
				body: data,
			}),
		}),

		updateAttemptDetail: builder.mutation<{ data: IAttemptDetails }, IAttemptDetailDTO>({
			query: data => ({
				url: `${API.Attempt.Details}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
		}),
	}),
})

export const { useGetAttemptDetailsQuery, useCreateAttemptDetailMutation, useUpdateAttemptDetailMutation } =
	attemptDetailsApiSlice
