import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const FinishIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 116.333 122.88' xmlSpace='preserve'>
				<path
					fillRule='evenodd'
					clipRule='evenodd'
					d='M20.416 17.743c46.635-32.238 50.118 31.566 95.917-8.271v65.01c-43.681 39.279-53.104-24.185-95.917 8.068V17.743zM8.898 0a8.9 8.9 0 0 1 8.899 8.898 8.896 8.896 0 0 1-4.375 7.663h.42V106.321h-.42c2.617 1.549 4.375 3.574 4.375 7.662 0 4.087-3.984 8.896-8.899 8.896-4.914 0-8.898-4.81-8.898-8.896 0-4.088 1.757-6.113 4.374-7.662h-.419V16.561h.419A8.9 8.9 0 0 1 8.898 0z'
				/>
			</svg>
		</SvgIcon>
	)
}
