package main

// Imports
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

// Main function
func main() {

	/*	// Operate endpoint
		http.HandleFunc("/operate", handleOperate)
		// Port 8080 listening
		fmt.Println("Listening on port 8080...")
		http.ListenAndServe(":8080", nil)
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/operate", handleOperate)

	// Inicia el servidor en el puerto 8080 con CORS habilitado
	handler := cors.Default().Handler(mux)
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println(err)
	}
}

func handleOperate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Expression captured
	//expression := r.FormValue("expression")
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
	result := 0.0
	// Switch to evaluate the operator
	switch operator {
	// Sum
	case "+":
		result = num1 + num2
	// Subtraction
	case "-":
		result = num1 - num2
	// Multiplication
	case "*":
		result = num1 * num2
	// Division
	case "/":
		// Invalid Operation
		if num2 == 0 {
			result = -1499
			// Valid Operation
		} else {
			result = num1 / num2
		}

	}
	fmt.Fprint(w, result)
}
