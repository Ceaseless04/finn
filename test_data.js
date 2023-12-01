import { restClient } from '@polygon.io/client-js';
import express from 'express';
import 'dotenv/config';

const app = express();
const rest = restClient(process.env.POLY_KEY);

// for aggregates
app.get('/', (req, res) => {
	rest.stocks.aggregates("AAPL", 1, "day", "2023-01-01", "2023-04-14").then((data) => {
		console.log(data.results);
		const high_cos_map = data.results.map((high) => high.h);
		

		let sum = 0.0;
		let avg = 0.0;
		let counter = 0;

		high_cos_map.forEach((cost) => {
			sum += cost;
			counter++;
			avg = sum / counter;
		});
		res.write(`The Average high cost for Apple stock price is: ${avg}`);
		res.json(high_cos_map);
		
	}).catch(e => {
		console.error('An error happened:', e);
	});
});

app.listen(process.env.PORT, () => {
	console.log("Server is up and running");
});