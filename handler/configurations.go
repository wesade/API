package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"uni/API/model"
)

var students = []model.Student{
	{
		"1",
		"Sarah",
		28,
	},
	{
		"2",
		"Peter",
		34,
	},
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/getAllStudents")
	json.NewEncoder(w).Encode(students)
}

func GetSingleStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/getSingleStudent")
	studentID := chi.URLParam(r, "id")

	for _, student := range students{
		if student.Id == studentID {
			json.NewEncoder(w).Encode(student)
		}
	}
}

func CreateNewStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/createNewStudent")

	request := model.Student{}
	json.NewDecoder(r.Body).Decode(&request)

	student := model.Student{
		Id: request.Id,
		StudentName: request.StudentName,
		StudentAge: request.StudentAge,
	}

	students = append(students, student)

	w.Write([]byte("Student added"))
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/deleteStudent")
	studentID := chi.URLParam(r, "id")

	for index, student := range students{
		if student.Id == studentID {
			students = append(students[:index], students[index+1:]...)
		}
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/updateStudent")
	studentID := chi.URLParam(r, "id")

	request := model.Student{}
	json.NewDecoder(r.Body).Decode(&request)

	updateStudent := model.Student{
		Id: request.Id,
		StudentName: request.StudentName,
		StudentAge: request.StudentAge,
	}

	for index, _ := range students {
		student := &students[index]
		if student.Id == studentID {
			student.StudentName = updateStudent.StudentName
			student.StudentAge = updateStudent.StudentAge
		}
	}
}
