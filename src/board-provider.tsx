import React, { useState } from 'react'

export type Board = string[][]
export const BOARD_SIZE = 4

export interface BoardState {
	tiles: string[][]
	setTile(row: number, col: number, value: string): void
}

const defaultBoardState: BoardState = {
	// Initialize 2D array of strings
	tiles: Array<string[]>(BOARD_SIZE)
		.fill([])
		.map(() => new Array<string>(BOARD_SIZE).fill('')),
	setTile: () => {
		// If the context is used outside of the provider throw an error
		throw Error('Unimplemented')
	},
}

export const boardContext = React.createContext(defaultBoardState)

const BoardProvider: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
	const [tiles, setTiles] = useState<string[][]>(defaultBoardState.tiles)

	const setTile = (row: number, col: number, value: string) => {
		const newTiles = tiles.map((row) => [...row])
		newTiles[row][col] = value.charAt(0)
		setTiles(newTiles)
	}

	return <boardContext.Provider value={{ tiles, setTile }}>{children}</boardContext.Provider>
}

export default BoardProvider
