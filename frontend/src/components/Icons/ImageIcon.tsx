import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const ImageIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 122.88 91.27'>
				<path
					fillRule='evenodd'
					d='M87.29 16.88a11.35 11.35 0 1 1-11.35 11.35 11.35 11.35 0 0 1 11.35-11.35zm27.33 74.39H8.26a8.27 8.27 0 0 1-5.83-2.44A8.24 8.24 0 0 1 0 83V8.26a8.26 8.26 0 0 1 2.42-5.84A8.26 8.26 0 0 1 8.26 0h106.36a8.26 8.26 0 0 1 5.83 2.43 8.26 8.26 0 0 1 2.42 5.84V83a8.24 8.24 0 0 1-2.42 5.83 8.27 8.27 0 0 1-5.83 2.44zm-7.35-9.43L87.6 50.46a4.52 4.52 0 0 0-7.65 0l-9.29 14.93 10.11 16.45h-4.15l-27.57-44c-2.54-3.39-6.61-3.13-8.88 0l-27 44H9.42V9.42h104v72.42z'
				/>
			</svg>
		</SvgIcon>
	)
}
