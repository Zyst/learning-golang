package main

import (
	"reflect"
	"testing"
)

func TestItem_GetByID(t *testing.T) {
	type fields struct {
		id    int
		name  string
		price int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Item
		wantErr bool
	}{
		{"Get Sword", fields{}, args{1}, Item{1, "Steel Sword", 1000}, false},
		{"Get Axe", fields{}, args{2}, Item{2, "Steel Axe", 800}, false},
		{"Get Error", fields{}, args{999}, Item{}, true},
		{"Get Negative Error", fields{}, args{-10029}, Item{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &Item{
				ID:    tt.fields.id,
				Name:  tt.fields.name,
				Price: tt.fields.price,
			}
			got, err := item.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Item.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Item.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_GetAll(t *testing.T) {
	type fields struct {
		id    int
		name  string
		price int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Item
		wantErr bool
	}{
	  {
			"Gotta buy 'em all!",
			fields{},
			[]Item{
				Item{1, "Steel Sword", 1000},
				Item{2, "Steel Axe", 800},
				Item{3, "Maple Bow", 1000},
				Item{4, "Maple Arrow", 20},
				Item{5, "Magic Wand", 2000},
				Item{6, "Oak Shield", 700},
				Item{7, "Steel Dagger", 500},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := &Item{
				ID:    tt.fields.id,
				Name:  tt.fields.name,
				Price: tt.fields.price,
			}
			got, err := item.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Item.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Item.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
