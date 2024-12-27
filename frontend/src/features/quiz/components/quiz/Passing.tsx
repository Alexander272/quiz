import { FC, useEffect, useState } from 'react'
import { Button, Stack, Typography, useTheme } from '@mui/material'
import { FormProvider, useForm } from 'react-hook-form'

import type { IQuestion } from '../../types/question'
import type { IQuiz, IUserQuizForm } from '../../types/quiz'
import { useDebounce } from '@/hooks/useDebounce'
import { ShortDivider } from '@/components/Divider/ShortDivider'
import { LongArrowIcon } from '@/components/Icons/LongArrowIcon'
import { FinishIcon } from '@/components/Icons/FinishIcon'
import { QuestionList } from '../question/QuestionList'
import { Question } from '../question/Question'

type Props = {
	attemptId: string
	quiz?: IQuiz
	questions: IQuestion[]
}

export const Passing: FC<Props> = ({ attemptId, quiz, questions }) => {
	const methods = useForm<IUserQuizForm>()

	const { palette } = useTheme()

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
			id: k,
			answers: Object.keys(form[k].answers).filter(a => form[k].answers[a]),
		}))

		const newData = {
			// quizId: quiz?.id,
			attemptId: attemptId,
			questions: questions,
		}
		console.log('newData', newData)
		console.log('dirty', methods.formState.dirtyFields)
	})

	const selectHandler = (index: number) => {
		setCurQuestion(index)
	}

	if (!quiz) return null
	return (
		<>
			<Typography fontSize={'1.4rem'} mb={1}>
				{quiz?.title}
			</Typography>
			<ShortDivider sx={{ mb: 4 }} />

			{/* //TODO надо сделать экран для старта теста */}

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

						{questions ? <Question data={questions[curQuestion]} /> : null}

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

							{curQuestion != questions.length - 1 ? (
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
								<Button onClick={saveHandler} variant='contained' sx={{ width: '100%', maxWidth: 200 }}>
									Завершить
									<FinishIcon fontSize={18} ml={1} fill={'#fff'} />
								</Button>
							)}
						</Stack>
					</Stack>
					<Stack width={'100%'} maxWidth={400}>
						<QuestionList data={questions} active={curQuestion} onSelect={selectHandler} />
					</Stack>
				</Stack>
			</FormProvider>
		</>
	)
}
