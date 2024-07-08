import { ChangeEvent, DragEvent, FC, useState } from 'react'
import { Box, Button, Stack, SxProps, Theme, Typography } from '@mui/material'

import { OpenExternalIcon } from '@/components/Icons/OpenExternalIcon'
import { TrashIcon } from '@/components/Icons/TrashIcon'
import { ImageIcon } from './ImageIcon'
import { Image } from './Image'
import Input from './Input'

type Props = {
	fullscreen?: boolean
	sx?: SxProps<Theme>
	value?: File | string
	onChange: (file?: File | null) => void
}

export const UploadImage: FC<Props> = props => {
	if (props.value) return <Preview {...props} />
	return <UploadInput {...props} />
}

const Preview: FC<Props> = ({ value, onChange, sx }) => {
	const deleteHandler = () => onChange(null)

	if (!value) return null
	return (
		<Box
			border={`1px solid #0000003b`}
			borderRadius={3}
			display={'flex'}
			flexDirection={'column'}
			justifyContent={'center'}
			alignItems={'center'}
			flexGrow={1}
			padding={2}
			position={'relative'}
			sx={{ ...sx }}
		>
			<Box
				position={'absolute'}
				width={'100%'}
				height={'100%'}
				borderRadius={3}
				display={'flex'}
				justifyContent={'center'}
				alignItems={'center'}
				sx={{
					background: '#00000063',
					opacity: 0,
					transition: 'all .3s ease-in-out',
					':hover': {
						opacity: '1',
					},
				}}
			>
				<Stack direction={'row'} spacing={2}>
					<Button
						href={typeof value == 'string' ? '/' + value : URL.createObjectURL(value)}
						target='_blank'
						rel='noopener noreferrer'
						sx={{ minWidth: 44, boxShadow: 'inset 0 0 0px 20px white' }}
					>
						<OpenExternalIcon fontSize={18} />
					</Button>
					<Button onClick={deleteHandler} sx={{ minWidth: 44, boxShadow: 'inset 0 0 0px 20px white' }}>
						<TrashIcon fontSize={18} />
					</Button>
				</Stack>
			</Box>
			<Image
				src={typeof value == 'string' ? '/' + value : URL.createObjectURL(value)}
				alt={typeof value == 'string' ? value : value.name}
			/>
		</Box>
	)
}

const UploadInput: FC<Props> = ({ onChange, sx }) => {
	const [hasDropFiles, setHasDropFiles] = useState(false)

	const changeHandler = (event: ChangeEvent<HTMLInputElement>) => {
		const files = event.target.files
		if (!files) return

		onChange(files[0])
	}

	const dragHandler = (event: DragEvent<HTMLDivElement>) => {
		event.preventDefault()
		event.stopPropagation()

		if (event.type === 'dragenter' || event.type === 'dragover') {
			setHasDropFiles(true)
		} else if (event.type === 'dragleave') {
			setHasDropFiles(false)
		}
	}

	const dropHandler = (event: DragEvent<HTMLDivElement>) => {
		event.preventDefault()
		event.stopPropagation()
		setHasDropFiles(false)

		const files = event.dataTransfer.files
		onChange(files[0])
	}

	return (
		<Box
			border={`1px solid #0000003b`}
			borderRadius={3}
			display={'flex'}
			flexDirection={'column'}
			justifyContent={'center'}
			alignItems={'center'}
			flexGrow={1}
			padding={2}
			boxShadow={hasDropFiles ? 'inset 0 0 20px #00000028' : undefined}
			sx={{
				cursor: 'pointer',
				transition: 'all 0.3s ease-in-out',
				':hover': {
					// boxShadow: 'inset 0 0 20px #00000028',
					borderColor: '#000000de',
				},
				...sx,
			}}
			component='label'
			onDragEnter={dragHandler}
			onDragLeave={dragHandler}
			onDragOver={dragHandler}
			onDrop={dropHandler}
		>
			<ImageIcon fontSize={100} fill={'#646464'} />
			<Typography fontSize={'1.1rem'} color={'#464646'}>
				Выберите или перетащите файл
			</Typography>
			<Input onChange={changeHandler} type='file' />
		</Box>
	)
}
