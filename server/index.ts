import express, { Express, Request, Response } from 'express';
import dotenv from 'dotenv';
dotenv.config();

const app: Express = express();
const port = process.env.PORT || 3000;

app.get('/', (req: Request, res: Response) => {
	res.send("Hello from Express using TypeScript");
});

app.listen(port, (err) => {
	if(err) console.error("Error in running server:" err);

	console.log("[server]: Server is running on port ", PORT);
});