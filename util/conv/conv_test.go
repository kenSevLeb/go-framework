package conv

import (
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

func TestByte2String(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				b: []byte("hello,world"),
			},
			want: "hello,world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Byte2String(tt.args.b); got != tt.want {
				t.Errorf("Byte2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt2String(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.args.i); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap2Struct(t *testing.T) {
	type args struct {
		mapInstance map[string]interface{}
		obj         interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				mapInstance: map[string]interface{}{
					"hello": "world",
				},
				obj: &struct {
					Hello string
				}{
					Hello: "world",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Map2Struct(tt.args.mapInstance, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Map2Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestString2Byte(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "success",
			args: args{
				s: "hello,world",
			},
			want: []byte("hello,world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2Byte(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String2Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString2Int(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "1",
			},
			want: 1,
		}, {
			name: "invalid",
			args: args{
				"1_a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2Int(tt.args.s); got != tt.want {
				t.Errorf("String2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStruct2Map(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "success",
			args: args{
				obj: struct {
					Hello  string `json:"hello"`
					Number int    `json:"number"`
				}{
					Hello:  "world",
					Number: 2020,
				},
			},
			want: map[string]interface{}{
				"hello":  "world",
				"number": 2020,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Struct2Map(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterToInt(t *testing.T) {
	assert.Equal(t, InterToInt("1"), 1)
	assert.Equal(t, InterToInt(1.0), 1)
	assert.Equal(t, InterToInt(1), 1)
}

func TestInterToStr(t *testing.T) {
	assert.Equal(t, InterToStr(1), "1")
	assert.Equal(t, InterToStr(1.0), "1")
	assert.Equal(t, InterToStr("1"), "1")
}
