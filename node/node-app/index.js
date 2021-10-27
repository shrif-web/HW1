
const express = require('express')
const { Client } = require('pg')
const app = express()
const port = 3000

app.get('/', async(req, res) => {
  const client = new Client({
    user: 'postgres',
    database: 'postgres',
    password: '.',
    ssl: false
  })  
  await client.connect()
  const _dbRes = await client.query('select name from student where id=1')
  await client.end()
  res.send({name: _dbRes.rows[0].name})
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
