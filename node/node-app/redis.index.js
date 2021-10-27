
const express = require('express')
const app = express()
const port = 3001
const redis = require("redis");
app.get('/', async(req, res) => {
  const client = redis.createClient();
  client.on("error", function (error){
    console.error(error);
  });
  client.set("key","value");
  client.get("key",(err, reply) => {
    console.log(reply);
    res.send({name : reply});
  });
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
