import React, { useContext } from 'react'
import { boardContext } from './board-provider'
import styles from './board.module.css'

const BOARD_SIZE = 4

const Board: React.FC = () => {
	const { tiles, setTile } = useContext(boardContext)

	return (
		<div className={styles.board}>
			{tiles.flat().map((letter, i) => {
				const row = Math.floor(i / BOARD_SIZE)
				const col = i % BOARD_SIZE

				return (
					<input
						key={i}
						type='text'
						className={styles.tile}
						value={letter[0] ?? ''}
						onChange={(e) => setTile(row, col, e.target.value)}
					/>
				)
			})}
		</div>
	)
}

export default Board
