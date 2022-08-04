package controllers

import (
	"fmt"
	"gorm-test/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllDepartments(t *testing.T) {
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
			controllers.GetAllDepartments(tt.args.c)
		})
	}
}

func TestCreateDepartment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes
	r := gin.Default()
	r.POST("/collage-api/department", controllers.CreateDepartment)

	// Create the mock request you'd like to test. Make sure the second argument
	// here is the same as one of the routes you defined in the router setup
	// block!
	req, err := http.NewRequest(http.MethodPost, "/collage-api/department", nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println(w.Body)

	// Check to see if the response was what you expected
	if w.Code == http.StatusCreated {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusCreated, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusCreated, w.Code)
	}

}
