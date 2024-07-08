import { Box, Breadcrumbs, Stack, Typography } from '@mui/material'
import { useParams } from 'react-router-dom'

import { AppRoutes } from '@/constants/routes'
import { PageBox } from '@/styled/PageBox'
import { QuizForm } from '@/features/quiz/components/quiz/QuizForm'
import { Breadcrumb } from '@/components/Breadcrumb/Breadcrumb'

export default function Edit() {
	const { id } = useParams()
	console.log('quizId', id)

	return (
		<PageBox>
			<Box
				borderRadius={3}
				paddingX={3}
				paddingY={3}
				width={'60%'}
				marginX={'auto'}
				border={'1px solid rgba(0, 0, 0, 0.12)'}
				// flexGrow={1}
				height={'fit-content'}
				minHeight={200}
				mb={6}
				display={'flex'}
				flexDirection={'column'}
				sx={{ backgroundColor: '#fff', userSelect: 'none' }}
			>
				<Breadcrumbs aria-label='breadcrumb'>
					<Breadcrumb to={AppRoutes.Home}>Главная</Breadcrumb>
					<Breadcrumb to={AppRoutes.Quizzes.My}>Мои тесты</Breadcrumb>
					<Breadcrumb to={AppRoutes.Quizzes.Edit} active>
						Редактировать
					</Breadcrumb>
				</Breadcrumbs>

				<Stack direction={'row'} justifyContent={'center'} alignItems={'center'} mt={1} mb={3}>
					<Typography fontSize={'1.3rem'} textAlign={'center'}>
						Редактировать тест
					</Typography>

					{/* <Tooltip title='Очистить форму'>
						<Button onClick={clearHandler} sx={{ minWidth: 40, ml: 0.5 }}>
							<RefreshIcon fontSize={16} />
						</Button>
					</Tooltip> */}
				</Stack>

				{/* 'b9bdb0c6-a011-41b8-8375-3721655747c4' */}
				<QuizForm quizId={id} method={'update'} />
			</Box>
		</PageBox>
	)
}
