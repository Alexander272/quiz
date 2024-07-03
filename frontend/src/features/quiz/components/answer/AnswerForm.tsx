import { FC } from 'react'
import { Button, Checkbox, Stack, TextField } from '@mui/material'
import { Controller, useFieldArray, useFormContext } from 'react-hook-form'

import type { IQuestionForm } from '../../types/question'
import { RoundPlusIcon } from '@/components/Icons/RoundPlusIcon'
import { PlusIcon } from '@/components/Icons/PlusIcon'

type Props = unknown

export const AnswerForm: FC<Props> = () => {
	const { control } = useFormContext<IQuestionForm>()
	const { fields } = useFieldArray({ control, name: 'answers' })

	return (
		<Stack spacing={1} mb={2}>
			{fields.map((f, i) => (
				<Stack key={f.id} direction={'row'} spacing={1} alignItems={'center'}>
					<Controller
						control={control}
						name={`answers.${i}.isCorrect`}
						render={({ field }) => <Checkbox checked={field.value} {...field} />}
					/>

					<Controller
						control={control}
						name='text'
						rules={{ required: { value: true, message: 'Поле обязательно для заполнения' } }}
						render={({ field, fieldState: { error } }) => (
							<TextField
								{...field}
								label='Ответ'
								fullWidth
								multiline
								minRows={2}
								error={Boolean(error)}
								helperText={error?.message}
							/>
						)}
					/>
				</Stack>
			))}

			<Button variant='outlined' sx={{ textTransform: 'inherit' }}>
				<RoundPlusIcon />
				<PlusIcon />
				Добавить вариант ответа
			</Button>
		</Stack>
	)
}
