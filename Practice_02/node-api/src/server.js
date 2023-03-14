const express = require("express");
const connection = require('./config/dbConfig');
const cors = require("cors");

// Express app Configuration
const app = express();
app.use(express.json());
app.use(cors());
// function called query using arrow function syntax. 
const query = (sql, args) => {
  /*This creates a new Promise object, which represents an asynchronous operation 
  that may or may not succeed.*/
  return new Promise((resolve, reject) => {
    connection.query(sql, args, (err, rows) => {
      if (err) return reject(err);
      resolve(rows);
    });
  });
};

// Principal route of the server
app.get("/", (req, res) => {
  res.send("Hi from backend, we are working hard!!");
});

// Get cpu & ram data route
app.get("/cpu-ram", async (req, res) => {
  try {
    // Execute the query
    const result = await query(
      `SELECT * FROM resource r ORDER BY r.resource DESC LIMIT 2;`
    );
    // Define the important data
    const dataToSend = result.length > 1 ? result[1] : result[0];
    // Send the query result
    res.send(dataToSend);
  } catch (error) {
    // In case of errors
    console.error(error);
    res.status(500).send("Internal Server Error");
  }
});


app.get("/get-processes", async (req, res) => {
  /*Constant variable result and assigns to it the result of an SQL query executed 
  by calling the query function*/
  const result = await query(`WITH resource_id AS (SELECT (MAX(resource)) AS resource FROM resource)SELECT p.resource, p.pid, p.name, p.user, p.status, p.ram_percentage, p.parent_process FROM process p, resource_id WHERE p.resource = resource_id.resource;`);
  /*Constant variable processCounts that is computed by calling the reduce function on
    the result array. The reduce function accumulates a count of the number of processes
    for each status, by using the status property of each process object.*/
  const processCounts = result.reduce((counts, process) => {
    counts[process.status] = (counts[process.status] || 0) + 1;
    return counts;
  }, {});
  /*Constant variable called "processes" that is computed by calling the map function 
  on the result array. The map function creates a new array of objects by copying all 
  properties of each process object in result except for the parent_process property. 
  It then adds a new property called parent_process with the value of the original 
  parent_process.
  */ 
  const processes = result.map(({ parent_process, ...rest }) => ({ ...rest, parent_process }));
  // Send the data
  res.send({
    ...processCounts,
    totalProcesses: result.length,
    processes,
  });
});


// Initialize the server
app.listen(8080, () => {
  console.log(`Server running on port: ${8080}`);
});
