const { default: axios } = require('axios');
const http = require('http');
require('dotenv').config();

const server = http.createServer((req, res) => {
  res.write("Hello from node!");
  res.end();
});

const goUrl = `http://127.0.0.1:${process.env.goPort}`;

server.listen(process.env.nodePort, "127.0.0.1", () => {
  console.log(`Node server listening on http://127.0.0.1:${process.env.nodePort}`);
});

setInterval(async () => {
  try {
    console.log('Sent request to go.')
    const response = await axios.get(goUrl)
    console.log(`Recieved from go: ${response.data}`)
  } catch (error) {
    console.log(`Request to go server (${goUrl}) failed.`)
  }
}, 2000)