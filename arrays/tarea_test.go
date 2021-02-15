package arrays

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		Students            []Students
		availableClassrooms uint
	}
	tests := []struct {
		name    string
		args    args
		want    [][]Students
		wantErr bool
	}{
		{name: "Retorna 3 listas de estudiantes", args: args{
			Students:            []Students{{"Pedro", "Zamora", 15}, {"Mariela", "Fernandez", 16}, {"Angel", "Navarro", 15}, {"Ana", "Cuenca", 16}, {"Monica", "Flores", 14}, {"Carlos", "Alvarado", 15}},
			availableClassrooms: 3,
		}, want: [][]Students{
			{{"Pedro", "Zamora", 15}, {"Mariela", "Fernandez", 16}},
			{{"Angel", "Navarro", 15}, {"Ana", "Cuenca", 16}},
			{{"Monica", "Flores", 14}, {"Carlos", "Alvarado", 15}},
		}, wantErr: false},
		{name: "Retorna 2 listas de estudiantes", args: args{
			Students:            []Students{{"Pedro", "Zamora", 15}, {"Mariela", "Fernandez", 16}, {"Angel", "Navarro", 15}, {"Ana", "Cuenca", 16}, {"Monica", "Flores", 14}, {"Carlos", "Alvarado", 15}},
			availableClassrooms: 2,
		}, want: [][]Students{
			{{"Pedro", "Zamora", 15}, {"Mariela", "Fernandez", 16}, {"Angel", "Navarro", 15}},
			{{"Ana", "Cuenca", 16}, {"Monica", "Flores", 14}, {"Carlos", "Alvarado", 15}},
		}, wantErr: false},
		{name: "Retorna 4 listas de estudiantes", args: args{
			Students:            []Students{{"Pedro", "Zamora", 15}, {"Mariela", "Fernandez", 16}, {"Angel", "Navarro", 15}, {"Ana", "Cuenca", 16}, {"Monica", "Flores", 14}, {"Carlos", "Alvarado", 15}},
			availableClassrooms: 4,
		}, want: [][]Students{
			{{"Pedro", "Zamora", 15}, {"Mariela", "Fernandez", 16}},
			{{"Angel", "Navarro", 15}, {"Ana", "Cuenca", 16}},
			{{"Monica", "Flores", 14}},
			{{"Carlos", "Alvarado", 15}},
		}, wantErr: false},
		{name: "Retorna error", args: args{
			Students:            nil,
			availableClassrooms: 4,
		}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := split(tt.args.Students, tt.args.availableClassrooms)
			if (err != nil) != tt.wantErr {
				t.Errorf("split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() got = %v, want %v", got, tt.want)
			}
		})
	}
}
