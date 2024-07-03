import { FC } from 'react'
import { Button, Checkbox, Divider, FormControlLabel, Stack, TextField, useTheme } from '@mui/material'
import { Controller, FormProvider, useForm } from 'react-hook-form'

import type { IQuestionForm } from '../../types/question'
import { UploadImage } from '@/features/files/components/UploadImage/UploadImage'
import { SaveIcon } from '@/components/Icons/SaveIcon'
import { TrashIcon } from '@/components/Icons/TrashIcon'
import { AnswerForm } from '../answer/AnswerForm'

const defaultValues: IQuestionForm = {
	number: 0,
	text: '',
	description: '',
	hasShuffle: true,
	level: '1',
	points: 1,
	time: '',
	answers: [
		{ number: 1, text: '', isCorrect: false },
		{ number: 2, text: '', isCorrect: false },
	],
}

type Props = unknown

export const QuestionForm: FC<Props> = () => {
	// const data = { data: { id: 'test' } }
	const { palette } = useTheme()

	const methods = useForm<IQuestionForm>({ defaultValues: defaultValues })
	const {
		control,
		formState: { isDirty },
	} = methods

	// if (!data?.data.id) return null
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
		>
			<FormProvider {...methods}>
				<Stack></Stack>

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
									label='Вопрос'
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

				<Stack spacing={1} direction={'row'} position={'absolute'} bottom={-14} right={30}>
					<Button
						disabled={!isDirty}
						variant='outlined'
						sx={{
							minWidth: 44,
							boxShadow: 'inset 0 0 0px 20px white',
							':hover': { boxShadow: 'inset 0 0 0px 20px #F5F6FA' },
						}}
					>
						<SaveIcon fontSize={18} fill={!isDirty ? palette.action.disabled : palette.primary.main} />
					</Button>
					<Button
						variant='outlined'
						color='error'
						sx={{
							minWidth: 44,
							boxShadow: 'inset 0 0 0px 20px white',
							':hover': { boxShadow: 'inset 0 0 0 20px #fdf7f7' },
						}}
					>
						<TrashIcon fontSize={18} fill={palette.error.main} />
					</Button>
				</Stack>
			</FormProvider>
		</Stack>
	)
}
