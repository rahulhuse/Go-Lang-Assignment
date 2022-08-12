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

func TestGetAllDepartments(t *testing.T) {
	r := gin.Default()
	r.GET("/collage-api/department", controllers.GetAllDepartments)
	req, err := http.NewRequest("GET", "/collage-api/department", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var dept []models.Department

	json.Unmarshal(w.Body.Bytes(), &dept)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, dept)

}

func TestCreateDepartment(t *testing.T) {
	r := gin.Default()
	r.POST("/collage-api/department", controllers.CreateDepartment)
	department := models.Department{
		DepartmentID:   1,
		DepartmentName: "OO",
		CollageID:      1,
	}
	jsonValue, _ := json.Marshal(department)
	req, err := http.NewRequest("POST", "/collage-api/department", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestGetDepartmentByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/collage-api/department/1", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	r := gin.Default()
	r.GET("/collage-api/department/:department_id", controllers.GetDepartmentByID)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpadateDepartment(t *testing.T) {

	var jsonStr = []byte(
		`{"ID": 2,
		"CreatedAt": "2022-08-10T19:09:05.07+05:30",
		"UpdatedAt": "2022-08-10T19:09:05.07+05:30",
		"DeletedAt": null,
		"departmentid": 1,
		"departmentname": "CIVIL",
		"collageid": 1
	  }`)

	req, err := http.NewRequest("PUT", "/collage-api/department/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/collage-api/department/:department_id", controllers.UpadateDepartment)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// reqNotFound, _ := http.NewRequest("PUT", "/collage-api/department/13", bytes.NewBuffer(jsonValue))
	// w = httptest.NewRecorder()
	// r.ServeHTTP(w, reqNotFound)
	// assert.Equal(t, http.StatusNotFound, w.Code)

}

func TestDeleteDepartmentByID(t *testing.T) {

	req, err := http.NewRequest(http.MethodDelete, "/collage-api/department/5", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/collage-api/department/:department_id", controllers.DeleteDepartmentByID)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
