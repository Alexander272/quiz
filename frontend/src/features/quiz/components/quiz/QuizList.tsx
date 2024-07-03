import { useGetQuizzesQuery } from '../../quizApiSlice'

export const QuizList = () => {
	const { data, isFetching } = useGetQuizzesQuery(null)

	console.log('fetching', data, isFetching)

	return <div>QuizList</div>
}
