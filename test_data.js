import { restClient } from '@polygon.io/client-js';
import 'dotenv/config';

const rest = restClient(process.env.POLY_KEY);

// for aggregates
rest.stocks.aggregates("AAPL", 1, "day", "2023-01-01", "2023-04-14").then((data) => {
	console.log(data.results);
	const high_cos_map = data.results.map((high) => high.h);
	console.log(high_cos_map);

	let sum = 0.0;
	let avg = 0.0;
	let counter = 0;

	high_cos_map.forEach((cost) => {
		sum += cost;
		counter++;
		avg = sum / counter;
	});

	console.log(`The Average high cost for Apple stock price is: ${avg}`);
}).catch(e => {
	console.error('An error happened:', e);
});