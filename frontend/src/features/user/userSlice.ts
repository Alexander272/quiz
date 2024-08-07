import { PayloadAction, createSlice } from '@reduxjs/toolkit'

import type { RootState } from '@/app/store'
import type { IUser } from './types/user'

export interface IUserState {
	id: string | null
	name?: string
	role: string | null
	menu: string[]
	token: string | null
}

const initialState: IUserState = {
	id: null,
	role: null,
	token: null,
	menu: [],
}

const userSlice = createSlice({
	name: 'user',
	initialState,
	reducers: {
		setUser: (state, action: PayloadAction<IUser>) => {
			state.id = action.payload.id
			state.name = action.payload.name
			state.role = action.payload.role
			state.menu = action.payload.menu
			state.token = action.payload.token
		},

		resetUser: () => initialState,
	},
})

export const getToken = (state: RootState) => state.user.token
export const getMenu = (state: RootState) => state.user.menu
export const getRole = (state: RootState) => state.user.role

export const userPath = userSlice.name
export const userReducer = userSlice.reducer

export const { setUser, resetUser } = userSlice.actions
