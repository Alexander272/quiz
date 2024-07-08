import { FC } from 'react'
import { Stack } from '@mui/material'

import { QuestionForm } from './QuestionForm'
import { useGetQuestionsQuery } from '../../questionApiSlice'
import { Fallback } from '@/components/Fallback/Fallback'

type Props = {
	quizId: string
}

export const QuestionListForm: FC<Props> = ({ quizId }) => {
	// const { palette } = useTheme()

	// const [newQuestions, setNewQuestions] = useState([])

	const { data, isFetching } = useGetQuestionsQuery(
		{ quiz: quizId, answers: true, shuffle: false },
		{ skip: !quizId }
	)

	if (!quizId) return null
	return (
		<Stack spacing={2} position={'relative'}>
			{data?.data.map(d => (
				<QuestionForm key={d.id} method='update' data={d} quizId={quizId} />
			))}
			<QuestionForm quizId={quizId} method='create' position={(data?.data.length || 0) + 1} />

			{isFetching && (
				<Fallback
					position={'absolute'}
					left={0}
					top={0}
					background={'#c6e5ff42'}
					zIndex={5}
					margin={'0!important'}
				/>
			)}

			{/* <Button variant='outlined' sx={{ width: 300, mt: 3, textTransform: 'inherit', mx: 'auto!important' }}>
				<PlusIcon fill={palette.primary.main} fontSize={14} mr={1} />
				Добавить вопрос
			</Button> */}
		</Stack>
	)
}
