import { ChangeEvent, FC } from 'react'

import { Icon, Input, Label, Wrapper } from './style'

type Props = {
	id: string
	value?: boolean
	onChange?: (event: ChangeEvent<HTMLInputElement>) => void
	label?: string
}

export const Checkbox: FC<Props> = ({ id, label, value, onChange }) => {
	return (
		<Wrapper>
			<Input id={id} checked={value} onChange={onChange} type='checkbox' />
			<Label htmlFor={id}>
				<Icon>
					<svg width='18px' height='18px' viewBox='0 0 18 18'>
						<path d='M1,9 L1,3.5 C1,2 2,1 3.5,1 L14.5,1 C16,1 17,2 17,3.5 L17,14.5 C17,16 16,17 14.5,17 L3.5,17 C2,17 1,16 1,14.5 L1,9 Z' />
						<polyline points='1 9 7 14 15 4' />
					</svg>
				</Icon>

				{label && <span>{label}</span>}
			</Label>
		</Wrapper>
	)
}
