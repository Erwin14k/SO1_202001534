const connection=require("./config/dbConfig")


async function processesQuery() {
    try {
    
    } catch (err) {
        console.error('Error Processes Query:', err);
    } finally {
        if (conn) {
        try {
            await conn.close();
        }catch (err) {
            console.error(err);
        }
        }
    }
}

module.exports = processesQuery;