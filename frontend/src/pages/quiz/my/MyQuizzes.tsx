import { Box } from '@mui/material'

import { MyQuizList } from '@/features/quiz/components/quiz/MyQuizList'
import { PageBox } from '@/styled/PageBox'

export default function MyQuizzes() {
	return (
		<PageBox>
			<Box
				borderRadius={3}
				paddingX={3}
				paddingY={3}
				width={'100%'}
				maxWidth={'1200px'}
				marginX={'auto'}
				border={'1px solid rgba(0, 0, 0, 0.12)'}
				// flexGrow={1}
				height={'fit-content'}
				minHeight={600}
				maxHeight={800}
				display={'flex'}
				flexDirection={'column'}
				sx={{ backgroundColor: '#fff', userSelect: 'none' }}
			>
				<MyQuizList />
			</Box>
		</PageBox>
	)
}
