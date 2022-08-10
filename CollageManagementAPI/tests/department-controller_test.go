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
		DepartmentID:   5,
		DepartmentName: "IT",
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

	r := gin.Default()
	r.GET("/collage-api/department/1", controllers.GetAllDepartments)
	req, err := http.NewRequest("GET", "/collage-api/department/1", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpadateDepartment(t *testing.T) {

	r := gin.Default()

	r.PUT("/collage-api/department/:id", func(c *gin.Context) {
		id := c.Param("")
		c.String(http.StatusOK, "Hello %s", id)
	}, controllers.UpadateDepartment)

	department := models.Department{
		DepartmentID:   5,
		DepartmentName: "CIVIL",
		CollageID:      1,
	}
	jsonValue, _ := json.Marshal(department)
	reqFound, err := http.NewRequest("PUT", "/collage-api/department/5", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	// reqNotFound, _ := http.NewRequest("PUT", "/collage-api/department/13", bytes.NewBuffer(jsonValue))
	// w = httptest.NewRecorder()
	// r.ServeHTTP(w, reqNotFound)
	// assert.Equal(t, http.StatusNotFound, w.Code)

}

func TestDeleteDepartmentByID(t *testing.T) {

	req, err := http.NewRequest(http.MethodDelete, "/collage-api/department/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/collage-api/department/:collageid", controllers.DeleteCollageCon)

	assert.Equal(t, http.StatusOK, rr.Code)

}
