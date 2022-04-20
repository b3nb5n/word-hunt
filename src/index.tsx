import React from 'react'
import ReactDOM from 'react-dom/client'
import Board from './board'
import BoardProvider from './board-provider'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>
		<BoardProvider>
			<Board />
		</BoardProvider>
	</React.StrictMode>
)
