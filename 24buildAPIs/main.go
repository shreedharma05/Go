package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// Middlewares or helper
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LCO")
	r := mux.NewRouter()
	// seeding values
	courses = append(courses, Course{CourseID: "1", CourseName: "NestJS", CoursePrice: 299, Author: &Author{FullName: "Tom", Website: "go.tom.dev"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "ReactJS", CoursePrice: 249, Author: &Author{FullName: "Jerry", Website: "go.jerry.dev"}})
	courses = append(courses, Course{CourseID: "3", CourseName: "KoaJS", CoursePrice: 249, Author: &Author{FullName: "Tom", Website: "go.tom.dev"}})
	courses = append(courses, Course{CourseID: "4", CourseName: "NextJS", CoursePrice: 299, Author: &Author{FullName: "Jerry", Website: "go.jerry.dev"}})

	// Routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses/add", addOneCourse).Methods("POST")
	r.HandleFunc("/course/update/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/delete/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to port
	http.ListenAndServe(":8000", r)

}

// controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" <h1>Welcome to Home Page of learncodeonline "))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // this sets the header to json response so that browser is aware of the media type (recieved)
	json.NewEncoder(w).Encode(courses)                 // NewEncoder creates an instance and the instance can convert data to json, json to byte code and writes it back to ResponseWriter
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// grab id using mux package or url
	params := mux.Vars(r)
	// loop through and find course that match
	fmt.Println("The params are: ", params)
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course available with given id")
	return
}

func addOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide some data")
	}
	// what if sent data is {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please provide a course name")
	}

	// Duplicate titles

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already available")
			return
		}
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop, id, remove, add updated course
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseID = params["id"]
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			break
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course successfully destroyed")
		}
	}
}
