import { FC, useEffect, useState } from 'react'
import { Box, Button, Stack, Typography, useTheme } from '@mui/material'
import { useParams } from 'react-router-dom'
import { FormProvider, useForm } from 'react-hook-form'

import type { IUserQuizForm } from '../../types/quiz'
import { useDebounce } from '@/hooks/useDebounce'
import { Fallback } from '@/components/Fallback/Fallback'
import { ShortDivider } from '@/components/Divider/ShortDivider'
import { LongArrowIcon } from '@/components/Icons/LongArrowIcon'
import { useGetQuizByIdQuery } from '../../quizApiSlice'
import { useGetQuestionsQuery } from '../../questionApiSlice'
import { QuestionList } from '../question/QuestionList'
import { Question } from '../question/Question'

type Props = unknown

export const Quiz: FC<Props> = () => {
	const { id } = useParams()
	const methods = useForm<IUserQuizForm>()

	const { palette } = useTheme()

	const { data: quiz, isFetching: isFetchQuiz } = useGetQuizByIdQuery(id || '', { skip: !id })
	const { data: questions, isFetching: isFetchQuestion } = useGetQuestionsQuery(
		{ quiz: id || '', answers: true },
		{ skip: !id }
	)

	const [curQuestion, setCurQuestion] = useState(0)

	const autosave = useDebounce(() => {
		saveHandler()
	}, 30000)

	useEffect(() => {
		if (methods.formState.dirtyFields) autosave()
	}, [methods.formState.dirtyFields, autosave])

	const saveHandler = methods.handleSubmit(form => {
		console.log('quiz', form)

		const questions = Object.keys(form).map(k => ({
			questionId: k,
			answers: Object.keys(form[k].answers).filter(a => form[k].answers[a]),
		}))

		const newData = {
			quizId: id,
			questions: questions,
		}
		console.log(newData)
	})

	const selectHandler = (index: number) => {
		setCurQuestion(index)
	}

	// if (!questions) return null
	if (isFetchQuestion || isFetchQuiz) return <Fallback />
	return (
		<Box sx={{ userSelect: 'none' }} display={'flex'} flexDirection={'column'} alignItems={'center'}>
			<Typography fontSize={'1.4rem'} mb={1}>
				{quiz?.data.title}
			</Typography>
			<ShortDivider sx={{ mb: 4 }} />

			{/* current question / total question */}
			{/* Timer (aside) */}
			{/* QuestionsList (aside) */}
			{/* Question */}

			<FormProvider {...methods}>
				<Stack direction={'row'} spacing={2} width={'100%'} maxWidth={1400}>
					<Stack
						spacing={2}
						width={'100%'}
						maxWidth={1000}
						padding={2}
						borderRadius={3}
						paddingX={3}
						paddingY={3}
						border={'1px solid rgba(0, 0, 0, 0.12)'}
						sx={{ backgroundColor: '#fff' }}
						overflow={'hidden'}
					>
						{/* <Box
							padding={2}
							borderRadius={3}
							paddingX={2}
							paddingY={1}
							border={'1px solid rgba(0, 0, 0, 0.12)'}
							display={'flex'}
							justifyContent={'center'}
							sx={{ backgroundColor: '#fff' }}
						>
							<Typography fontSize={'1.5rem'}>{curQuestion + 1} / {questions.data.length}</Typography>
						</Box> */}

						{questions ? <Question data={questions?.data[curQuestion]} /> : null}

						<Stack spacing={2} direction={'row'} justifyContent={'center'}>
							<Button
								disabled={curQuestion == 0}
								onClick={() => selectHandler(curQuestion - 1)}
								sx={{ width: '100%', maxWidth: 200, ':hover': { svg: { right: 6 } } }}
							>
								<LongArrowIcon
									fontSize={18}
									transform={'rotate(180deg)'}
									fill={curQuestion == 0 ? palette.action.disabled : palette.primary.main}
									mr={1}
									position={'relative'}
									right={0}
									transition={'all .3s ease-in-out'}
								/>
								Предыдущий
							</Button>

							{curQuestion != (questions?.data.length || 0) - 1 ? (
								<Button
									onClick={() => selectHandler(curQuestion + 1)}
									variant='outlined'
									sx={{ width: '100%', maxWidth: 200, ':hover': { svg: { left: 6 } } }}
								>
									Следующий
									<LongArrowIcon
										fontSize={18}
										ml={1}
										fill={palette.primary.main}
										position={'relative'}
										left={0}
										transition={'all .3s ease-in-out'}
									/>
								</Button>
							) : (
								<Button
									onClick={saveHandler}
									variant='contained'
									sx={{ width: '100%', maxWidth: 200, ':hover': { svg: { left: 6 } } }}
								>
									Завершить
									{/* //TODO подумать о другой иконке */}
									<LongArrowIcon
										fontSize={18}
										ml={1}
										fill={'#fff'}
										position={'relative'}
										left={0}
										transition={'all .3s ease-in-out'}
									/>
								</Button>
							)}
						</Stack>
					</Stack>
					<Stack width={'100%'} maxWidth={400}>
						<QuestionList data={questions?.data || []} active={curQuestion} onSelect={selectHandler} />
					</Stack>
				</Stack>
			</FormProvider>
		</Box>
	)
}
