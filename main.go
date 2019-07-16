package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type employee struct {
	firstName string
	lastName  string
}

type project struct {
	Name      string
	Client    string
	StartDate string
	Status    string
}

var employees = []employee{
	{
		firstName: "Christian",
		lastName:  "Campos",
	},
	{
		firstName: "Keith",
		lastName:  "Grant",
	},
	{
		firstName: "James",
		lastName:  "Hellar",
	},
	{
		firstName: "Nima",
		lastName:  "Jalali",
	},
	{
		firstName: "Shawyan",
		lastName:  "Rahbar",
	},
	{
		firstName: "Claire",
		lastName:  "Rhoda",
	},
}

func setupRoutes() {

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		getEmployee(w, r)
	})

	http.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		createProject(w, r)
	})

	// Ending slash indicates a subtree
	http.HandleFunc("/project/", func(w http.ResponseWriter, r *http.Request) {
		deleteProject(w, r)
	})
}

func getEmployee(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Fetching employee")

	firstName := r.URL.Query().Get("firstName")

	fmt.Printf("this is the first name %+v\n", firstName)
	if firstName == "" {
		return errors.New("empty query string parameter")
	}

	// pseudo-query of finding employee with given name
	for _, employee := range employees {
		if firstName == employee.firstName {
			fmt.Fprintf(w, "we found the employee, %+v\n", employee.lastName)

			w.Write([]byte(employee.lastName))

			return nil
		}
	}

	fmt.Fprintf(w, "unable to find requested employee with name: %+v\n", firstName)

	w.WriteHeader(404)

	return nil
}

func createProject(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("creating project")

	//extract body and put into project struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var project project

	err = json.Unmarshal(body, &project)
	if err != nil {
		return err
	}

	if project.StartDate == "" {
		return errors.New("empty start date")
	}

	const layout = "2006-01-02"
	startDate, _ := time.Parse(layout, project.StartDate)

	t := time.Now()
	formattedTime := t.Format(layout)

	currentDate, err := time.Parse(layout, formattedTime)
	if err != nil {
		fmt.Println("error parsing currentDate ", currentDate)
		return err
	}

	if startDate.Before(currentDate) {
		project.Status = "infeasible"
		w.WriteHeader(400)
		fmt.Fprintln(w, "start date before current date")
		return err
	}

	project.Status = "pending"

	responseBytes, err := json.Marshal(&project)
	if err != nil {
		fmt.Printf("there was an error marshalling json %+v\n", responseBytes)
		return err
	}

	w.WriteHeader(200)
	w.Write(responseBytes)

	return nil
}

func deleteProject(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Deleting project!")

	idString := strings.Replace(r.URL.Path, "/project/", "", 1)

	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Fprintln(w, "error converting id string to integer")
		return err
	}

	err = deletePage(id)
	if err != nil {

		if err == errNotFound {
			w.WriteHeader(404)
			return err
		}

		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %+v\n", err)

		return err
	}

	fmt.Fprintln(w, "page deleted")
	w.WriteHeader(204)

	return nil

}

var errNotFound = errors.New("not found")

// Random int determines error type
func deletePage(pageID int) error {
	rand.Seed(int64(pageID))
	deleteStatus := rand.Intn(3)

	switch deleteStatus {
	case 0:
		return nil
	case 1:
		return errNotFound
	default:
		return errors.New("internal error")
	}
}

func main() {
	fmt.Println("HTTP Server Practice App")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
