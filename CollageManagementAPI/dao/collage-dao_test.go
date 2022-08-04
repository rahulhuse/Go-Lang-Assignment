package dao

import (
	"gorm-test/models"
	"reflect"
	"testing"
)

func TestCollageNew(t *testing.T) {
	tests := []struct {
		name string
		want CollageService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CollageNew(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CollageNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollageRepo_UpdateCollage(t *testing.T) {
	type args struct {
		collage   *models.Collage
		collageid string
	}
	tests := []struct {
		name        string
		collageRepo *CollageRepo
		args        args
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.collageRepo.UpdateCollage(tt.args.collage, tt.args.collageid); (err != nil) != tt.wantErr {
				t.Errorf("CollageRepo.UpdateCollage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
