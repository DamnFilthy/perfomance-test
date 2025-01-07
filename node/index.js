const express = require("express")
const fs = require("fs");
const { platform } = require("os");

const app = express();
const port = 3005;

app.get("/", (req, res) => {
  res.send("Hello from Node!");
});

app.get("/node-file", (req, res) => {
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
        platform: "Node v20",
        time: performance.now() - start,
        item: middleItem
    });

    const duration = performance.now() - start;

    console.log("Total duration:", duration)
});

app.get("/node-test", (req, res) => {
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
        platform: "Node v20",
        time: performance.now() - start,
        item: middleItem
    });

    const duration = performance.now() - start;

    console.log("Total duration:", duration)
});

app.listen(port, () => {
  console.log(`Listening on port ${port}...`);
});