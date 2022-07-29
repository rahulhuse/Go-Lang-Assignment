package service

import (
	"Collage-Management-App/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUpdateCollage(t *testing.T) {
	type args struct {
		collage   *models.Collage
		collageid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateCollage(tt.args.collage, tt.args.collageid); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCollage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteCollage(t *testing.T) {
	type args struct {
		collage    *models.Collage
		collage_id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteCollage(tt.args.collage, tt.args.collage_id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCollage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateCollage(t *testing.T) {
	type args struct {
		collage *[]models.Collage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateCollage(tt.args.collage); (err != nil) != tt.wantErr {
				t.Errorf("CreateCollage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCollageByID(t *testing.T) {
	type args struct {
		collage   *models.Collage
		collageid string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetCollageByID(tt.args.collage, tt.args.collageid); (err != nil) != tt.wantErr {
				t.Errorf("GetCollageByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllCollages(t *testing.T) {
	type args struct {
		collage *[]models.Collage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetAllCollages(tt.args.collage); (err != nil) != tt.wantErr {
				t.Errorf("GetAllCollages() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
