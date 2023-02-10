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
	fmt.Println()
	fmt.Println()

	defer db.Close()

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
	// Multiplication
	case "*":
		result = num1 * num2
	// Division
	case "/":
		// If num2 is 0, is an invalid operation
		if num2 == 0 {
			result = -1499

		} else {
			result = num1 / num2
		}

	}
	fmt.Fprint(w, result)
}

/*func handleGetLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("si entramos")
	rows, err := db.Query("SELECT * FROM logs")
	if err != nil {
		http.Error(w, "Error in get_logs Query  :(", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []LogStruct
	for rows.Next() {
		var log LogStruct
		if err := rows.Scan(&log.ID, &log.LeftOperand, &log.RightOperand, &log.Operator, &log.Result); err != nil {
			http.Error(w, "Error while scanning logs table records :(", http.StatusInternalServerError)
			return
		}
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error while processing logs table records :(", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		http.Error(w, "Error encoding records from logs table to JSON", http.StatusInternalServerError)
		return
	}
}*/

/*func handleGetLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM logs")
	if err != nil {
		fmt.Println("errro111")
		http.Error(w, "Error in get_logs Query  :(", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []LogStruct
	for rows.Next() {
		var log LogStruct
		var date string
		if err := rows.Scan(&log.ID, &log.LeftOperand, &log.RightOperand, &log.Operator, &log.Result, &date); err != nil {
			fmt.Println("errro222")
			http.Error(w, "Error while scanning logs table records :(", http.StatusInternalServerError)
			return
		}
		log.Date = date
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("errro333")
		http.Error(w, "Error while processing logs table records :(", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		fmt.Println("errro444")
		http.Error(w, "Error encoding records from logs table to JSON", http.StatusInternalServerError)
		return
	}
}*/

func handleGetLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT * FROM logs")
	if err != nil {
		http.Error(w, "Error in get_logs Query  :(", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []LogStruct
	for rows.Next() {
		var log LogStruct
		var date string
		if err := rows.Scan(&log.ID, &log.RightOperand, &log.Operator, &log.LeftOperand, &log.Result, &date); err != nil {
			http.Error(w, "Error while scanning logs table records :(", http.StatusInternalServerError)
			return
		}
		log.Date = date
		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error while processing logs table records :(", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		http.Error(w, "Error encoding records from logs table to JSON", http.StatusInternalServerError)
		return
	}
}
