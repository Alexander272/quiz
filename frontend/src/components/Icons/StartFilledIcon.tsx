import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const StartFilledIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg
				xmlns='http://www.w3.org/2000/svg'
				viewBox='0 0 122.88 122.88'
				enableBackground={'new 0 0 122.88 122.88'}
				xmlSpace='preserve'
			>
				<path
					fillRule='evenodd'
					clipRule={'evenodd'}
					d='M61.44 0c33.93 0 61.44 27.51 61.44 61.44s-27.51 61.44-61.44 61.44S0 95.37 0 61.44 27.51 0 61.44 0zm23.47 65.52c3.41-2.2 3.41-4.66 0-6.61L49.63 38.63c-2.78-1.75-5.69-.72-5.61 2.92l.11 40.98c.24 3.94 2.49 5.02 5.8 3.19l34.98-20.2z'
				/>
			</svg>
		</SvgIcon>
	)
}
