import { FC } from 'react'
import { Button, Checkbox, Stack, TextField, useTheme } from '@mui/material'
import { Controller, useFieldArray, useFormContext } from 'react-hook-form'

import type { IQuestionForm } from '../../types/question'
import { PlusIcon } from '@/components/Icons/PlusIcon'
import { TimesIcon } from '@/components/Icons/TimesIcon'

type Props = unknown

export const AnswerForm: FC<Props> = () => {
	const { palette } = useTheme()

	const { control } = useFormContext<IQuestionForm>()
	const { fields, append, remove } = useFieldArray({ control, name: 'answers' })

	const addNewHandler = () => {
		append({ number: fields.length + 1, text: '', isCorrect: false })
	}

	const removeHandler = (index: number) => {
		remove(index)
	}

	return (
		<Stack spacing={2} mb={1}>
			{fields.map((f, i) => (
				<Stack key={f.id} direction={'row'} spacing={1} alignItems={'flex-start'}>
					<Controller
						control={control}
						name={`answers.${i}.isCorrect`}
						render={({ field }) => <Checkbox checked={field.value} {...field} />}
					/>

					<Controller
						control={control}
						name={`answers.${i}.text`}
						rules={{ required: { value: true, message: 'Поле обязательно для заполнения' } }}
						render={({ field, fieldState: { error } }) => (
							<TextField
								{...field}
								label='Ответ'
								fullWidth
								multiline
								// minRows={2}
								error={Boolean(error)}
								helperText={error?.message}
							/>
						)}
					/>

					<Button
						onClick={() => removeHandler(i)}
						variant='outlined'
						color='gray'
						sx={{ minWidth: 20, padding: 1.4 }}
					>
						<TimesIcon fontSize={12} />
					</Button>
				</Stack>
			))}

			<Button
				onClick={addNewHandler}
				variant='outlined'
				sx={{ width: 300, textTransform: 'inherit', mx: 'auto!important' }}
			>
				{/* <RoundPlusIcon fill={palette.primary.main} fontSize={20} /> */}
				<PlusIcon fill={palette.primary.main} fontSize={14} mr={1} />
				Добавить вариант ответа
			</Button>
		</Stack>
	)
}
