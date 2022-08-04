package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm-test/controllers"
	"gorm-test/database"
	"gorm-test/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

	a := assert.New(t)
	database.ConnectDB()

	collage := models.Collage{
		CollageID:      1,
		CollageName:    "IIT",
		CollageEmail:   "iit@org.com",
		CollageMobile:  "88885858",
		CollageAddress: "Pune",
		Departments:    nil,
		Staffs:         nil,
		Students:       nil,
	}

	reqBody, err := json.Marshal(collage)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateBookRouter(bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Collage{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	actual.Model = gorm.Model{}
	expected := collage
	a.Equal(expected, actual)
}

func setCreateBookRouter(body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	r.POST("/collage-api/collage", controllers.CreateCollageCon)
	req, err := http.NewRequest(http.MethodPost, "/collage-api/collage", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
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
