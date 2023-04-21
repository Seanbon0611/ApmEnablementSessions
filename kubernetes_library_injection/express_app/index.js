const express = require("express");
const app = express();
const port = 3005;


app.get("/", (req, res) => {
res.status(200).send("Welcome To This Express App!");
})

app.listen(port, () => {
console.log(`Example app listening on port ${port}`);
});
