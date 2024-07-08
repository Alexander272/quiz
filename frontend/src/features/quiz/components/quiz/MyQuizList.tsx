import { Card, CardContent, CardMedia, Stack, Typography } from '@mui/material'

import { useGetMyQuizzesQuery } from '../../quizApiSlice'

export const MyQuizList = () => {
	const { data, isFetching } = useGetMyQuizzesQuery(null)

	console.log(data?.data, isFetching)

	return (
		<Stack direction={'row'} spacing={3}>
			{data?.data.map(d => (
				<Card key={d.id} sx={{ width: '100%', maxWidth: 345 }}>
					<CardMedia component='img' height='140' image={d.image} alt={d.image} />
					<CardContent>
						<Typography gutterBottom variant='h5' component='div'>
							{d.title}
						</Typography>
						<Typography variant='body2' color='text.secondary'>
							{d.description}
						</Typography>
					</CardContent>
					{/* //TODO добавить кнопку для просмотра результатов и предпросмотр теста */}
				</Card>
			))}
		</Stack>
	)
}
