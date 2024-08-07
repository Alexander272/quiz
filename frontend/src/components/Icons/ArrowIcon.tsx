import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const ArrowIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg
				xmlns='http://www.w3.org/2000/svg'
				viewBox='0 0 66.91 122.88'
				enableBackground={'new 0 0 66.91 122.88'}
				xmlSpace='preserve'
			>
				<path d='M1.95 111.2a6.875 6.875 0 0 0 .14 9.73 6.875 6.875 0 0 0 9.73-.14L64.94 66l-4.93-4.79 4.95 4.8c2.65-2.74 2.59-7.11-.15-9.76-.08-.08-.16-.15-.24-.22L11.81 2.09c-2.65-2.73-7-2.79-9.73-.14-2.72 2.65-2.78 7-.13 9.73l48.46 49.55L1.95 111.2z' />
			</svg>
		</SvgIcon>
	)
}
