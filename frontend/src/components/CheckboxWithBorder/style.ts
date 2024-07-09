import { keyframes } from '@emotion/react'
import styled from '@emotion/styled'

const bottomCheck = keyframes`
    0% {
        height: 0;
    }
    100% {
        height: calc(var(--checkbox-height) / 2);
    }
`
const topCheck = keyframes`
    0% {
        height: 0;
    }
    50% {
        height: 0;
    }
    100% {
        height: calc(var(--checkbox-height) * 1.2);
    }
`

export const Input = styled.input`
	display: none;
`

export const Label = styled.label`
	height: var(--checkbox-height);
	width: var(--checkbox-height);
	background-color: transparent;
	border: calc(var(--checkbox-height) * 0.1) solid var(--color);
	border-radius: 5px;
	position: relative;
	display: inline-block;
	box-sizing: border-box;
	transition: border-color ease 0.2s;
	cursor: pointer;

	&:hover {
		border-color: var(--hover-color);
	}

	&::before,
	&::after {
		content: '';
		box-sizing: border-box;
		position: absolute;
		height: 0;
		width: calc(var(--checkbox-height) * 0.2);
		background-color: var(--active-color);
		display: inline-block;
		transform-origin: left top;
		border-radius: 5px;
		transition: opacity ease 0.5;
	}

	&::before {
		top: calc(var(--checkbox-height) * 0.72);
		left: calc(var(--checkbox-height) * 0.41);
		box-shadow: 0 0 0 calc(var(--checkbox-height) * 0.05) var(--background-color);
		transform: rotate(-135deg);
	}
	&::after {
		top: calc(var(--checkbox-height) * 0.37);
		left: calc(var(--checkbox-height) * 0.05);
		transform: rotate(-45deg);
	}
`

type Props = {
	color?: string
	activeColor?: string
	hoverColor?: string
}

export const Wrapper = styled.div<Props>`
	box-sizing: border-box;
	--background-color: #fff;
	--checkbox-height: 25px;
	--color: ${props => props.color || '#000'};
	--active-color: ${props => props.activeColor || '#34b93d'};
	--hover-color: ${props => props.hoverColor || '#000'};
	height: var(--checkbox-height);

	${Input}:checked + ${Label} {
		border-color: var(--active-color);
	}

	${Input}:checked + ${Label}::after {
		height: calc(var(--checkbox-height) / 2);
		animation: ${bottomCheck} 0.2s ease 0s forwards;
	}
	${Input}:checked + ${Label}::before {
		height: calc(var(--checkbox-height) * 1.2);
		animation: ${topCheck} 0.4s ease 0s forwards;
	}
`
