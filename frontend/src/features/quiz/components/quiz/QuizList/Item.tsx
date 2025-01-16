import { FC } from 'react'
import { Button, Stack, Typography } from '@mui/material'
import { useNavigate } from 'react-router-dom'

import type { IQuiz } from '@/features/quiz/types/quiz'
import { AppRoutes } from '@/constants/routes'

type Props = {
	data: IQuiz
}

export const Item: FC<Props> = ({ data }) => {
	const navigate = useNavigate()

	const selectHandler = () => {
		navigate(AppRoutes.Quiz.replace(':id', data.id))
	}

	return (
		<Stack direction={'row'} sx={{ border: '1px solid #E0E0E0', borderRadius: 2 }}>
			{/* Image */}
			<Stack spacing={1}>
				<Typography fontSize={'1.2rem'} fontWeight={'bold'}>
					{data.title}
				</Typography>
				{data.description && <Typography>{data.description}</Typography>}
			</Stack>
			<Stack>
				<Typography>{/* Time */}</Typography>
				<Button variant='outlined' onClick={selectHandler}>
					Начать //
				</Button>
			</Stack>
		</Stack>
	)
}
