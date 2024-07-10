import { FC } from 'react'
import { IQuestion } from '../../types/question'
import { Stack, Typography } from '@mui/material'
import { ShortDivider } from '@/components/Divider/ShortDivider'
import { useFormContext } from 'react-hook-form'
import { IUserQuizForm } from '../../types/quiz'

type Props = {
	data: IQuestion[]
	active: number
	onSelect: (index: number) => void
	height?: string | number
}

export const QuestionList: FC<Props> = ({ data, active, height, onSelect }) => {
	const { watch } = useFormContext<IUserQuizForm>()
	const form = watch()

	return (
		<Stack
			height={'100%'}
			maxHeight={height || 500}
			padding={2}
			borderRadius={3}
			paddingX={3}
			paddingY={3}
			border={'1px solid rgba(0, 0, 0, 0.12)'}
			sx={{ backgroundColor: '#fff' }}
		>
			<Typography>Список вопросов</Typography>
			<ShortDivider sx={{ width: 70, mb: 3, mt: 1 }} />

			<Stack spacing={1}>
				{/* //TODO надо бы наверное виртуальный список делать */}
				{data.map((d, i) => {
					let complete = false
					if (form[d.id]?.answers) {
						complete = Object.values(form[d.id]?.answers)?.includes(true)
					}

					return (
						<Stack
							key={d.id}
							onClick={() => onSelect(i)}
							paddingY={1}
							paddingX={2}
							border={`1px solid ${active == i ? '#4285f4' : '#E0E0E0'}`}
							borderRadius={2}
							sx={{
								cursor: 'pointer',
								backgroundColor: complete ? '#efedff3b' : undefined,
								transition: 'all 0.3s ease-in-out',
								':hover': { borderColor: '#585858' },
							}}
						>
							{d.text}
						</Stack>
					)
				})}
			</Stack>
		</Stack>
	)
}
