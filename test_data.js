import { restClient } from '@polygon.io/client-js';
import 'dotenv/config';

const rest = restClient(process.env.POLY_KEY);

rest.stocks.aggregates("AAPL", 1, "day", "2023-01-01", "2023-04-14").then((data) => {
	console.log(data);
}).catch(e => {
	console.error('An error happened:', e);
});