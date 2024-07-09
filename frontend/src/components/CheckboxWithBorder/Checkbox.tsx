import { ChangeEvent, FC } from 'react'

import { Input, Label, Wrapper } from './style'

type Props = {
	id: string
	value?: boolean
	onChange?: (event: ChangeEvent<HTMLInputElement>) => void
	color?: string
}

export const Checkbox: FC<Props> = ({ id, value, onChange }) => {
	return (
		<Wrapper color='#00000099'>
			<Input checked={value} onChange={onChange} type='checkbox' id={id} />
			<Label htmlFor={id} />
		</Wrapper>
	)
}
