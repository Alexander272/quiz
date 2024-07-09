import { FC, useEffect } from 'react'
import { Box, Stack, Typography } from '@mui/material'
import { useParams } from 'react-router-dom'
import { FormProvider, useForm } from 'react-hook-form'

import type { IUserQuizForm } from '../../types/quiz'
import { Fallback } from '@/components/Fallback/Fallback'
import { ShortDivider } from '@/components/Divider/ShortDivider'
import { useGetQuizByIdQuery } from '../../quizApiSlice'
import { useGetQuestionsQuery } from '../../questionApiSlice'
import { Question } from '../question/Question'
import { useDebounce } from '@/hooks/useDebounce'

type Props = unknown

export const Quiz: FC<Props> = () => {
	const { id } = useParams()
	const methods = useForm<IUserQuizForm>()

	const { data: quiz, isFetching: isFetchQuiz } = useGetQuizByIdQuery(id || '', { skip: !id })
	const { data: questions, isFetching: isFetchQuestion } = useGetQuestionsQuery(
		{ quiz: id || '', answers: true },
		{ skip: !id }
	)

	const save = useDebounce(() => {
		saveHandler()
	}, 30000)

	useEffect(() => {
		if (methods.formState.dirtyFields) save()
	}, [methods.formState.dirtyFields, save])

	const saveHandler = methods.handleSubmit(form => {
		console.log('quiz', form)

		const questions = Object.keys(form).map(k => ({
			questionId: k,
			answers: Object.keys(form[k].answers).filter(a => form[k].answers[a]),
		}))
		console.log(questions)

		// const newData = {
		// 	quizId: id,
		// 	questions:questions,
		// }
	})

	if (isFetchQuestion || isFetchQuiz) return <Fallback />
	if (!questions) return null
	return (
		<Box sx={{ userSelect: 'none' }} display={'flex'} flexDirection={'column'} alignItems={'center'}>
			<Typography fontSize={'1.4rem'} mb={1}>
				{quiz?.data.title}
			</Typography>
			<ShortDivider sx={{ mb: 4 }} />

			{/* current question / total question */}
			{/* Timer */}
			{/* QuestionsList (aside) */}
			{/* Question */}

			<FormProvider {...methods}>
				<Stack direction={'row'} spacing={2} width={'100%'} maxWidth={1000}>
					<Stack spacing={2} width={'100%'} maxWidth={1000}>
						<Question data={questions?.data[0]} />
					</Stack>
					{/* <Stack></Stack> */}
				</Stack>
			</FormProvider>
		</Box>
	)
}
