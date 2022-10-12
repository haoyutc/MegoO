package homework

import (
	"fmt"
	"strings"
	"testing"
)

func TestAlternatePrintNumberAndString(t *testing.T) {
	alternatePrintNumberAndString()
}

func TestStringsCount(t *testing.T) {
	if strings.Count("abcdefaabbcc", "a") == 3 {
		fmt.Println("string a repeat 3 times")
	}
}

func TestStringsIndex(t *testing.T) {
	if strings.Index("abcdef", "c") == 2 {
		fmt.Println("string c index at 2")
	}
}

func TestIsUniqueString(t *testing.T) {
	fmt.Println("aabbccd isUnique: ", isUniqueString("aabbccd"))
	fmt.Println("abcdef isUnique: ", isUniqueString("abcdef"))
	fmt.Println("aabbccd isUnique: ", isUniqueString2("aabbccd"))
	fmt.Println("abcdef isUnique: ", isUniqueString2("abcdef"))
}

func TestReverseString(t *testing.T) {
	fmt.Println(reverseString("abcd"))
}

func Test_reverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
		{name: "return true", args: args{
			s: "abcd",
		}, want: "dcba", want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := reverseString(tt.args.s)
			if got != tt.want {
				t.Errorf("reverseString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("reverseString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isRegroup(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "s1 equal s2", args: args{
			s1: "abcd",
			s2: "bcad",
		}, want: true},
		{name: "s1 not equal s2", args: args{
			s1: "asdbe",
			s2: "asdcw",
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRegroup(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isRegroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name: "replace blank return true",
			args: args{
				s: "as db wqw ",
			},
			want:  "as%20db%20wqw%20",
			want1: true,
		},
		{
			name:  "replace blank return false",
			args:  args{s: "a sd@asa!"},
			want:  "a sd@asa!",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := replaceBlank(tt.args.s)
			if got != tt.want {
				t.Errorf("replaceBlank() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("replaceBlank() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
