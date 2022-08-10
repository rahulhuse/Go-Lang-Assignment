package tests

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

	var jsonStr = []byte(`[{"ID":8,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"2022-08-05T18:42:35.764+05:30","DeletedAt":null,"collageid":8,"collagename":"abc","collageemail":"abc","collagemobile":"5567","collageaddress":"Mumbai","Departments":null,"Staffs":null,"Students":null}]`)

	req, err := http.NewRequest("POST", "/collage-api/collage", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/collage-api/collage", controllers.CreateCollageCon)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `[{"ID":8,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"2022-08-05T18:42:35.764+05:30","DeletedAt":null,"collageid":8,"collagename":"abc","collageemail":"abc","collagemobile":"5567","collageaddress":"Mumbai","Departments":null,"Staffs":null,"Students":null}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestGetCollageByIDCon(t *testing.T) {

	req, err := http.NewRequest("GET", "/collage-api/collage/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/collage-api/collage/:collageid", controllers.GetCollageByIDCon)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"2022-08-05T18:42:35.764+05:30","DeletedAt":null,"collageid":1,"collagename":"abc","collageemail":"abc","collagemobile":"5567","collageaddress":"oune","Departments":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"departmentid":1,"departmentname":"IT","collageid":1}],"Staffs":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"staffid":1,"staffname":"abc","staffmobilenumber":"495959","collageid":1}],"Students":[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"studentid":1,"studentname":"qqq","studentmobilenumber":"222","collageid":1}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateCollageCon(t *testing.T) {
	var jsonStr = []byte(`{"ID":8,"CreatedAt":"2022-08-08T13:25:43.214+05:30","UpdatedAt":"2022-08-08T13:39:33.1783916+05:30","DeletedAt":null,"collageid":8,"collagename":"abc","collageemail":"abc","collagemobile":"5567","collageaddress":"Pune","Departments":null,"Staffs":null,"Students":null}`)

	req, err := http.NewRequest("PUT", "/collage-api/collage/8", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.PUT("/collage-api/collage/:collageid", controllers.UpdateCollageCon)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// expected := `{"ID":8,"CreatedAt":"2022-08-08T13:25:43.214+05:30","UpdatedAt":"2022-08-08T13:39:33.1783916+05:30","DeletedAt":null,"collageid":8,"collagename":"abc","collageemail":"abc","collagemobile":"5567","collageaddress":"Pune","Departments":null,"Staffs":null,"Students":null}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}

func TestDeleteCollageCon(t *testing.T) {

	req, err := http.NewRequest(http.MethodDelete, "/collage-api/collage/7", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := gin.Default()
	r.DELETE("/collage-api/collage/:collageid", controllers.DeleteCollageCon)
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
