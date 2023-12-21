require("dotenv").config();
const express= require("express");
const app = express();

app.get("/", (req, res) => {
	res.send("Connected to express server created by Kris!");
});


app.listen(process.env.PORT, () => {
	console.log("Connected to server");
});