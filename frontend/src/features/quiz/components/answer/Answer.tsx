import { FC } from 'react'
import { Stack, Typography } from '@mui/material'
import { Controller, useFormContext } from 'react-hook-form'

import type { IAnswer } from '../../types/answer'
import type { IUserQuizForm } from '../../types/quiz'
import { Checkbox } from '@/components/Checkbox/Checkbox'

type Props = {
	data: IAnswer
}

export const Answer: FC<Props> = ({ data }) => {
	const { control } = useFormContext<IUserQuizForm>()

	return (
		<Controller
			control={control}
			name={`${data.questionId}.answers.${data.id}`}
			render={({ field }) => (
				<Stack
					onClick={() => field.onChange(!field.value)}
					direction={'row'}
					spacing={2}
					alignItems={'center'}
					paddingY={1}
					paddingX={2}
					border={`1px solid ${field.value ? '#4285f4' : '#E0E0E0'}`}
					boxShadow={field.value ? 'inset 0 0 8px #4285f44f' : undefined}
					borderRadius={2}
					sx={{ transition: 'all 0.3s ease-in-out', cursor: 'pointer' }}
				>
					<Checkbox id={data.id} value={field.value || false} onChange={field.onChange} />
					<Typography>{data.text}</Typography>
				</Stack>
			)}
		/>
	)
}
