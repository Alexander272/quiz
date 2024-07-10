import { FC } from 'react'
import { Box, Grow, Stack, Typography } from '@mui/material'

import type { IQuestion } from '../../types/question'
import { Image } from '@/features/files/components/UploadImage/Image'
import { ShortDivider } from '@/components/Divider/ShortDivider'
import { Answer } from '../answer/Answer'

type Props = {
	data: IQuestion
}

export const Question: FC<Props> = ({ data }) => {
	return (
		<Grow key={data.id} in timeout={500}>
			<Box>
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
		</Grow>
	)
}

// const grow = keyframes`
//  	0% {
//     	opacity: 0;
//     	transform: scale(0);
//   	}
//   	100% {
//    		opacity: 1;
//     	transform: scale(1);
//   	}
// `

// export const Question: FC<Props> = ({ data }) => {
// 	return (
// 		<Box key={data.id} sx={{ animation: `${grow} 0.3s ease` }}>
// 			<Stack spacing={4} mb={1}>
// 				{data.image ? (
// 					<Box display={'flex'} maxWidth={300} maxHeight={200}>
// 						<Image src={data.image} alt={data.image} />
// 					</Box>
// 				) : null}

// 				<Typography fontSize={'1.1rem'} align='justify' paddingX={1}>
// 					{data.text}
// 				</Typography>
// 			</Stack>
// 			<ShortDivider sx={{ mb: 3 }} />

// 			<Stack spacing={1.4}>
// 				{data.answers.map(a => (
// 					<Answer key={a.id} data={a} />
// 				))}
// 			</Stack>
// 		</Box>
// 	)
// }
