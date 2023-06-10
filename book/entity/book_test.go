package entity

import (
	"reflect"
	"testing"
)

func TestBookEntity_GetID(t *testing.T) {
	type fields struct {
		ID      uint
		Title   string
		SoldOut bool
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "Success",
			fields: fields{
				ID:      1,
				Title:   "Book 1",
				SoldOut: false,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BookEntity{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				SoldOut: tt.fields.SoldOut,
			}
			if got := b.GetID(); got != tt.want {
				t.Errorf("BookEntity.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookEntity_GetTitle(t *testing.T) {
	type fields struct {
		ID      uint
		Title   string
		SoldOut bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success",
			fields: fields{
				ID:      1,
				Title:   "Book 1",
				SoldOut: false,
			},
			want: "Book 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BookEntity{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				SoldOut: tt.fields.SoldOut,
			}
			if got := b.GetTitle(); got != tt.want {
				t.Errorf("BookEntity.GetTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookEntity_IsAvailable(t *testing.T) {
	type fields struct {
		ID      uint
		Title   string
		SoldOut bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Book Available",
			fields: fields{
				ID:      1,
				Title:   "Book 1",
				SoldOut: false,
			},
			want: true,
		},
		{
			name: "Book Not Available",
			fields: fields{
				ID:      1,
				Title:   "Book 1",
				SoldOut: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BookEntity{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				SoldOut: tt.fields.SoldOut,
			}
			if got := b.IsAvailable(); got != tt.want {
				t.Errorf("BookEntity.IsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookEntity_GetEntity(t *testing.T) {
	type fields struct {
		ID      uint
		Title   string
		SoldOut bool
	}
	tests := []struct {
		name   string
		fields fields
		want   BookEntity
	}{
		{
			name: "Book Not Available",
			fields: fields{
				ID:      1,
				Title:   "Book 1",
				SoldOut: true,
			},
			want: BookEntity{
				ID:      1,
				Title:   "Book 1",
				SoldOut: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BookEntity{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				SoldOut: tt.fields.SoldOut,
			}
			if got := b.GetEntity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookEntity.GetEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
