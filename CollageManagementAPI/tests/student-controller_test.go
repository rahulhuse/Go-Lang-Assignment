package tests

import (
	"bytes"
	"encoding/json"
	"gorm-test/controllers"
	"gorm-test/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllStudents(t *testing.T) {
	r := gin.Default()
	r.GET("/collage-api/student", controllers.GetAllStudents)
	req, err := http.NewRequest("GET", "/collage-api/student", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var stu []models.Student

	json.Unmarshal(w.Body.Bytes(), &stu)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, stu)
}

func TestCreateStudent(t *testing.T) {
	r := gin.Default()
	r.POST("/collage-api/student", controllers.CreateStudent)
	student := models.Student{
		StudentID:           4,
		StudentName:         "Lokesh",
		StudentMobileNumber: "67788",
		CollageID:           2,
	}
	jsonValue, _ := json.Marshal(student)
	req, err := http.NewRequest("POST", "/collage-api/student", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetStudentByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/collage-api/student/2", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	r := gin.Default()
	r.GET("/collage-api/student/:student_id", controllers.GetStudentByID)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpadateStudent(t *testing.T) {
	var jsonStr = []byte(
		`{"ID": 1,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"studentid": 1,
			"studentname": "NNN",
			"studentmobilenumber": "1111",
			"collageid": 1
		  }`)

	req, err := http.NewRequest("PUT", "/collage-api/student/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/collage-api/student/:student_id", controllers.UpadateStudent)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteStudentByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/collage-api/student/4", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/collage-api/student/:student_id", controllers.DeleteStudentByID)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
