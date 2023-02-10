package main

// Imports
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	// CORS
	"log"

	"github.com/rs/cors"
)

// Struct to store the log data
type LogStruct struct {
	ID           int     `json:"id"`
	RightOperand int     `json:"right_operand"`
	Operator     string  `json:"operator"`
	LeftOperand  int     `json:"left_operand"`
	Result       float64 `json:"result"`
	Date         string  `json:"date_created"`
}

// Main function
func main() {

	// DB connection
	db, dbErr := sql.Open("mysql", "root:@{}Ee[]#$#$#$14kFer@tcp(127.0.0.1:3306)/Calculator")
	// Verify the DB Connection
	if dbErr != nil {
		log.Fatalf("Error Connection To the Database: %v", dbErr)
	}
	// Connection Established Message
	fmt.Println("Connection Established To The DataBase........")
	// MySql Version
	var version string
	// Mysql Version Query
	versionError := db.QueryRow("SELECT VERSION()").Scan(&version)
	// Verify Query
	if versionError != nil {
		log.Fatal(versionError)
	}
	// Mysql Version Message
	fmt.Println("Mysql_Version: " + version)
	defer db.Close()

	// Server Configuration
	mux := http.NewServeMux()
	mux.HandleFunc("/operate", func(w http.ResponseWriter, r *http.Request) {
		handleOperate(w, r, db)
	})
	mux.HandleFunc("/get-logs", func(w http.ResponseWriter, r *http.Request) {
		handleGetLogs(w, r, db)
	})
	// Server initialization on port 8080 with CORS enabled.
	handler := cors.Default().Handler(mux)
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", handler)
	// In case an error occurs
	if err != nil {
		fmt.Println(err)
	}
}

func handleOperate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Expression captured
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	expression := data["expression"]
	//fmt.Println(expression)
	if expression == "" {
		http.Error(w, "Missing expression", http.StatusBadRequest)
		return
	}
	// The two numbers variables
	var num1, num2 float64
	// Operator variable
	var operator string

	var i int
	// For loop to obtain the operator position
	for i = range expression {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/' {
			if i != 0 {
				break
			}

		}
	}
	// The left operand
	num1, _ = strconv.ParseFloat(expression[:i], 64)
	// The operator
	operator = expression[i : i+1]
	// The right operand
	num2, _ = strconv.ParseFloat(expression[i+1:], 64)

	// Float var to store the operation result
	var result float64 = 0.0
	// Switch to evaluate the operator
	switch operator {
	// Sum
	case "+":
		result = num1 + num2
		query := fmt.Sprintf("INSERT INTO logs (right_operand, left_operand, operator, result) VALUES (%d, %d, '%s', %f)", int(num1), int(num2), "+", result)
		_, err3 := db.Exec(query)
		if err3 != nil {
			log.Fatalf("Error, Cannot insert the log: %v", err)
		}
		fmt.Println("Log Inserted Successfully :)")
	// Subtraction
	case "-":
		result = num1 - num2
		query := fmt.Sprintf("INSERT INTO logs (right_operand, left_operand, operator, result) VALUES (%d, %d, '%s', %f)", int(num1), int(num2), "-", result)
		_, err3 := db.Exec(query)
		if err3 != nil {
			log.Fatalf("Error, Cannot insert the log: %v", err)
		}
		fmt.Println("Log Inserted Successfully :)")
	// Multiplication
	case "*":
		result = num1 * num2
		query := fmt.Sprintf("INSERT INTO logs (right_operand, left_operand, operator, result) VALUES (%d, %d, '%s', %f)", int(num1), int(num2), "*", result)
		_, err3 := db.Exec(query)
		if err3 != nil {
			log.Fatalf("Error, Cannot insert the log: %v", err)
		}
		fmt.Println("Log Inserted Successfully :)")
	// Division
	case "/":
		// If num2 is 0, is an invalid operation
		if num2 == 0 {
			result = -1499
			query := fmt.Sprintf("INSERT INTO logs (right_operand, left_operand, operator, result) VALUES (%d, %d, '%s', %f)", int(num1), int(num2), "/", result)
			_, err3 := db.Exec(query)
			if err3 != nil {
				log.Fatalf("Error, Cannot insert the log: %v", err)
			}
			fmt.Println("Log Inserted Successfully :)")
		} else {
			result = num1 / num2
			query := fmt.Sprintf("INSERT INTO logs (right_operand, left_operand, operator, result) VALUES (%d, %d, '%s', %f)", int(num1), int(num2), "/", result)
			_, err3 := db.Exec(query)
			if err3 != nil {
				log.Fatalf("Error, Cannot insert the log: %v", err)
			}
			fmt.Println("Log Inserted Successfully :)")
		}

	}
	fmt.Fprint(w, result)
}

func handleGetLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Perform a query to the "logs" table using the db.Query method and store the results in the rows variable.
	rows, err := db.Query("SELECT * FROM logs")
	// Check for errors executing the query
	if err != nil {
		http.Error(w, "Error in get_logs Query  :(", http.StatusInternalServerError)
		return
	}
	// We use 'defer rows.Close()' to make sure the connection to the database is closed.
	defer rows.Close()

	//Initialize a slice to hold the logs
	var logs []LogStruct
	// Iterate over each row returned by the query
	for rows.Next() {
		// Create a LogStruct variable to hold the values from each row
		var log LogStruct
		var date string
		// Scan the values from each row into the LogStruct variable
		if err := rows.Scan(&log.ID, &log.RightOperand, &log.Operator, &log.LeftOperand, &log.Result, &date); err != nil {
			// Return a 500 Internal Server Error response with a message indicating the scan error
			http.Error(w, "Error while scanning logs table records :(", http.StatusInternalServerError)
			return
		}
		log.Date = date
		// Append the LogStruct to the logs slice
		logs = append(logs, log)
	}

	// Check for errors after iterating over all the rows
	if err := rows.Err(); err != nil {
		// Return a 500 Internal Server Error response with a message indicating the processing error
		http.Error(w, "Error while processing logs table records :(", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to "application/json"
	w.Header().Set("Content-Type", "application/json")
	// Encode the logs slice as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		// Return a 500 Internal Server Error response with a message indicating the encoding error
		http.Error(w, "Error encoding records from logs table to JSON", http.StatusInternalServerError)
		return
	}
}
