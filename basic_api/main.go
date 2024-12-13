package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Define model for course and author

type Author struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Course struct {
	ID         string  `json:"id"`
	CourseName string  `json:"course_name"`
	Price      float32 `json:"price"`
	Author     *Author `json:"author"`
}

func (c *Course) ReturnCourseWithId() Course {
	uid := uuid.New()
	c.ID = uid.String()
	return *c
}

// Simulating db with Course type slice
var courses []Course

// Define controllers

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// If no courses exist, ensure courses is an empty slice
	if courses == nil {
		courses = []Course{}
	}
	// Create the response structure
	response := map[string]interface{}{
		"courses":   courses,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	params := mux.Vars(r)
	course_id := params["id"]
	for _, c := range courses {
		if c.ID == course_id {
			course = c
			break
		}
	}
	// Create the response structure
	response := map[string]interface{}{
		"course":    course,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	course = course.ReturnCourseWithId()
	courses = append(courses, course)
	// Create the response structure
	response := map[string]interface{}{
		"course":    course,
		"timestamp": time.Now().Format(time.RFC3339),
	}
	// Set the status code for successful creation (201 Created)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	params := mux.Vars(r)
	course_id := params["id"]

	for index, c := range courses {
		if c.ID == course_id {
			courses[index].CourseName = course.CourseName
			courses[index].Price = course.Price
			courses[index].Author = course.Author
			courses[index].ID = course_id
			// Create the response structure
			response := map[string]interface{}{
				"course":    courses[index],
				"detail":    fmt.Sprintf("Course with ID %s is updated", course_id),
				"timestamp": time.Now().Format(time.RFC3339),
			}
			w.WriteHeader(http.StatusOK) // Return 200 OK status
			json.NewEncoder(w).Encode(response)
			return
		}
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	course_id := params["id"]

	for index, course := range courses {
		if course.ID == course_id {
			courses = append(courses[:index], courses[index+1:]...)
		}
	}
	// Create the response structure
	response := map[string]interface{}{
		"detail":    fmt.Sprintf("Course with ID %s is deleted", course_id),
		"timestamp": time.Now().Format(time.RFC3339),
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/api/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/api/course", CreateCourse).Methods("POST")
	r.HandleFunc("/api/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/api/course/{id}", DeleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
