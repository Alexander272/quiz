import { Stack } from '@mui/material'

import { useGetQuizzesQuery } from '@/features/quiz/quizApiSlice'
import { Item } from './Item'

export const QuizList = () => {
	const { data, isFetching } = useGetQuizzesQuery(null)

	console.log('fetching', data, isFetching)

	return (
		<Stack spacing={2}>
			{data?.data.map(d => (
				<Item key={d.id} data={d} />
			))}
		</Stack>
	)
}
