import { Component } from 'solid-js'
import classes from './tile.module.css'

export interface TileProps {
	value: string
}

const Tile: Component<TileProps> = ({ value }) => {
	return <input type='text' class={classes.tile} value={value} />
}

export default Tile
