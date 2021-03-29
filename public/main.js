let depth = 6;
let limit = 16;
let loading = false;

const button = document.querySelector('button');
button.onclick = async () => {
	if (loading) return;
	const tiles = Array.from(document.querySelectorAll('.tile'));
	const letters = tiles.map((tile) => `${tile.value}`);

	const endpoint =
		location.hostname == 'localhost' || location.hostname == '127.0.0.1'
			? 'http://localhost:8080'
			: `https://word-hunt.web.app/api`;

	loading = true;
	button.innerText = '...';

	try {
		const response = await fetch(endpoint, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ letters, depth, limit }),
		});

		const words = await response.json();
		const wordsList = document.querySelector('ol');

		while (wordsList.firstChild) {
			wordsList.removeChild(wordsList.firstChild);
		}

		words.forEach((word) => {
			const li = document.createElement('li');
			li.className = 'word';
			li.innerText = word;
			wordsList.appendChild(li);
		});
	} catch (e) {
		console.log(e);
	} finally {
		loading = false;
		button.innerText = 'Get Words';
	}
};

document.querySelector('#depth').oninput = (e) => {
	const { min, max } = e.target;
	const value = parseInt(e.target.value);
	const percentage = ((value - min) / (max - min)) * 100;

	depth = value;
	const label = document.querySelector('label[for=depth]');
	label.innerText = `Depth - ${value}`;
	e.target.style.background = `linear-gradient(to right, black 0%, black ${percentage}%, white ${percentage}%, white 100%)`;
};
