var mysql = require("mysql");

// Database connecion configration
const connection = mysql.createConnection({
  host: "",
  user: "root",
  password: "secret",
  database: "practice02",
  port: 3306,
});

// Verify connection status.
connection.connect((error) => {
  if (error) {
    console.error("Error connecting to database:", error);
  } else {
    console.log("Connected to database!");
  }
});

module.exports= connection;
