import { FC } from 'react'
import { Box, Stack, Typography } from '@mui/material'

import type { IQuestion } from '../../types/question'
import { Image } from '@/features/files/components/UploadImage/Image'
import { ShortDivider } from '@/components/Divider/ShortDivider'
import { Answer } from '../answer/Answer'

type Props = {
	data: IQuestion
}

export const Question: FC<Props> = ({ data }) => {
	return (
		<Box
			padding={2}
			borderRadius={3}
			paddingX={3}
			paddingY={3}
			border={'1px solid rgba(0, 0, 0, 0.12)'}
			sx={{ backgroundColor: '#fff' }}
		>
			<Stack spacing={4} mb={1}>
				{data.image ? (
					<Box display={'flex'} maxWidth={300} maxHeight={200}>
						<Image src={data.image} alt={data.image} />
					</Box>
				) : null}

				<Typography fontSize={'1.1rem'} align='justify' paddingX={1}>
					{data.text}
				</Typography>
			</Stack>
			<ShortDivider sx={{ mb: 3 }} />

			<Stack spacing={1.4}>
				{data.answers.map(a => (
					<Answer key={a.id} data={a} />
				))}
			</Stack>
		</Box>
	)
}
