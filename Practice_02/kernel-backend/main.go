package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Process Struct
type Process struct {
	Pid      int    `json:"pid",omitempty`
	Name     string `json:"name",omitempty`
	User     int    `json:"user",omitempty`
	Status   int    `json:"status",omitempty`
	Ram      int    `json:"ram",omitempty`
	Children []int  `json:"children",omitempty`
}

// Resource Struct
type Resource struct {
	Cpu   float32   `json:"cpu",omitempty`
	Ram   Ram       `json:"ram",omitempty`
	Procs []Process `json:"procs",omitempty`
}

// Parent Process Struct
type Parent struct {
	Value    *Process
	Children []*Process
}

// Ram Struct
type Ram struct {
	TotalRam    int `json:"total_ram"`
	FreeRam     int `json:"free_ram"`
	OccupiedRam int `json:"ram_occupied"`
}

// Mysql DB Connection
var connection = MySQLConn()

// Connection function
func MySQLConn() *sql.DB {
	connString := "root:secret@tcp(34.133.134.153:3306)/practice02"
	connection, error1 := sql.Open("mysql", connString)
	if error1 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println("Connection MySQL")
	}
	if err := connection.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	return connection
}

func main() {
	var out [3][]byte
	var error1 [5]error
	var output [3]string

	for {
		// Cat to the ram module
		cmdram := exec.Command("sh", "-c", "cat /proc/ram_202001534")
		out[1], error1[1] = cmdram.CombinedOutput()
		if error1[1] != nil {
			fmt.Println(error1[1])
		}
		// Save the ram module information on the output array
		output[1] = string(out[1][:])

		// Cat to the cpu module
		cmdcpu := exec.Command("sh", "-c", "cat /proc/cpu_202001534")
		out[2], error1[2] = cmdcpu.CombinedOutput()
		if error1[2] != nil {
			fmt.Println(error1[2])
		}
		output[2] = string(out[2][:])
		// Split the cpu module information
		parts := strings.Split(output[2], "],")
		// Get the cpu value
    cpuValue := parts[0]
		cpuValue=strings.TrimPrefix(cpuValue, string(cpuValue[0]))
		// Save the cpu value on the output array
		output[0]=cpuValue
		output[2]=parts[1]
		// Json output
		jsonstring := fmt.Sprintf("{\"cpu\":%s,\"ram\":%s,\"procs\":%s}", output[0], output[1], output[2])

		var temporalData Resource
		unmarshallData := json.Unmarshal([]byte(jsonstring), &temporalData)
		if unmarshallData != nil {
			fmt.Println(unmarshallData)
		}
		// Start a transaction on the connection.
		tx, err := connection.Begin()
		if err != nil {
			fmt.Println(err)
			return
		}
		// If an error occurs, a Rollback (undo changes) is performed before the function ends.
		defer tx.Rollback()
		/* The "query" variable is the SQL query for the insert. The values ​​of "temporalData.Cpu"
		and "temporalData.Ram.TotalRam-temporalData.Ram.FreeRam)*100/float64(temporalData.Ram.TotalRam)"
		are inserted into the "cpu_data" and "ram_data" fields respectively.
		*/
		query := `INSERT INTO resource(date_resource,cpu_data,ram_data) VALUES (NOW(),?,?);`
		result, err := tx.Exec(query, temporalData.Cpu, float64(temporalData.Ram.OccupiedRam)/float64(temporalData.Ram.TotalRam)*100)
		if err != nil {
			fmt.Println(err)
			return
		}
		/* Gets the ID of the newly inserted row in the "resourceId" variable. If there are
		any errors during the execution of the query, the error is printed to the console.
		*/
		resourceId, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Insert processes and their children in the table 'process'
		var stmt *sql.Stmt
		query = `INSERT INTO process(pid, name, user, status, ram_percentage, parent_process, resource) VALUES (?, ?, ?, ?, ?, ?, ?);`
		// Prepare the query to avoid SQL injections.
		stmt, err = tx.Prepare(query)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Close the declaration when the function ends.
		defer stmt.Close()
		/* The processes obtained in 'temporalData' are traversed to insert them into the
		'process' table*/
		for i := 0; i < len(temporalData.Procs); i++ {
			temporalProc := temporalData.Procs[i]
			if len(temporalProc.Children) > 0 {
				// Status control
				state := ""
				switch temporalProc.Status {
				case 0:
					state = "Running"
				case 1:
					state = "Suspended"
				case 2:
					state = "Suspended"
				case 4:
					state = "Stopped"
				case 32:
					state = "Zombie"
				default:
					state = "Suspended"
				}
				res, err := stmt.Exec(temporalProc.Pid, temporalProc.Name, temporalProc.User, state, float64(temporalProc.Ram)/float64(temporalData.Ram.TotalRam), nil, resourceId)
				if err != nil {
					fmt.Println(err)
					return
				}
				// Get the ID of the inserted parent process
				parentId, err := res.LastInsertId()
				if err != nil {
					fmt.Println(err)
					return
				}
				// Insert children processes
				for j := 0; j < len(temporalProc.Children); j++ {
					ch := temporalProc.Children[j]
					for k := 0; k < len(temporalData.Procs); k++ {
						pr := temporalData.Procs[k]
						if pr.Pid == ch {
							chState := ""
							switch pr.Status {
							case 0:
								chState = "Running"
							case 1:
								chState = "Suspended"
							case 2:
								chState = "Suspended"
							case 4:
								chState = "Stopped"
							case 8:
								chState = "Zombie"
							case 32:
								chState = "Zombie"
							default:
								chState = "Suspended"
							}
							_, err = stmt.Exec(pr.Pid, pr.Name, pr.User, chState, float64(pr.Ram)/float64(temporalData.Ram.TotalRam), parentId, resourceId)
							if err != nil {
								fmt.Println(err)
								return
							}
						}
					}
				}
			}
		}
		// Confirm all the operations
		err = tx.Commit()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Round of processes finished
		fmt.Println("Finished Round")
		time.Sleep(5 * time.Second)
	}
}
