import { config } from 'dotenv';
import { createConnection } from 'mysql';
config();


// Database connecion configration
const connection = createConnection({
    host: process.env.HOST,
    user: process.env.USER,
    password: process.env.PASSWORD,
    database: process.env.DATABASE,
    port: process.env.PORT,
});

// Verify connection status.
connection.connect((error) => {
    if (error) {
        console.error('Error connecting to database:', error);
    } else {
    console.log('Connected to database!');
    }
});

export default connection;
