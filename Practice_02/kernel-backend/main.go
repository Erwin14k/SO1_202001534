package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
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
	TotalRam int `json:"totalram"`
	FreeRam  int `json:"freeram"`
}

// Mysql DB Connection
var conn = MySQLConn()

func MySQLConn() *sql.DB {
	connString := "root:secret@tcp(:3306)/practice02"
	conn, error1 := sql.Open("mysql", connString)
	if error1 != nil {
		fmt.Println(error1)
	} else {
		fmt.Println("Connection MySQL")
	}
	return conn
}

func main() {
	var out [3][]byte
	var error1 [5]error
	var output [3]string

	for {
		cmd := exec.Command("sh", "-c", "cat /proc/stat | grep cpu | tail -1 | awk '{print ($5*100)/($2+$3+$4+$5+$6+$7+$8+$9+$10)}' | awk '{print 100-$1}'")
		out[0], error1[0] = cmd.CombinedOutput()

		if error1[0] != nil {
			fmt.Println(error1[0])
		}
		output[0] = string(out[0][:len(out[0])-1])

		cmdram := exec.Command("sh", "-c", "cat /proc/ram_202001534")
		out[1], error1[1] = cmdram.CombinedOutput()
		if error1[1] != nil {
			fmt.Println(error1[1])
		}
		output[1] = string(out[1][:])

		cmdcpu := exec.Command("sh", "-c", "cat /proc/cpu_202001534")
		out[2], error1[2] = cmdcpu.CombinedOutput()
		if error1[2] != nil {
			fmt.Println(error1[2])
		}
		output[2] = string(out[2][:])

		jsonstring := fmt.Sprintf("{\"cpu\":%s,\"ram\":%s,\"procs\":%s}", output[0], output[1], output[2])

		var temporalData Resource
		unmarshallData := json.Unmarshal([]byte(jsonstring), &temporalData)
		if unmarshallData != nil {
			fmt.Println(unmarshallData)
		}

		query := `INSERT INTO resource(date_resource,cpu_data,ram_data) VALUES (NOW(),?,?);`
		result, er := conn.Exec(query, temporalData.Cpu, float64(temporalData.Ram.TotalRam-temporalData.Ram.FreeRam)*100/float64(temporalData.Ram.TotalRam))
		if er != nil {
			fmt.Println(er)
		}

		resourceId, _ := result.LastInsertId()

		var parents []*Parent

		for i := 0; i < len(temporalData.Procs); i++ {
			temporalProc := temporalData.Procs[i]
			if len(temporalProc.Children) > 0 {
				var parent Parent
				parent.Value = &temporalProc
				for j := 0; j < len(temporalProc.Children); j++ {
					c := temporalProc.Children[j]
					for k := 0; k < len(temporalData.Procs); k++ {
						pr := temporalData.Procs[k]
						if pr.Pid == c {
							parent.Children = append(parent.Children, &pr)
						}
					}
				}
				parents = append(parents, &parent)
			}
		}
		for i := 0; i < len(parents); i++ {
			temporalParent := *parents[i].Value
			query := `INSERT INTO process (pid,name,user,status,ram_percentage,resource) VALUES (?,?,?,?,?,?);`
			state := ""
			switch temporalParent.Status {
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
			res, error2 := conn.Exec(query, temporalParent.Pid, temporalParent.Name, temporalParent.User, state, float64(temporalParent.Ram)/float64(temporalData.Ram.TotalRam), resourceId)
			if error2 != nil {
				fmt.Println(error2)
			}
			parentId, _ := res.LastInsertId()
			for j := 0; j < len(parents[i].Children); j++ {
				query2 := `INSERT INTO process(pid, name, user, status, ram_percentage, parent_process, resource) VALUES (?,?,?,?,?,?,?);`
				ch := *parents[i].Children[j]
				chState := ""
				switch ch.Status {
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
				default:
					chState = "Suspended"
				}
				_, er3 := conn.Exec(query2, ch.Pid, ch.Name, ch.User, chState, float64(ch.Ram)/float64(temporalData.Ram.TotalRam), parentId, resourceId)
				if er3 != nil {
					fmt.Println(er3)
				}
			}
		}
		fmt.Println(resourceId)
		time.Sleep(1 * time.Second)
	}
}
