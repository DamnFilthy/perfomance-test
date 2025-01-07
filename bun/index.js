import express from "express";

const fs = require("fs");

const app = express();
const port = 3007;

app.get("/", (req, res) => {
  res.send("Hello from Bun!");
});

app.get("/bun-file", (req, res) => {
    const start = performance.now();
    
    const filePath = 'result.json';
    const result = [];

    for (let i = 0; i <= 10_000_000; i++) {
        result.push({
            id: i,
            name: 'User' + i,
            content: '<h1>Hello, World!</h1>'
        })
    }

    fs.writeFileSync(filePath, JSON.stringify(result));
    const data = fs.readFileSync(filePath, 'utf8');
    const fileResult = JSON.parse(data);

    const fileResultFilter = fileResult.filter(item => item.id % 2 !== 0);
    const middleItem = fileResultFilter[Math.round(fileResult.length / 2)];

    res.send({
        platform: "Bun v1",
        time: performance.now() - start,
        item: middleItem
    });

    const duration = performance.now() - start;

    console.log("Total duration:", duration)
});

app.get("/bun-test", (req, res) => {
    const start = performance.now();
    
    const result = [];

    for (let i = 0; i <= 10_000_000; i++) {
        result.push({
            id: i,
            name: 'User' + i,
            content: '<h1>Hello, World!</h1>'
        })
    }

    const filterResult = result.filter(item => item.id % 2 !== 0);
    const middleItem = filterResult[Math.round(filterResult.length / 2)];

    res.send({
        platform: "Bun v1",
        time: performance.now() - start,
        item: middleItem
    });

    const duration = performance.now() - start;

    console.log("Total duration:", duration)
});

app.listen(port, () => {
  console.log(`Listening on port ${3007}...`);
});