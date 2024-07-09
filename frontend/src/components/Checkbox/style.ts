import styled from '@emotion/styled'

export const Icon = styled.span`
	height: 18px;
	width: 18px;
	position: relative;

	svg {
		position: relative;
		z-index: 1;
		fill: none;
		stroke-linecap: round;
		stroke-linejoin: round;
		stroke: #c8ccd4;
		stroke-width: 1.5;
		transform: translate3d(0, 0, 0);
		transition: all 0.2s ease;

		path {
			stroke-dasharray: 60;
			stroke-dashoffset: 0;
		}

		polyline {
			stroke-dasharray: 22;
			stroke-dashoffset: 66;
		}
	}

	&::before {
		content: '';
		position: absolute;
		top: -10px;
		left: -10px;
		width: 38px;
		height: 38px;
		border-radius: 50%;
		background: rgba(34, 50, 84, 0.03);
		opacity: 0;
		transition: opacity 0.2s ease;
	}

	&:hover {
		&::before {
			opacity: 1;
		}

		svg {
			stroke: #4285f4;
		}
	}
`

export const Label = styled.label`
	/* display: inline-block; */
	cursor: pointer;
	position: relative;
	margin: auto;
	/* width: 18px; */
	/* height: 18px; */
	-webkit-tap-highlight-color: transparent;
	transform: translate3d(0, 0, 0);

	display: inline-flex;
	align-items: center;
	vertical-align: middle;
	gap: 14px;
	padding: 6px;
`

export const Input = styled.input`
	display: none;
	visibility: hidden;
`

export const Wrapper = styled.div`
	/* height: 18px; */

	${Input}:checked + ${Label} ${Icon} svg {
		stroke: #4285f4;

		path {
			stroke-dashoffset: 60;
			transition: all 0.3s linear;
		}

		polyline {
			stroke-dashoffset: 42;
			transition: all 0.2s linear;
			transition-delay: 0.15s;
		}
	}
`
