import { FC, MouseEvent, PropsWithChildren, useRef, useState } from 'react'
import { Box, Button, Popover, Stack } from '@mui/material'

type Props = {
	onClick: () => void
	fullWidth?: boolean
	disabled?: boolean
	buttonComponent?: JSX.Element
}

export const Confirm: FC<PropsWithChildren<Props>> = ({ children, fullWidth, disabled, onClick, buttonComponent }) => {
	const [open, setOpen] = useState(false)
	const anchor = useRef<HTMLButtonElement>(null)

	const toggleHandler = (event: MouseEvent) => {
		if (disabled) return
		event.stopPropagation()
		setOpen(prev => !prev)
	}

	const confirmHandler = (event: MouseEvent) => {
		toggleHandler(event)
		onClick()
	}

	return (
		<Box ref={anchor} onClick={toggleHandler} sx={{ width: fullWidth ? '100%' : 'inherit' }}>
			{buttonComponent ? (
				buttonComponent
			) : (
				<Button disabled={disabled} variant='contained' color='error' fullWidth>
					Удалить
				</Button>
			)}

			<Popover
				open={open}
				anchorEl={anchor.current}
				onClose={toggleHandler}
				anchorOrigin={{
					vertical: 'center',
					horizontal: 'center',
				}}
				transformOrigin={{
					vertical: 'top',
					horizontal: 'center',
				}}
				// sx={{ mt: 1 }}
			>
				<Stack spacing={2} paddingX={2} paddingY={1.2}>
					<Box>{children}</Box>

					<Stack direction='row' spacing={2}>
						<Button onClick={confirmHandler} variant='contained' color='error' fullWidth>
							Да
						</Button>
						<Button onClick={toggleHandler} variant='outlined' fullWidth>
							Отмена
						</Button>
					</Stack>
				</Stack>
			</Popover>
		</Box>
	)
}
