package strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefault(t *testing.T) {
	type args struct {
		value        string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				value:        "",
				defaultValue: "hello",
			},
			want: "hello",
		},
		{
			name: "01",
			args: args{
				value:        "world",
				defaultValue: "hello",
			},
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplode(t *testing.T) {
	type args struct {
		str       string
		separator string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "01",
			args: args{
				str:       "1,2,3",
				separator: ",",
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Explode(tt.args.str, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				s: "test",
			},
			want: "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.s); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImplode(t *testing.T) {
	type args struct {
		items     interface{}
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				items:     []string{"a", "b"},
				separator: ",",
			},
			want: "a,b",
		},
		{
			name: "02",
			args: args{
				items:     []string{"hello"},
				separator: ",",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Implode(tt.args.items, tt.args.separator); got != tt.want {
				t.Errorf("Implode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{
				s: "test",
			},
			want: "098f6bcd4621d373cade4e832627b4f6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.s); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID(t *testing.T) {
	fmt.Println(UUID())
}

func TestSubStr(t *testing.T) {
	fmt.Println(SubStr("hello world", 5))
}

func TestRand(t *testing.T) {
	fmt.Println(Rand(6))
	fmt.Println(Rand(8))
	fmt.Println(Rand(10))
}
