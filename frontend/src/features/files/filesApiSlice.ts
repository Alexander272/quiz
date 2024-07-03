import type { IUploadFiles } from './types/files'
import { HttpCodes } from '@/constants/httpCodes'
import { apiSlice } from '@/app/apiSlice'

const filesApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		uploadFiles: builder.mutation<null, IUploadFiles>({
			query: data => ({
				url: '',
				method: 'POST',
				body: data.data,
				validateStatus: response => response.status === HttpCodes.CREATED,
			}),
			// invalidatesTags: []
		}),
	}),
})

export const { useUploadFilesMutation } = filesApiSlice
