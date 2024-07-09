import { FC } from 'react'
import {
	Button,
	Checkbox,
	CircularProgress,
	Divider,
	FormControlLabel,
	Stack,
	TextField,
	Typography,
	useTheme,
} from '@mui/material'
import { Controller, FormProvider, useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IQuestion, IQuestionDTO, IQuestionForm } from '../../types/question'
import { UploadImage } from '@/features/files/components/UploadImage/UploadImage'
import { SaveIcon } from '@/components/Icons/SaveIcon'
import { TrashIcon } from '@/components/Icons/TrashIcon'
import { AnswerForm } from '../answer/AnswerForm'
import { useCreateQuestionMutation, useDeleteQuestionMutation, useUpdateQuestionMutation } from '../../questionApiSlice'
import { Fallback } from '@/components/Fallback/Fallback'
import { RefreshIcon } from '@/components/Icons/RefreshIcon'
import { useDeleteFileMutation, useUploadFilesMutation } from '@/features/files/filesApiSlice'
import { Confirm } from '@/components/Confirm/Confirm'
import { WarningIcon } from '@/components/Icons/WarningIcon'

const defaultValues: IQuestionForm = {
	number: 1,
	text: '',
	description: '',
	hasShuffle: true,
	level: '1',
	points: 1,
	time: 0,
	answers: [
		{ number: 1, text: '', isCorrect: false },
		{ number: 2, text: '', isCorrect: false },
	],
}

type Props = {
	quizId: string
	method: 'create' | 'update'
	position?: number
	data?: IQuestion
}

export const QuestionForm: FC<Props> = ({ quizId, method, position, data }) => {
	const { palette } = useTheme()

	const [create, { isLoading }] = useCreateQuestionMutation()
	const [update, { isLoading: isLoadUpdate }] = useUpdateQuestionMutation()
	const [remove, { isLoading: isLoadRemove }] = useDeleteQuestionMutation()

	const [upload] = useUploadFilesMutation()
	const [removeImage] = useDeleteFileMutation()

	//TODO поправить данные
	const methods = useForm<IQuestionForm>({ values: data || { ...defaultValues, number: position || 1 } })
	const {
		control,
		handleSubmit,
		reset,
		formState: { dirtyFields },
	} = methods

	const resetHandler = () => reset()

	const saveHandler = async (form: IQuestionForm) => {
		const newData: IQuestionDTO = {
			quizId: quizId,
			...form,
			number: data?.number || position || 1,
			image: '',
			// imageLink: !form.image || typeof form.image == 'object' ? '' : data?.image,
			// image: typeof form.image == 'string' ? undefined : form.image,
		}
		console.log('data', data)
		console.log('question', form)

		try {
			if (!form.image || typeof form.image == 'object') {
				if (data?.image) {
					removeImage(data.image)
					// удалить картинку если она была
				}
				// сохранить картинку и записать полученный путь в newData
				if (form.image) {
					const fileData = new FormData()
					console.log(form.image)
					fileData.append('path', `${quizId}/${data?.id || 'temp'}/${(form.image as File).name}`)
					fileData.append('image', form.image as File)
					const payload = await upload({ data: fileData }).unwrap()
					newData.image = payload.data
				}
			}

			if (method == 'create') {
				await create(newData).unwrap()
				reset(defaultValues)
			} else {
				await update(newData).unwrap()
				reset(newData)
			}
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	const deleteHandler = async () => {
		if (!data?.id) return

		try {
			await remove({ id: data.id, quizId: data.quizId })
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<Stack
			// spacing={1}
			border={'1px solid #E0E0E0'}
			borderRadius={3}
			paddingY={1.5}
			paddingX={2}
			width={'100%'}
			// maxWidth={'98%'}
			position={'relative'}
			component={'form'}
			onSubmit={handleSubmit(saveHandler)}
		>
			{isLoading || isLoadUpdate ? (
				<Fallback position={'absolute'} left={0} top={0} background={'#c6e5ff42'} zIndex={5} />
			) : null}

			<FormProvider {...methods}>
				<Stack
					direction={'row'}
					justifyContent={'space-between'}
					alignItems={'center'}
					paddingX={2}
					paddingY={1.5}
					// borderRadius={2}
					mb={2}
					borderBottom={'1px solid #E0E0E0'}
					// sx={{ background: '#f8fafc' }}
				>
					<Stack direction={'row'}>
						<Typography fontSize={'1.2rem'}>Вопрос №{data?.number || position}</Typography>
					</Stack>

					<Stack spacing={1} direction={'row'}>
						<Button
							onClick={resetHandler}
							disabled={!Object.keys(dirtyFields).length}
							variant='outlined'
							color='gray'
							sx={{ minWidth: 44 }}
						>
							<RefreshIcon
								fontSize={18}
								fill={!Object.keys(dirtyFields).length ? palette.action.disabled : palette.gray.main}
							/>
						</Button>

						<Button
							disabled={!Object.keys(dirtyFields).length}
							variant='outlined'
							type='submit'
							sx={{ minWidth: 44 }}
						>
							<SaveIcon
								fontSize={18}
								fill={!Object.keys(dirtyFields).length ? palette.action.disabled : palette.primary.main}
							/>
						</Button>

						<Confirm
							onClick={deleteHandler}
							disabled={!data}
							buttonComponent={
								<Button disabled={!data} variant='outlined' color='error' sx={{ minWidth: 44 }}>
									{isLoadRemove ? (
										<CircularProgress size={16} color='error' />
									) : (
										<TrashIcon
											fontSize={18}
											fill={!data ? palette.action.disabled : palette.error.main}
										/>
									)}
								</Button>
							}
						>
							<Stack direction={'row'} spacing={1} alignItems={'center'} justifyContent={'center'} mb={1}>
								<WarningIcon fill={palette.error.main} />
								<Typography fontSize={'1.1rem'} fontWeight='bold'>
									Удаление
								</Typography>
							</Stack>

							<Typography maxWidth={300} align='center'>
								Вы уверены, что хотите удалить вопрос?
							</Typography>
						</Confirm>
					</Stack>
				</Stack>

				<Stack spacing={1.5} direction={'row'}>
					<Controller
						name='image'
						control={methods.control}
						render={({ field }) => (
							<UploadImage
								value={field.value}
								onChange={field.onChange}
								sx={{ flexBasis: '30%', maxHeight: 210 }}
							/>
						)}
					/>

					<Stack spacing={1.5} flexBasis={'70%'}>
						<Controller
							control={control}
							name='text'
							rules={{
								required: { value: true, message: 'Поле обязательно для заполнения' },
								minLength: { value: 3, message: 'Минимальная длина вопроса 3 символа' },
							}}
							render={({ field, fieldState: { error } }) => (
								<TextField
									{...field}
									label='Текст вопроса'
									fullWidth
									multiline
									minRows={3}
									error={Boolean(error)}
									helperText={error?.message}
								/>
							)}
						/>

						<Controller
							control={control}
							name='description'
							render={({ field }) => (
								<TextField {...field} label='Комментарий' fullWidth multiline minRows={2} />
							)}
						/>

						<Controller
							control={control}
							name='hasShuffle'
							render={({ field }) => (
								<FormControlLabel
									label='Перемешивать варианты ответов'
									control={<Checkbox checked={field.value} {...field} sx={{ my: '-3px' }} />}
								/>
							)}
						/>
					</Stack>
				</Stack>

				{/* //TODO куда-то еще надо ответы запихивать и кнопку для их добавления */}
				<Divider sx={{ width: '86%', marginX: 'auto', mt: 2, mb: 3 }} />
				<AnswerForm />
			</FormProvider>
		</Stack>
	)
}
