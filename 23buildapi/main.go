package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learning to build Api in Go ")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{Courseid: "2", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{Fullname: "Vaibhav", Website: "abcsd.com"}})
	courses = append(courses, Course{Courseid: "3", CourseName: "VueJs", CoursePrice: 199, Author: &Author{Fullname: "Vaibhav", Website: "vuejs.com"}})

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

// model for course - ideally should be in a seperate file
type Course struct {
	Courseid    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware / helper --file
func (c *Course) IsEmpty() bool {
	//return c.Courseid == "" && c.CourseName == ""
	return c.CourseName == ""

}

//controllers = file

// serveHome route. to send api
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API by Vaibhav "))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // getting data from database
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r) // take variables/parameters sent through request into params

	// loop through each courses, find matching id and return the response
	for _, course := range courses {
		if course.Courseid == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course.")
	w.Header().Set("Content-Type", "application/json")

	// what if request body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	//what if request body = {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data. Its empty")
		return
	}

	// what if same data/coursename is inserted again
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseName == params["coursename"] {
			json.NewEncoder(w).Encode("Data(coursename) alreasy exists.")
			return
		}
	}

	// generate unique id, string
	rand.Seed(time.Now().UnixNano())               //generate a random int number
	course.Courseid = strconv.Itoa(rand.Intn(100)) //convert it to string
	courses = append(courses, course)              // adding new course to DB
	json.NewEncoder(w).Encode(course)              //probably sending the new generated course
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course.")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop DB, find matching id, remove it, add new with id
	for index, course := range courses {
		if params["id"] == course.Courseid {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.Courseid = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting one course.")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.Courseid == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted.")
			break
		}
	}
}
