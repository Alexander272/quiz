import { Box, Typography } from '@mui/material'

import { PageBox } from '@/styled/PageBox'
import { CreateQuiz } from '@/features/quiz/components/quiz/CreateQuiz'

export default function Create() {
	return (
		<PageBox>
			<Box
				borderRadius={3}
				paddingX={3}
				paddingY={3}
				width={'60%'}
				marginX={'auto'}
				border={'1px solid rgba(0, 0, 0, 0.12)'}
				// flexGrow={1}
				height={'fit-content'}
				minHeight={200}
				display={'flex'}
				flexDirection={'column'}
				sx={{ backgroundColor: '#fff', userSelect: 'none' }}
			>
				<Typography fontSize={'1.3rem'} textAlign={'center'} mt={1} mb={3}>
					Создать тест
				</Typography>
				<CreateQuiz />
			</Box>
		</PageBox>
	)
}
