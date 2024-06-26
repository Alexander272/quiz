import { createListenerMiddleware, TypedStartListening } from '@reduxjs/toolkit'

import { apiSlice } from '@/app/apiSlice'
import { AppDispatch, RootState } from '@/app/store'
import { resetUser } from '@/features/user/userSlice'
// import { resetTable } from '@/features/table/tableSlice'
// import { resetModal } from '@/features/modal/modalSlice'

export const resetStoreListener = createListenerMiddleware()

const startResetStoreListener = resetStoreListener.startListening as TypedStartListening<RootState, AppDispatch>

startResetStoreListener({
	actionCreator: resetUser,
	effect: async (_, listenerApi) => {
		await listenerApi.delay(100)
		//TODO
		// listenerApi.dispatch(resetTable())
		// listenerApi.dispatch(resetModal())
		listenerApi.dispatch(apiSlice.util.resetApiState())
	},
})
