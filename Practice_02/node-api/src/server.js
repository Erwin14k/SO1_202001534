const express = require('express');
const processesQuery = require('./processesQuery');

// Express app Configuration
const app = express();
app.use(express.json());
app.use(cors());


// Principal route of the server
app.get('/', (req, res) => {
    res.send('Hi from backend, we are working hard!!');
});

// Get cp & ram data route
app.get('/cpu-ram', async (req, res) => {
    try {
        await processesQuery();
        res.send('');
    } catch (err) {
        console.error(err);
        res.status(500).send('Error Loading Temporal table :(');
    }
});

// Get processes route
app.get('/processes', async (req, res) => {
    try {
        await processesQuery();
        res.send('Temporal Table Loaded Successfully!!');
    } catch (err) {
        console.error(err);
        res.status(500).send('Error Loading Temporal table :(');
    }
});

// Initialize the server
app.listen(3000, () => {
    console.log(`Server running on port: ${3000}`);
});