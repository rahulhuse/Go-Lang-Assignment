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

func TestGetAllStaff(t *testing.T) {
	r := gin.Default()
	r.GET("/collage-api/staff", controllers.GetAllStaff)
	req, err := http.NewRequest("GET", "/collage-api/staff", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var staff []models.Staff

	json.Unmarshal(w.Body.Bytes(), &staff)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, staff)
}

func TestCreateStaff(t *testing.T) {
	r := gin.Default()
	r.POST("/collage-api/staff", controllers.CreateStaff)
	staff := models.Staff{
		StaffID:           5,
		StaffName:         "PP",
		StaffMobileNumber: "7890",
		CollageID:         1,
	}
	jsonValue, _ := json.Marshal(staff)
	req, err := http.NewRequest("POST", "/collage-api/staff", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetStaffByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/collage-api/staff/4", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	r := gin.Default()
	r.GET("/collage-api/staff/:staff_id", controllers.GetStaffByID)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpadateStaff(t *testing.T) {
	var jsonStr = []byte(
		`{"ID": 2,
			"CreatedAt": "0001-01-01T00:00:00Z",
			"UpdatedAt": "2022-08-08T11:04:07.59+05:30",
			"DeletedAt": null,
			"staffid": 2,
			"staffname": "TTT",
			"staffmobilenumber": "222",
			"collageid": 2
		  }`)

	req, err := http.NewRequest("PUT", "/collage-api/staff/2", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/collage-api/staff/:staff_id", controllers.UpadateStaff)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteStaffByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/collage-api/staff/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/collage-api/staff/:staff_id", controllers.DeleteStaffByID)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
