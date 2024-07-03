import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const NoImage: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 256 256' xmlSpace='preserve'>
				<g
					stroke='none'
					strokeWidth={0}
					strokeDasharray={'none'}
					strokeLinecap='butt'
					strokeLinejoin='miter'
					strokeMiterlimit={10}
					fill='none'
					fillRule='nonzero'
					opacity={1}
				>
					<path
						d='M89 20.938a1 1 0 0 0-1 1v46.125a7.631 7.631 0 0 1-2.898 5.983L62.328 50.71a1.001 1.001 0 0 0-1.372-.057L45.058 64.479l-2.862-2.942a1 1 0 1 0-1.434 1.394l3.521 3.62c.37.38.972.405 1.373.058l15.899-13.826 21.783 22.32a7.607 7.607 0 0 1-2.987.608H24.7a1 1 0 1 0 0 2h55.651c5.32 0 9.648-4.328 9.648-9.647V21.938a.998.998 0 0 0-.999-1zM89.744 4.864a1.002 1.002 0 0 0-1.412-.077l-8.363 7.502H9.648C4.328 12.29 0 16.618 0 21.938v46.125c0 4.528 3.141 8.328 7.356 9.361l-7.024 6.3a1 1 0 0 0 1.335 1.488l88-78.935a1 1 0 0 0 .077-1.413zM9.648 14.29h68.091L34.215 53.33 23.428 42.239a1 1 0 0 0-1.385-.046L2 60.201V21.938c0-4.217 3.431-7.648 7.648-7.648zM2 68.063v-5.172l20.665-18.568 10.061 10.345-23.44 21.024C5.238 75.501 2 72.157 2 68.063z'
						stroke='none'
						strokeWidth={1}
						strokeDasharray={'none'}
						strokeLinecap='butt'
						strokeLinejoin='miter'
						strokeMiterlimit={10}
						fill='#000'
						fillRule='nonzero'
						opacity={1}
						transform='matrix(2.81 0 0 2.81 1.407 1.407)'
					/>
					<path
						d='M32.607 35.608c-4.044 0-7.335-3.291-7.335-7.335s3.291-7.335 7.335-7.335 7.335 3.291 7.335 7.335-3.29 7.335-7.335 7.335zm0-12.67c-2.942 0-5.335 2.393-5.335 5.335s2.393 5.335 5.335 5.335 5.335-2.393 5.335-5.335-2.393-5.335-5.335-5.335z'
						stroke='none'
						strokeWidth={1}
						strokeDasharray={'none'}
						strokeLinecap='butt'
						strokeLinejoin='miter'
						strokeMiterlimit={10}
						fill='#000'
						fillRule='nonzero'
						opacity={1}
						transform='matrix(2.81 0 0 2.81 1.407 1.407)'
					/>
				</g>
			</svg>
		</SvgIcon>
	)
}
