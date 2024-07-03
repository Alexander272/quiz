import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const VisibleIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 122.88 68.18'>
				<path d='M61.44 13.81a20.31 20.31 0 1 1-14.34 6 20.24 20.24 0 0 1 14.34-6zM1.05 31.31a106.72 106.72 0 0 1 10.32-10.88C25.74 7.35 42.08.36 59 0s34.09 5.92 50.35 19.32a121.91 121.91 0 0 1 12.54 12 4 4 0 0 1 .25 5 79.88 79.88 0 0 1-15.38 16.41 69.53 69.53 0 0 1-43.33 15.45 76 76 0 0 1-44.26-14.36A89.35 89.35 0 0 1 .86 36.44a3.94 3.94 0 0 1 .19-5.13zm15.63-5A99.4 99.4 0 0 0 9.09 34a80.86 80.86 0 0 0 14.62 13.37A68.26 68.26 0 0 0 63.4 60.3a61.69 61.69 0 0 0 38.41-13.72 70.84 70.84 0 0 0 12-12.3 110.45 110.45 0 0 0-9.5-8.86C89.56 13.26 74.08 7.58 59.11 7.89s-29.48 6.59-42.43 18.38zm39.69-7.79a7.87 7.87 0 1 1-7.87 7.87 7.86 7.86 0 0 1 7.87-7.87z' />
			</svg>
		</SvgIcon>
	)
}
