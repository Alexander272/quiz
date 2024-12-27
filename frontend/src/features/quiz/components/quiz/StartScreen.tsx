import { FC } from 'react'
import { Box, Button, Stack, Typography, useTheme } from '@mui/material'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IQuiz } from '../../types/quiz'
import { Image } from '@/features/files/components/UploadImage/Image'
import { StartIcon } from '@/components/Icons/StartIcon'
import { useCreateAttemptMutation, useGetActiveAttemptQuery } from '../../attemptApiSlice'

type Props = {
	quiz?: IQuiz
	questions: number
	onStart?: (id: string) => void
}

export const StartScreen: FC<Props> = ({ quiz, questions, onStart }) => {
	const { palette } = useTheme()

	const { data } = useGetActiveAttemptQuery({ quiz: quiz?.id || '' }, { skip: !quiz })
	const [create] = useCreateAttemptMutation()

	const startHandler = async () => {
		if (!quiz) return
		console.log('start')

		if (!data?.data.length) {
			try {
				const newData = { scheduleId: quiz.scheduleId, total: questions }
				const payload = await create(newData).unwrap()

				onStart && onStart(payload.id)
			} catch (error) {
				const fetchError = error as IFetchError
				toast.error(fetchError.data.message, { autoClose: false })
			}
		} else {
			onStart && onStart(data.data[0].id)
		}
	}

	console.log('active attempt', data?.data)

	if (!quiz) return null
	return (
		<Stack
			spacing={2}
			width={'100%'}
			maxWidth={800}
			padding={2}
			my={'auto'}
			borderRadius={3}
			paddingX={3}
			paddingY={3}
			border={'1px solid rgba(0, 0, 0, 0.12)'}
			sx={{ backgroundColor: '#fff' }}
		>
			<Stack spacing={4}>
				{quiz.image ? (
					<Box display={'flex'} maxWidth={400} maxHeight={400}>
						<Image src={quiz.image} alt={quiz.image} />
					</Box>
				) : null}

				<Stack>
					<Typography fontSize={'1.4rem'} align='center' mb={1}>
						{quiz.title}
					</Typography>
					<Typography align='justify' mb={2}>
						{quiz.description}
					</Typography>

					<Typography>
						Количество вопросов:{' '}
						<Typography component={'span'} fontWeight={'bold'}>
							{questions}
						</Typography>
					</Typography>
					<Typography>
						Время:{' '}
						<Typography component={'span'} fontWeight={'bold'}>
							{quiz.time > 0 ? quiz.time : 'Неограниченно'}
						</Typography>
					</Typography>

					<Button
						onClick={startHandler}
						variant='outlined'
						fullWidth
						sx={{ mt: 3, maxWidth: 260, mx: 'auto' }}
					>
						{(data?.data.length || 0) > 0 ? 'Продолжить' : 'Начать'}
						<StartIcon fontSize={18} fill={palette.primary.main} ml={1} />
					</Button>
				</Stack>
			</Stack>
		</Stack>
	)
}
