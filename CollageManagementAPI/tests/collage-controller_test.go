package controllers

import (
	"bytes"
	"fmt"
	"gorm-test/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllCollagesCon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/collage-api/collage", controllers.GetAllCollagesCon)

	req, err := http.NewRequest(http.MethodGet, "/collage-api/collage", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestCreateCollageCon(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"collage_id":1,"collage_name":"IIT","collage_email":"iit@org.com","collage_mobile":"88885858","collage_address":"Pune"}`)

	req, err := http.NewRequest("POST", "/collage-api/collage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/collage-api/collage", controllers.CreateCollageCon)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"collage_id":1,"collage_name":"IIT","collage_email":"iit@org.com","collage_mobile":"88885858","collage_address":"Pune"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestGetCollageByIDCon(t *testing.T) {

	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/collage-api/collage", controllers.GetStudentByID)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":1,"collage_id":1,"collage_name":"IIT","collage_email":"iit@org.com","collage_mobile":"88885858","collage_address":"Pune"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateCollageCon(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.UpdateCollageCon(tt.args.c)
		})
	}
}

func TestDeleteCollageCon(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.DeleteCollageCon(tt.args.c)
		})
	}
}
