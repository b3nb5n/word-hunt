import { Component } from 'solid-js'
import classes from './board.module.css'
import Tile from './tile'

export interface BoardProps {
	size?: number
}

const Board: Component<BoardProps> = ({ size = 4 }) => {
	const letters = Array(size)
		.fill(null)
		.map(() => new Array<string>(size).fill(''))
	const tiles = letters.flat().map((letter) => <Tile value={letter} />)

	return <div class={classes.board}>{tiles}</div>
}

export default Board
