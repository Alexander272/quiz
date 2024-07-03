import { Checkbox, Divider, FormControlLabel, Stack, TextField } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'

import type { IQuizForm } from '../../types/quiz'
import { QuestionForm } from '../question/QuestionForm'
import { UploadImage } from '@/features/files/components/UploadImage/UploadImage'

const defaultValues: IQuizForm = {
	title: '',
	description: '',
	numberOfAttempts: 3,
	hasShuffle: true,
	hasSkippable: false,
	showList: false,
	showAnswers: false,
	showResults: false,
	time: '',
}

export const CreateQuiz = () => {
	const methods = useForm<IQuizForm>({ defaultValues: defaultValues })

	return (
		<Stack spacing={3} justifyContent={'center'} alignItems={'center'}>
			<Stack direction={'row'} spacing={1.5} width={'100%'}>
				<Controller
					name='image'
					control={methods.control}
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
						control={methods.control}
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
						control={methods.control}
						name='description'
						render={({ field }) => (
							<TextField {...field} label='Описание' fullWidth multiline minRows={4} />
						)}
					/>

					<Controller
						control={methods.control}
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

			{/* //TODO добавить кнопку для сохранения и оставшиеся поля */}

			{/* //TODO  а точно ли здесь мне нужно вставлять вопросы */}
			<Divider sx={{ width: '90%' }} />
			<QuestionForm />
		</Stack>
	)
}
