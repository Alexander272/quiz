import { FC, useState } from 'react'
import { Button, Checkbox, Divider, FormControlLabel, Stack, TextField, useTheme } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IQuizDTO, IQuizForm } from '../../types/quiz'
import { UploadImage } from '@/features/files/components/UploadImage/UploadImage'
import { Fallback } from '@/components/Fallback/Fallback'
import { SaveIcon } from '@/components/Icons/SaveIcon'
import { localKeys } from '../../constants/localKeys'
import { useCreateQuizMutation, useGetQuizByIdQuery, useUpdateQuizMutation } from '../../quizApiSlice'
import { QuestionListForm } from '../question/QuestionListForm'

const defaultValues: IQuizForm = {
	title: '',
	description: '',
	numberOfAttempts: 3,
	hasShuffle: true,
	hasSkippable: false,
	showList: false,
	showAnswers: false,
	showResults: false,
	time: 0,
}

type Props = {
	quizId?: string
	method: 'create' | 'update'
}

export const QuizForm: FC<Props> = ({ quizId = '', method }) => {
	const { palette } = useTheme()

	const { data, isFetching } = useGetQuizByIdQuery(quizId, { skip: !quizId || method == 'create' })
	const [quiz, setQuiz] = useState(quizId)
	const [create] = useCreateQuizMutation()
	const [update] = useUpdateQuizMutation()

	const {
		control,
		reset,
		handleSubmit,
		formState: { dirtyFields },
	} = useForm<IQuizForm>({ values: data?.data || defaultValues })

	const saveHandler = async (form: IQuizForm) => {
		const newData: IQuizDTO = {
			...form,
			id: data?.data.id,
			imageLink: !form.image || typeof form.image == 'object' ? '' : data?.data.image,
			image: typeof form.image == 'string' ? undefined : form.image,
		}

		// const formData = new FormData()
		// Object.keys(form).map(k => formData.append(k, form[k as 'title']))
		// formData.append('imageLink', data?.data.image || '')

		try {
			if (method == 'create') {
				const payload = await create(newData).unwrap()
				localStorage.setItem(localKeys.Quiz, JSON.stringify({ id: payload.id, time: Date.now() }))
				setQuiz(payload.id)
				toast.success(payload.message)
			} else {
				const payload = await update(newData).unwrap()
				reset(form)
				toast.success(payload.message)
			}
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<Stack>
			<Stack
				justifyContent={'center'}
				alignItems={'center'}
				position={'relative'}
				component={'form'}
				onSubmit={handleSubmit(saveHandler)}
			>
				{isFetching && <Fallback position={'absolute'} left={0} top={0} background={'#c6e5ff42'} zIndex={5} />}

				<Stack direction={'row'} spacing={1.5} width={'100%'}>
					<Controller
						name='image'
						control={control}
						render={({ field }) => (
							<UploadImage
								value={field.value}
								onChange={field.onChange}
								sx={{ flexBasis: '30%', maxHeight: 210 }}
							/>
							// <Stack spacing={1} flexBasis={'30%'} alignItems={'flex-start'}>
							// 	<UploadButton onChange={files => field.onChange(files[0])} />
							// 	<Preview file={field.value} onDelete={() => field.onChange()} height={200} />
							// </Stack>
						)}
					/>

					<Stack spacing={1.5} flexBasis={'70%'}>
						<Controller
							control={control}
							name='title'
							rules={{
								required: { value: true, message: 'Поле обязательно для заполнения' },
								minLength: { value: 3, message: 'Минимальная длина названия 3 символа' },
								maxLength: { value: 255, message: 'Максимальная длина названия 255 символов' },
							}}
							render={({ field, fieldState: { error } }) => (
								<TextField
									{...field}
									label='Название'
									fullWidth
									error={Boolean(error)}
									helperText={error?.message}
								/>
							)}
						/>

						<Controller
							control={control}
							name='description'
							render={({ field }) => (
								<TextField {...field} label='Описание' fullWidth multiline minRows={4} />
							)}
						/>

						<Controller
							control={control}
							name='hasShuffle'
							render={({ field }) => (
								<FormControlLabel
									label='Перемешивать вопросы'
									control={<Checkbox checked={field.value} {...field} sx={{ my: '-3px' }} />}
								/>
							)}
						/>
					</Stack>
				</Stack>

				{/* //TODO добавить кнопку опубликовать и оставшиеся поля */}
				<Stack direction={'row'} spacing={2} position={'absolute'} bottom={16}>
					<Button
						disabled={!Object.keys(dirtyFields).length}
						variant='outlined'
						type='submit'
						sx={{
							textTransform: 'inherit',
							boxShadow: 'inset 0 0 0px 20px white',
							':hover': { boxShadow: 'inset 0 0 0px 20px #F5F6FA' },
						}}
					>
						<SaveIcon
							fill={!Object.keys(dirtyFields).length ? palette.action.disabled : palette.primary.main}
							fontSize={18}
							mr={1}
						/>
						Сохранить
					</Button>

					{/* //TODO добавить кнопку для удаления теста */}
				</Stack>
				<Divider sx={{ width: '90%', mt: 3, mb: 4 }} />
			</Stack>

			<QuestionListForm quizId={quiz} />
		</Stack>
	)
}
