package main

import (
	"reflect"
	"testing"
)

func Test_cleanInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Space around sentence", args: args{"  hello  world  "}, want: []string{"hello", "world"}},
		{name: "Sentence with CAPS and CamelCase", args: args{"Charmander Bulbasaur PIKACHU"}, want: []string{"charmander", "bulbasaur", "pikachu"}},
		{name: "Unchanged Text", args: args{"hello world"}, want: []string{"hello", "world"}},
		{name: "Empty text", args: args{""}, want: []string{}},
		{name: "Single Word", args: args{"word"}, want: []string{"word"}},
		{name: "Single Uppercase Word", args: args{"Word"}, want: []string{"word"}},
		{name: "Single Caps Word", args: args{"WORD"}, want: []string{"word"}},
		{name: "Multi Line nonsenses", args: args{"Hello\nWorld!"}, want: []string{"hello", "world!"}},
		{name: "Space and multi lines", args: args{"Hello This\nis a world"}, want: []string{"hello", "this", "is", "a", "world"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
