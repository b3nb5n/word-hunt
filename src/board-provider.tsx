import React, { PropsWithChildren, useEffect, useState } from 'react'
import wasmInit from '../assets/word_hunt.wasm'

type Solutions = Record<string, string>
interface Board {
	tiles: string[][]
	setTile(row: number, col: number, value: string): void
	solve(): Promise<Solutions>
}

const BOARD_SIZE = 4
const defaultHandler = () => {
	throw Error('Unimplemented')
}

const defaultBoard: Board = {
	tiles: Array<string[]>(BOARD_SIZE).map(() =>
		Array<string>(BOARD_SIZE).fill('')
	),
	setTile: defaultHandler,
	solve: defaultHandler,
}

export const boardContext = React.createContext(defaultBoard)
const BoardProvider: React.FC<PropsWithChildren<{}>> = ({
	children,
}) => {
	const [tiles, setTiles] = useState(defaultBoard.tiles)
	const [solutions, setSolutions] = useState<Solutions>()

	useEffect(() => {
		wasmInit({}).then(console.log)
	}, [])

	const setTile = (row: number, col: number, value: string) => {
		if (row < 0 || row > BOARD_SIZE || col < 0 || col > BOARD_SIZE)
			return
		const newTiles = tiles.map((row) => [...row])
		newTiles[row][col] = value.charAt(0).toLowerCase()
		setTiles(newTiles)
	}

	const solve = async () => {
		if (!solutions) setSolutions({})
		return {}
	}

	return (
		<boardContext.Provider value={{ tiles, setTile, solve }}>
			{children}
		</boardContext.Provider>
	)
}

export default BoardProvider
