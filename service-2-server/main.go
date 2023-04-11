package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", getEmployeeJson)
	log.Println("API endpoint running at -> http://localhost:4317/get")
	http.ListenAndServe(":4317", nil)
}

type Employee struct {
	Id         string `json:"id"`
	First      string `json:"first"`
	Last       string `json:"last"`
	Department string `json:"department ,omitempty"`
}

var employees = map[string]Employee{
	"0000000000": Employee{Id: "101", First: "Susan", Last: "Matthew", Department: "HR"},
	"0000000001": Employee{Id: "102", First: "Bill", Last: "Gates", Department: "Finance"},
	"0000000002": Employee{Id: "103", First: "Prateek", Last: "Singh", Department: "Engineering"},
	"0000000003": Employee{Id: "104", First: "Rakesh", Last: "Singh", Department: "IT"},
}

func getEmployees() []Employee {
	values := make([]Employee, len(employees))
	i := 0
	for _, emp := range employees {
		values[i] = emp
		i++
	}
	return values
}

func getEmployeeJson(w http.ResponseWriter, r *http.Request) {
	emps := getEmployees()
	data, err := json.Marshal(emps)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
