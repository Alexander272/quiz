// import { toast } from 'react-toastify'

// import type { IBaseFetchError } from '@/app/types/error'
import type { IUploadFiles } from './types/files'
import { HttpCodes } from '@/constants/httpCodes'
import { API } from '@/app/api'
import { apiSlice } from '@/app/apiSlice'

const filesApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		uploadFiles: builder.mutation<{ data: string }, IUploadFiles>({
			query: data => ({
				url: API.Media,
				method: 'POST',
				body: data.data,
				validateStatus: response => response.status === HttpCodes.CREATED,
			}),
			// onQueryStarted: async (_arg, api) => {
			// 	try {
			// 		await api.queryFulfilled
			// 	} catch (error) {
			// 		console.log(error)
			// 		const fetchError = (error as IBaseFetchError).error
			// 		toast.error(fetchError.data.message, { autoClose: false })
			// 	}
			// },
			// invalidatesTags: []
		}),
		deleteFile: builder.mutation<null, string>({
			query: path => ({
				url: API.Media,
				method: 'DELETE',
				params: new URLSearchParams({ path: path }),
			}),
			// onQueryStarted: async (_arg, api) => {
			// 	try {
			// 		await api.queryFulfilled
			// 	} catch (error) {
			// 		console.log(error)
			// 		const fetchError = (error as IBaseFetchError).error
			// 		toast.error(fetchError.data.message, { autoClose: false })
			// 	}
			// },
		}),
	}),
})

export const { useUploadFilesMutation, useDeleteFileMutation } = filesApiSlice
