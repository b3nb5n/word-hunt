import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import init, { greet } from './wasm/pkg/wasm'

init().then(() => greet('world'))

ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>
		<App />
	</React.StrictMode>
)
