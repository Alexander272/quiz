import type { PropsWithChildren } from 'react'
import { Navigate, Outlet, useLocation } from 'react-router-dom'

import { AppRoutes } from '@/constants/routes'
import { useAppSelector } from '@/hooks/redux'
import { getToken } from '@/features/user/userSlice'
// import { Forbidden } from '../forbidden/ForbiddenLazy'

// проверка авторизации пользователя
export default function PrivateRoute({ children }: PropsWithChildren) {
	const token = useAppSelector(getToken)
	// const menu = useAppSelector(getMenu)
	const location = useLocation()

	if (!token) return <Navigate to={AppRoutes.Auth} state={{ from: location }} />
	// if (!menu || !menu.length) return <Forbidden />

	if (!children) return <Outlet />
	return children
}
