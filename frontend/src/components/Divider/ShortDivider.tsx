import { FC } from 'react'
import { Divider, SxProps, Theme } from '@mui/material'

type Props = {
	sx?: SxProps<Theme>
	flexItem?: boolean
}

export const ShortDivider: FC<Props> = ({ flexItem, sx }) => {
	return <Divider flexItem={flexItem} sx={{ width: '120px', borderColor: '#4285f4', borderBottomWidth: 2, ...sx }} />
}
