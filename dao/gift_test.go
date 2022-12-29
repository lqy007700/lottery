package dao

import (
	"gorm.io/gorm"
	"log"
	"lottery/datasource"
	"lottery/models"
	"reflect"
	"testing"
)

var db *gorm.DB

func init() {
	db = datasource.New().DB
}

func TestGift_Count(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				db: tt.fields.db,
			}
			if got := g.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGift_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		gift *models.Gift
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				db: tt.fields.db,
			}
			if err := g.Create(tt.args.gift); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGift_Delete(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				db: tt.fields.db,
			}
			if err := g.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGift_Get(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.Gift
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				db: tt.fields.db,
			}
			if got := g.Get(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGift_GetAll(t *testing.T) {
	gift := NewGiftDao(db)
	res := gift.GetAll()
	log.Println(res)
}

func TestGift_Update(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		gift *models.Gift
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gift{
				db: tt.fields.db,
			}
			if err := g.Update(tt.args.gift); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewGiftDao(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *Gift
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGiftDao(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGiftDao() = %v, want %v", got, tt.want)
			}
		})
	}
}
