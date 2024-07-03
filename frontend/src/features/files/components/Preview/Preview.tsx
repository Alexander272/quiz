import { FC } from 'react'
import { Box, Button, useTheme } from '@mui/material'

import { NoImage } from './NoImage'
import { TimesIcon } from '@/components/Icons/TimesIcon'

type Props = {
	file?: File
	width?: number | string
	height?: number | string
	onDelete?: () => void
}

export const Preview: FC<Props> = ({ file, width, height, onDelete }) => {
	console.log(file)
	const { palette } = useTheme()

	return (
		<Box
			display={'flex'}
			justifyContent={'center'}
			alignItems={'center'}
			width={width || '100%'}
			maxHeight={height}
			position={'relative'}
		>
			{!file ? (
				<NoImage fontSize={90} height={110} />
			) : (
				<>
					<Button
						onClick={onDelete}
						variant='outlined'
						// color='error'
						color='gray'
						sx={{ position: 'absolute', right: 5, top: 5, minWidth: 44 }}
					>
						<TimesIcon fontSize={16} fill={palette.gray.main} />
					</Button>
					<img src={URL.createObjectURL(file)} alt={file.name} />
				</>
			)}
		</Box>
	)
}
