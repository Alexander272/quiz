import { ChangeEvent, FC } from 'react'
import { Typography } from '@mui/material'

import Input from './styled/Input'
import Button from './styled/Button'
import { UploadIcon } from './UploadIcon'

type Props = {
	// onChange: (event: ChangeEvent<HTMLInputElement>) => void
	onChange: (files: FileList) => void
	multiple?: boolean
}

export const UploadButton: FC<Props> = ({ onChange, multiple }) => {
	const changeHandler = (event: ChangeEvent<HTMLInputElement>) => {
		const files = event.target.files
		if (!files) return

		onChange(files)
	}

	return (
		<Button component='label'>
			<UploadIcon />
			<Typography ml={1}>Загрузить {'файл' + (multiple ? 'ы' : '')}</Typography>

			<Input onChange={changeHandler} type='file' multiple={multiple} />
		</Button>
	)
}
