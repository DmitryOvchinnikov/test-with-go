package underscore

import "testing"

//func TestCamel(t *testing.T) {
//	type args struct {
//		str string
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{"simple", args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
//		{"spaces", args{"with a space"}, "with a space"},
//		{"ends with capital", args{"endsWithA"}, "ends_with_a"},
//	}
//	// setup
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Camel(tt.args.str); got != tt.want {
//				t.Fatalf("Camel() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestCamel(t *testing.T) {
	arg := "thisIsACamelCaseString"
	want := "this_is_a_camel_case_string"
	got := Camel(arg)
	if got != want {
		t.Errorf("Camel(%q) = %q; want %q", arg, got, want)
	}
}
